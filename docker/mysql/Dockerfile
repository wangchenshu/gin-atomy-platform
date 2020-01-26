FROM mysql:5.7

ADD initdb.d/atomy.sql /docker-entrypoint-initdb.d

RUN touch /var/log/mysql/mysqld.log