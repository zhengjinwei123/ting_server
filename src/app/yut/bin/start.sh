#!/bin/bash

set -o pipefail

script_name=`basename $0`
script_abs_name=`readlink -f $0`
script_path=`dirname $script_abs_name`

bin_dir=$script_path
proj_home=$bin_dir/..
server_config_file=$proj_home/settings/config.xml

if [ ! -f "$server_config_file" ]
then
    echo "$server_config_file is missing"
    exit 1
fi

if [ ! -d "$bin_dir"/run ]
then
    mkdir "$bin_dir"/run
fi

bash "$script_path"/yut-server.init start $server_config_file
