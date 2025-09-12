#!/bin/bash
# wait-for-db.sh
set -e

host="$1"
shift
cmd="$@"

echo "Waiting for database at $host..."

until mysql -h "$host" -u"$MYSQL_USER" -p"$MYSQL_PASSWORD" "$MYSQL_DATABASE" -e "SELECT 1" &> /dev/null; do
  echo "Database not ready yet... retrying in 2 seconds."
  sleep 2
done

echo "Database is up! Starting Apache..."
exec $cmd
