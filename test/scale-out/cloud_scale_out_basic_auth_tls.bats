# Note: Intended to be run as "make run-cloud-scale-out-tests"
#       Makefile target installs & checks all necessary tooling
#       Extra tools that are not covered in Makefile target needs to be added in verify_prerequisites()

NUM_ZOT_INSTANCES=6
ZOT_CLUSTER_MEMBERS_PATCH_FILE="${BATS_FILE_TMPDIR}/members-patch.json"

load helpers_zot
load helpers_cloud
load helpers_haproxy

# HAProxy runs in SSL Pass-through mode as Zot servers handle TLS
function setup_haproxy() {
    local haproxy_root_dir="${BATS_FILE_TMPDIR}/haproxy"
    local haproxy_cfg_file="${haproxy_root_dir}/haproxy-test.cfg"

    mkdir -p ${haproxy_root_dir}

    cat > ${haproxy_cfg_file}<<EOF
global
    log ${haproxy_root_dir}/log    local0
    log ${haproxy_root_dir}/log    local1 notice
    maxconn 20000
    stats timeout 30s
    daemon

defaults
    log     global
    mode    tcp
    option  tcplog
    option  dontlognull
    timeout connect 5000
    timeout client  50000
    timeout server  50000

frontend zot
    bind *:8000
    default_backend zot-cluster

backend zot-cluster
    option ssl-hello-chk
    balance roundrobin
EOF
    # Populate server list
    generate_haproxy_server_list ${NUM_ZOT_INSTANCES} >> ${haproxy_cfg_file}

    haproxy_start ${haproxy_cfg_file}
}

function launch_zot_server() {
  local zot_server_address=${1}
  local zot_server_port=${2}
  echo "Launching Zot server ${zot_server_address}:${zot_server_port}" >&3

  # Setup zot server config
  local zot_root_dir=${BATS_FILE_TMPDIR}/zot
  mkdir -p ${zot_root_dir}

  local zot_config_file="${BATS_FILE_TMPDIR}/zot_config_${zot_server_address}_${zot_server_port}.json"
  echo "CONFIG: ${zot_config_file}" >&3

  local zot_log_file="${BATS_FILE_TMPDIR}/zot-${zot_server_address}-${zot_server_port}.log"
  echo "LOG: ${zot_log_file}" >&3

  cat > ${zot_config_file}<<EOF
{
    "distSpecVersion": "1.1.0",
    "storage": {
        "rootDirectory": "${zot_root_dir}",
        "dedupe": false,
        "remoteCache": true,
        "storageDriver": {
            "name": "s3",
            "rootdirectory": "/zot",
            "region": "us-east-2",
            "regionendpoint": "localhost:4566",
            "bucket": "zot-storage-test",
            "secure": false,
            "skipverify": false
        },
        "cacheDriver": {
            "name": "dynamodb",
            "endpoint": "http://localhost:4566",
            "region": "us-east-2",
            "cacheTablename": "BlobTable",
            "repoMetaTablename": "RepoMetadataTable",
            "imageMetaTablename": "ImageMetaTable",
            "repoBlobsInfoTablename": "RepoBlobsInfoTable",
            "userDataTablename": "UserDataTable",
            "apiKeyTablename":"ApiKeyTable",
            "versionTablename": "Version"
        }
    },
    "http": {
        "address": "${zot_server_address}",
        "port": "${zot_server_port}",
        "realm": "zot",
        "auth": {
            "htpasswd": {
                "path": "${ZOT_HTPASSWD_PATH}"
            }
        },
        "tls": {
            "cert": "${ROOT_DIR}/test/data/server.cert",
            "key": "${ROOT_DIR}/test/data/server.key",
            "cacert": "${ROOT_DIR}/test/data/ca.crt"
        }
    },
    "cluster": {
      "members": [],
      "hashKey": "loremipsumdolors"
    },
    "log": {
        "level": "debug",
        "output": "${zot_log_file}"
    }
}
EOF
    update_zot_cluster_member_list_in_config_file ${zot_config_file} ${ZOT_CLUSTER_MEMBERS_PATCH_FILE}
    zot_serve ${zot_config_file}
    wait_zot_reachable ${zot_server_port}
}

# Setup function for single zot instance
function setup() {
    # Verify prerequisites are available
    if ! $(verify_prerequisites); then
        exit 1
    fi

    setup_cloud_services
    setup_local_htpasswd

    generate_zot_cluster_member_list ${NUM_ZOT_INSTANCES} ${ZOT_CLUSTER_MEMBERS_PATCH_FILE}

    for ((i=0;i<${NUM_ZOT_INSTANCES};i++)); do
        launch_zot_server 127.0.0.1 $(( 10000 + $i ))
    done

    # list all zot processes that were started
    ps -ef | grep ".*zot.*serve.*" | grep -v grep >&3

    setup_haproxy

    # list haproxy processes that were started
    ps -ef | grep "haproxy" | grep -v grep >&3
}

function teardown() {
    local zot_root_dir=${BATS_FILE_TMPDIR}/zot
    haproxy_stop_all
    zot_stop_all
    rm -rf ${zot_root_dir}
    teardown_cloud_services
}

@test "Check for successful zb run on haproxy frontend" {
    zb_run "cloud-scale-out-basic-auth-tls-bats" "https://127.0.0.1:8000" 1 1 "${ZOT_AUTH_USER}:${ZOT_AUTH_PASS}"
}
