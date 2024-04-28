package api

import (
	"bytes"
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"net/http"

	"github.com/dchest/siphash"
	"github.com/gorilla/mux"

	"zotregistry.dev/zot/pkg/api/constants"
)

func ClusterProxy(ctrlr *Controller) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
			config := ctrlr.Config
			logger := ctrlr.Log

			// if no cluster or single-node cluster, handle locally
			if config.Cluster == nil || len(config.Cluster.Members) == 1 {
				next.ServeHTTP(response, request)

				return
			}

			vars := mux.Vars(request)
			name, ok := vars["name"]

			if !ok || name == "" {
				response.WriteHeader(http.StatusNotFound)

				return
			}

			h := siphash.New([]byte(config.Cluster.HashKey))
			h.Write([]byte(name))
			sum64 := h.Sum64()

			targetMember := config.Cluster.Members[sum64%uint64(len(config.Cluster.Members))]

			// from the member list and our DNS/IP address, figure out if this request should be handled locally
			localMember := fmt.Sprintf("%s:%s", config.HTTP.Address, config.HTTP.Port)
			logger.Debug().Str(constants.RepositoryLogKey, name).Msg(
				fmt.Sprintf("local member is %s and target member is %s", localMember, targetMember),
			)

			if targetMember == localMember {
				logger.Debug().Str(constants.RepositoryLogKey, name).Msg("handling the request locally")
				next.ServeHTTP(response, request)

				return
			}
			logger.Debug().Str(constants.RepositoryLogKey, name).Msg("proxying the request")

			proxyResponse, err := proxyHTTPRequest(request.Context(), request, targetMember, ctrlr)
			if err != nil {
				http.Error(response, err.Error(), http.StatusInternalServerError)
				logger.Error().Str(constants.RepositoryLogKey, name).Msg(
					fmt.Sprintf("error while proxying request %s", err.Error()),
				)

				return
			}
			defer proxyResponse.Body.Close()

			copyHeader(response.Header(), proxyResponse.Header)
			response.WriteHeader(proxyResponse.StatusCode)
			_, _ = io.Copy(response, proxyResponse.Body)
		})
	}
}

func proxyHTTPRequest(ctx context.Context, req *http.Request,
	targetMember string, ctrlr *Controller,
) (*http.Response, error) {
	cloneURL := *req.URL

	proxyQueryScheme := "http"
	if ctrlr.Config.HTTP.TLS != nil {
		proxyQueryScheme = "https"
	}

	cloneURL.Scheme = proxyQueryScheme
	cloneURL.Host = targetMember

	clonedBody := cloneRequestBody(req)

	fwdRequest, err := http.NewRequestWithContext(ctx, req.Method, cloneURL.String(), clonedBody)
	if err != nil {
		return nil, err
	}

	copyHeader(fwdRequest.Header, req.Header)

	resp, err := getHTTPClient(ctrlr).Do(fwdRequest)
	if err != nil {
		return nil, err
	}

	var b bytes.Buffer
	_, _ = io.Copy(&b, resp.Body)
	resp.Body.Close()
	resp.Body = io.NopCloser(bytes.NewReader(b.Bytes()))

	return resp, nil
}

func getHTTPClient(ctrlr *Controller) *http.Client {
	transport := getTransport(ctrlr)

	return &http.Client{
		Transport: transport,
	}
}

func getTransport(ctrlr *Controller) *http.Transport {
	transport := http.Transport{}

	if ctrlr.Config.HTTP.TLS != nil {
		transport.TLSClientConfig = &tls.Config{
			MinVersion: tls.VersionTLS12,
			RootCAs:    ctrlr.Server.TLSConfig.ClientCAs,
		}
	}

	return &transport
}

func cloneRequestBody(src *http.Request) io.Reader {
	var bCloneForOriginal, bCloneForCopy bytes.Buffer
	multiWriter := io.MultiWriter(&bCloneForOriginal, &bCloneForCopy)
	numBytesCopied, _ := io.Copy(multiWriter, src.Body)

	// If the body is a type of io.NopCloser and length is 0,
	// the Content-Length header is not sent in the proxied request
	// Explicitly returning http.NoBody allows the implementation
	// to set the header
	// Ref: https://github.com/golang/go/issues/34295
	if numBytesCopied == 0 {
		src.Body = http.NoBody

		return http.NoBody
	}

	src.Body = io.NopCloser(&bCloneForOriginal)

	return bytes.NewReader(bCloneForCopy.Bytes())
}

func copyHeader(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}
