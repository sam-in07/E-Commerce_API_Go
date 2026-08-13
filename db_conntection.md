Go + PostgreSQL + Goose + sqlc — Windows Setup Guide

A quick setup reference for Go API projects using:

Go
PostgreSQL installed directly on Windows
Goose for database migrations
sqlc for generating Go database code
.env for database configuration

Docker is not required when PostgreSQL is already installed and running directly on Windows.

1. PostgreSQL Setup

Use an existing PostgreSQL installation.

Example credentials:

Host:     localhost
Port:     5432
User:     postgres
Password: admin
Database: <project-database-name>


For this project:

Database: ecom-api-go

Test PostgreSQL

From PowerShell:

psql -U postgres -d ecom-api-go


Enter the PostgreSQL password.

If connected successfully:

ecom-api-go=#


Check the connection:

\conninfo


Check tables:

\dt


Exit:

\q

2. Do NOT create a Docker PostgreSQL container

If PostgreSQL is already installed locally, a docker-compose.yaml such as:

services:
  postgres:
    image: postgres:16-alpine


is not necessary.

Do not run:

docker compose up


unless you intentionally want to use Docker PostgreSQL.

The local PostgreSQL server is enough.

3. Go Project

Go project example:

D:\webProjects\E-Commerce_API_Go


Go should already be installed.

Check:

go version


Run the application:

go run ./cmd


Expected output:

connected to database
server has started at addr :8080

4. Environment Variables

Create a .env file in the project root.

Example:

GOOSE_DBSTRING="host=localhost user=postgres password=admin dbname=ecom-api-go sslmode=disable"
GOOSE_DRIVER=postgres
GOOSE_MIGRATION_DIR=./internal/adapters/postgresql/migrations


Important:

The dbname must match the PostgreSQL database you actually created.

Example:

ecom-api-go


Do not accidentally use:

ecom


if the project is configured to use:

ecom-api-go

5. Go Database Connection

The application should use the same database connection information.

Example:

dsn: env.GetString(
    "GOOSE_DBSTRING",
    "host=localhost user=postgres password=admin dbname=ecom-api-go sslmode=disable",
),


The .env value and the fallback value should point to the same database.

6. Install sqlc

sqlc is installed separately from the project.

You do NOT need to create a special folder.

Run:

go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest


You can run this from anywhere.

Verify:

sqlc version


Example:

v1.31.1


Once installed, you don't need to install sqlc again for every project.

7. Goose Migrations

Goose manages the database schema.

Migration files should be inside the configured directory:

internal/
└── adapters/
    └── postgresql/
        └── migrations/


Example:

00001_create_products.sql
00002_create_users.sql
00003_create_orders.sql

Run migrations

Go to the project root:

cd D:\webProjects\E-Commerce_API_Go


Then:

goose up


Expected output:

OK   00001_create_products.sql
goose: successfully migrated database to version: 1


This means the migration was successfully applied.

8. Verify Database Tables

Connect to PostgreSQL:

psql -U postgres -d ecom-api-go


Then:

\dt


Expected example:

Schema |       Name       | Type  | Owner
-------+------------------+-------+--------
public | goose_db_version | table | postgres
public | products         | table | postgres


goose_db_version is created by Goose to track migrations.

Your application tables, such as:

products
users
orders


are created by your migration files.

9. sqlc

After the database/schema and SQL queries are correctly configured, run:

sqlc generate


Run this from the project root:

D:\webProjects\E-Commerce_API_Go


sqlc reads the project's sqlc.yaml configuration and generates Go database code.

Do NOT manually write the generated sqlc files.

10. Recommended Setup Order

For a new project, follow this order:

1. Install Go
       ↓
2. Install PostgreSQL
       ↓
3. Create project database
       ↓
4. Configure .env
       ↓
5. Verify PostgreSQL connection
       ↓
6. Set up Goose migrations
       ↓
7. Run: goose up
       ↓
8. Verify tables with: \dt
       ↓
9. Configure sqlc.yaml
       ↓
10. Run: sqlc generate
       ↓
11. Run Go API
       ↓
12. Test API at localhost:8080

11. Useful Commands
Check Go
go version

Check sqlc
sqlc version

Check Goose
goose version

Start Go API
go run ./cmd

Run all pending migrations
goose up

Roll back the latest migration
goose down

Check PostgreSQL connection
psql -U postgres -d <database-name>

List tables

Inside psql:

\dt

Show connection information
\conninfo

Exit psql
\q

12. Troubleshooting
"database does not exist"

Create the database first:

CREATE DATABASE <database-name>;


Then connect again.

"password authentication failed"

Check:

User
Password
Host
Port


Example:

user=postgres
password=admin
host=localhost
port=5432

"\dt shows no tables"

The database exists, but migrations probably haven't been run.

From the project root:

goose up


Then check:

\dt

"sqlc is not recognized"

Check:

sqlc version


If it isn't found, Go's binary directory may not be in the Windows PATH.

API starts but database connection fails

Check the DSN:

host=localhost
user=postgres
password=admin
dbname=<correct-database>
sslmode=disable


Make sure the database name exactly matches PostgreSQL.

13. Final Checklist

Before considering the project database setup complete:

 PostgreSQL installed and running
 Database created
 PostgreSQL credentials verified
 .env configured
 Go can connect to PostgreSQL
 Goose installed
 Migration files exist
 goose up succeeds
 \dt shows application tables
 sqlc installed
 sqlc.yaml configured
 sqlc generate succeeds
 Go API starts successfully
 API responds on localhost:8080
The short version

For the next Go project, remember:

# 1. Go to project
cd D:\path\to\project

# 2. Configure .env
# PostgreSQL connection

# 3. Run migrations
goose up

# 4. Check PostgreSQL
psql -U postgres -d <database-name>

# 5. Check tables
\dt

# 6. Generate sqlc code
sqlc generate

# 7. Start API
go run ./cmd


No Docker is necessary if PostgreSQL is installed and running locally.


...........E-Commerce_API_Go> goose -s create create_orders  sql  

2026/08/13 21:03:43 Created new file: internal\adapters\postgresql\migrations\00002_create_orders.sql


...........E-Commerce_API_Go> goose up 

2026/08/13 21:04:40 OK   00002_create_orders.sql (69.68ms)

2026/08/13 21:04:40 goose: successfully migrated database to version: 2

...........E-Commerce_API_Go>  sqlc generate



goose -s create create_products  sql

sqlc generate

postgres://postgres:admin@localhost:5432/ecom?sslmode=disable

