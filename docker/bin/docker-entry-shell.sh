#!/bin/bash
export PATH=$PATH:/bin:/sbin:/usr/bin:/usr/sbin:/usr/local/bin:/usr/local/sbin:~/bin

if [ ! -z $1 ]; then
    docker container exec -it $1 bash
else
    echo "input container id"
fi

exit 0
