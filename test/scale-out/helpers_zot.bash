ROOT_DIR=$(git rev-parse --show-toplevel)
OS=$(go env GOOS)
ARCH=$(go env GOARCH)
ZOT_PATH=${ROOT_DIR}/bin/zot-${OS}-${ARCH}
ZLI_PATH=${ROOT_DIR}/bin/zli-${OS}-${ARCH}
ZOT_MINIMAL_PATH=${ROOT_DIR}/bin/zot-${OS}-${ARCH}-minimal

# basic auth
ZOT_AUTH_USER=poweruser
ZOT_AUTH_PASS=sup*rSecr9T
ZOT_CREDS_PATH="${BATS_FILE_TMPDIR}/creds"
ZOT_HTPASSWD_PATH="${ZOT_CREDS_PATH}/htpasswd"

# zb
ZB_PATH=${ROOT_DIR}/bin/zb-${OS}-${ARCH}
ZB_RESULTS_PATH=${ROOT_DIR}/zb-results
ZB_CI_CD_OUTPUT_FILE=${ROOT_DIR}/ci-cd.json

function verify_prerequisites {
    if [ ! -f ${ZOT_PATH} ]; then
        echo "you need to build ${ZOT_PATH} before running the tests" >&3
        return 1
    fi

    if [ ! -f ${ZB_PATH} ]; then
        echo "you need to build ${ZB_PATH} before running the tests" >&3
        return 1
    fi

    if [ ! $(command -v skopeo) ]; then
        echo "you need to install skopeo as a prerequisite to running the tests" >&3
        return 1
    fi

    if [ ! $(command -v awslocal) ] &>/dev/null; then
        echo "you need to install aws cli as a prerequisite to running the tests" >&3
        return 1
    fi

    if [ ! $(command -v haproxy) ] &>/dev/null; then
        echo "you need to install haproxy as a prerequisite to running the tests" >&3
        return 1
    fi

    return 0
}

function get_free_port(){
    while true
    do
        random_port=$(( ((RANDOM<<15)|RANDOM) % 49152 + 10000 ))
        status="$(nc -z 127.0.0.1 $random_port < /dev/null &>/dev/null; echo $?)"
        if [ "${status}" != "0" ]; then
            free_port=${random_port};
            break;
        fi
    done

    echo ${free_port}
}

function zot_serve() {
    local config_file=${1}
    ${ZOT_PATH} serve ${config_file} &
}

# stops all zot instances started by the test
function zot_stop_all() {
    pkill zot
}

function wait_zot_reachable() {
    local zot_port=${1}
    local zot_url=http://127.0.0.1:${zot_port}/v2/_catalog
    curl --connect-timeout 3 \
        --max-time 5 \
        --retry 20 \
        --retry-delay 1 \
        --retry-max-time 180 \
        --retry-connrefused \
        ${zot_url}
}

function zli_add_config() {
    local registry_name=${1}
    local registry_url=${2}
    # Clean up old configuration for the same registry
    if ${ZLI_PATH} config --list | grep -q ${registry_name}; then
        ${ZLI_PATH} config remove ${registry_name}
    fi
    # Add the new registry
    ${ZLI_PATH} config add ${registry_name} ${registry_url}
}

function zb_run() {
    local test_name=${1}
    local zot_address=${2}
    local concurrent_reqs=${3}
    local num_requests=${4}
    local credentials=${5}

    if [ ! -d "${ZB_RESULTS_PATH}" ]; then
        mkdir -p "${ZB_RESULTS_PATH}"
    fi

    if [ -z "${credentials}" ]; then
        ${ZB_PATH} -c ${concurrent_reqs} -n ${num_requests} ${zot_address} -o ci-cd --skip-cleanup
    else
        ${ZB_PATH} -c ${concurrent_reqs} -n ${num_requests} -A ${credentials} ${zot_address} -o ci-cd --skip-cleanup
    fi

    if [ -f "${ZB_CI_CD_OUTPUT_FILE}" ]; then
        mv "${ZB_CI_CD_OUTPUT_FILE}" "${ZB_RESULTS_PATH}/${test_name}-results.json"
    fi
}

function setup_local_htpasswd() {
    create_htpasswd_file "${ZOT_CREDS_PATH}" "${ZOT_HTPASSWD_PATH}" ${ZOT_AUTH_USER} ${ZOT_AUTH_PASS}
}

function create_htpasswd_file() {
    local creds_dir_path="${1}"
    local htpasswd_file_path="${2}"
    local user=${3}
    local password=${4}

    mkdir -p "${creds_dir_path}"
    htpasswd -b -c -B "${htpasswd_file_path}" ${user} ${password}
}

function generate_zot_cluster_member_list() {
    local num_zot_instances=${1}
    local patch_file_path=${2}
    local temp_file="/tmp/jq-dump.json"
    echo "{\"cluster\":{\"members\":[]}}" > ${patch_file_path}

    for ((i=0;i<${num_zot_instances};i++)); do
        local member="127.0.0.1:$(( 10000 + $i ))"
        jq ".cluster.members += [\"${member}\"]" ${patch_file_path} > ${temp_file} && \
        mv ${temp_file} ${patch_file_path}
    done

    echo "cluster members patch file" >&3
    cat ${patch_file_path} >&3
}

function update_zot_cluster_member_list_in_config_file() {
    local zot_cfg_file=${1}
    local zot_members_patch_file=${2}
    local temp_file="/tmp/jq-dump.json"

    jq -s '.[0] * .[1]' ${zot_cfg_file} ${zot_members_patch_file} > ${temp_file} && \
    mv ${temp_file} ${zot_cfg_file}
}
