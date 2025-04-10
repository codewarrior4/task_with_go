#!/usr/bin/env sh
# wait-for-it.sh

host="$1"
port="$2"
shift 2

until nc -z "$host" "$port"; do
  echo "Waiting for $host:$port to be ready..."
  sleep 2
done

exec "$@"
