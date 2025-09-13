#!/bin/bash

# Scripts used to execute database migration files

echo "Start database migration..."

# Get the migration file directory
MIGRATION_DIR="$(dirname "$(dirname "$0")")/migration"

# Check if the migration directory exists
if [ ! -d "$MIGRATION_DIR" ]; then
  echo "Error: Migration directory $MIGRATION_DIR does not exist"
  exit 1
fi

# Execute migrations in the Docker container
echo "Execute migration files in the MySQL container..."

# Traverse all SQL files in the migration directory
for sql_file in "$MIGRATION_DIR"/*.sql; do
  if [ -f "$sql_file" ]; then
    filename=$(basename "$sql_file")
    echo "Execute migration: $filename"

    # Copy the SQL file to the container and execute it
    docker cp "$sql_file" gva-mysql:/tmp/
    docker exec gva-mysql sh -c "mysql -uroot -p525374187 my_landing_page < /tmp/$filename"

    # Check the execution result
    if [ $? -eq 0 ]; then
      echo "✅ Successfully executed migration: $filename"
    else
      echo "❌ Failed to execute migration: $filename"
      exit 1
    fi
  fi
done

echo "Database migration completed!"
