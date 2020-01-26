#!/bin/sh

sed -i "s/<SERVER_PORT>/$SERVER_PORT/g" .env
sed -i "s/<DATABASE_HOST>/$DATABASE_HOST/g" .env
sed -i "s/<DATABASE_PORT>/$DATABASE_PORT/g" .env
sed -i "s/<DATABASE_USER>/$DATABASE_USER/g" .env
sed -i "s/<DATABASE_PASSWORD>/$DATABASE_PASSWORD/g" .env
sed -i "s/<DATABASE_DB>/$DATABASE_DB/g" .env

exec "$@"