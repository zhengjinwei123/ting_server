#!/bin/bash

set -o pipefail

script_name=`basename $0`
script_abs_name=`readlink -f $0`
script_path=`dirname $script_abs_name`

bin_dir=$script_path
proj_home=$bin_dir/..
server_config_file=$proj_home/settings/config.xml
is_kill=0

if [ ! -f "$server_config_file" ]
then
    echo "$server_config_file is missing"
    exit 1
fi

if [ ! -d "$bin_dir"/run ]
then
    mkdir "$bin_dir"/run
fi

wait_last_process() {
	while true
	do
		if [ $is_kill -eq 0 ]; then
			sleep 1000d &
		fi

		sleep_id=$(pgrep sleep)
		if [ "$sleep_id" ]; then
			wait $sleep_id
		fi

		if [ "$(pgrep yut_server)" ]; then
			sleep 1
			continue
		fi

		break
	done
}

kill_server() {
	base "$script_path"/stop.sh

	sleep 3s

	is_kill=1
}


main() {
	trap 'kill_server' SIGTERM

    bash "$script_path"/yut-server.init start $server_config_file
    
	wait_last_process
}

main
