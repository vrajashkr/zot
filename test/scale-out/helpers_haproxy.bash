function generate_haproxy_server_list() {
    local num_instances=${1}
    for ((i=0;i<${num_instances};i++)) do
        local port=$(( 10000 + $i ))
        echo "    server zot${i} 127.0.0.1:${port} check"
    done
}

# stops all haproxy instances started by the test
function haproxy_stop_all() {
    pkill haproxy
}

# starts one haproxy instance with the given config file
# expects the haproxy config to specify daemon mode
function haproxy_start() {
    local haproxy_cfg_file=${1}

    # Check the config file
    haproxy -f ${haproxy_cfg_file} -c >&3

    # Start haproxy
    haproxy -f ${haproxy_cfg_file}
}
