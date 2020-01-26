#!/bin/bash
export PATH=$PATH:/bin:/sbin:/usr/bin:/usr/sbin:/usr/local/bin:/usr/local/sbin:~/bin

if [ ! -z $1 ]; then
    cat atomy.sql | docker exec -i $1 /usr/bin/mysql -u root --password=123456 atomy
else
    echo "input container id"
fi

exit 0
