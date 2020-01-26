#!/bin/bash
export PATH=$PATH:/bin:/sbin:/usr/bin:/usr/sbin:/usr/local/bin:/usr/local/sbin:~/bin

if [ ! -z $1 ]; then
    docker exec $1 /usr/bin/mysqldump -u root --password=123456 atomy > atomy.sql
else
    echo "input container id"
fi

exit 0
