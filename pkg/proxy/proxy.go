package proxy

import (
	"bytes"
	"context"
	"io"
	"net/http"

	"zotregistry.dev/zot/pkg/api/config"
	"zotregistry.dev/zot/pkg/api/constants"
	"zotregistry.dev/zot/pkg/common"
)

// proxy the request to the target member and return a pointer to the response or an error.
func ProxyHTTPRequest(ctx context.Context, req *http.Request,
	targetMember string, config *config.Config,
) (*http.Response, error) {
	cloneURL := *req.URL

	proxyQueryScheme := "http"
	if config.HTTP.TLS != nil {
		proxyQueryScheme = "https"
	}

	cloneURL.Scheme = proxyQueryScheme
	cloneURL.Host = targetMember

	clonedBody := CloneRequestBody(req)

	fwdRequest, err := http.NewRequestWithContext(ctx, req.Method, cloneURL.String(), clonedBody)
	if err != nil {
		return nil, err
	}

	CopyHeader(fwdRequest.Header, req.Header)

	// always set hop count to 1 for now.
	// the handler wrapper above will terminate the process if it sees a request that
	// already has a hop count but is due for proxying.
	fwdRequest.Header.Set(constants.ScaleOutHopCountHeader, "1")

	clientOpts := common.HTTPClientOptions{
		TLSEnabled: config.HTTP.TLS != nil,
		VerifyTLS:  config.HTTP.TLS != nil, // for now, always verify TLS when TLS mode is enabled
		Host:       targetMember,
	}

	tlsConfig := config.Cluster.TLS
	if tlsConfig != nil {
		clientOpts.CertOptions.ClientCertFile = tlsConfig.Cert
		clientOpts.CertOptions.ClientKeyFile = tlsConfig.Key
		clientOpts.CertOptions.RootCaCertFile = tlsConfig.CACert
	}

	httpClient, err := common.CreateHTTPClient(&clientOpts)
	if err != nil {
		return nil, err
	}

	resp, err := httpClient.Do(fwdRequest)
	if err != nil {
		return nil, err
	}

	var clonedRespBody bytes.Buffer

	// copy out the contents into a new buffer as the response body
	// stream should be closed to get all the data out.
	_, _ = io.Copy(&clonedRespBody, resp.Body)
	resp.Body.Close()

	// after closing the original body, substitute it with a new reader
	// using the buffer that was just created.
	// this buffer should be closed later by the consumer of the response.
	resp.Body = io.NopCloser(bytes.NewReader(clonedRespBody.Bytes()))

	return resp, nil
}

func CloneRequestBody(src *http.Request) io.Reader {
	var bCloneForOriginal, bCloneForCopy bytes.Buffer
	multiWriter := io.MultiWriter(&bCloneForOriginal, &bCloneForCopy)
	numBytesCopied, _ := io.Copy(multiWriter, src.Body)

	// if the body is a type of io.NopCloser and length is 0,
	// the Content-Length header is not sent in the proxied request.
	// explicitly returning http.NoBody allows the implementation
	// to set the header.
	// ref: https://github.com/golang/go/issues/34295
	if numBytesCopied == 0 {
		src.Body = http.NoBody

		return http.NoBody
	}

	src.Body = io.NopCloser(&bCloneForOriginal)

	return bytes.NewReader(bCloneForCopy.Bytes())
}

func CopyHeader(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}
