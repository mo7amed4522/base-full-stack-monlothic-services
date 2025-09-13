# Database Migration Guide

本The directory contains the project's database migration files, which are used to update and modify the database structure.。

## Migrate files

- `add_google_fields_to_bx_users.sql`: Add Google login related fields to the bx_users table

## Execute migrations

The project provides two scripts to execute database migrations:

1. **Only execute migrations**

   ```bash
   ./server/scripts/run_migrations.sh
   ```

   This script will execute all SQL files in the `migration` directory.

2. **Update Docker container and execute migrations**

   ```bash
   ./server/scripts/update_docker_and_migrate.sh
   ```

   This script will stop and restart the Docker container, then execute the migration script.

## Add new migration files

1. Create a new SQL file in the `migration` directory, with a filename that describes the migration purpose
2. Write the SQL statements to implement the required database changes
3. Use the above scripts to execute the migrations

## Attention

- Migration files should be idempotent, i.e., they can be executed multiple times without causing errors
- Before executing migrations, ensure that the database has been backed up
- Migration files should contain clear comments, describing their purpose and impact
