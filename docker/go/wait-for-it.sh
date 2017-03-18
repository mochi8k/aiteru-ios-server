#!/bin/bash

host=redis
port=6379
cmd=$@

echo Host:$host
echo Port:$port
echo CMD:$cmd

until redis-cli -h $host -p $port ping; do
  >&2 echo "Redis is unavailable -sleeping"
  sleep 1
done

>&2 echo "Redis is up -executing command"

exec $cmd
