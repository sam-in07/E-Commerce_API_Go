[The 12 factor App ](https://12factor.net/) 


<img src="picss/workingflow.png" alt="Description" style="width:200%;">


<img src="picss\flow.png" alt="Description" style="width:300%;">


# E-Commerce API

https://roadmap.sh/projects/ecommerce-api

A backend REST API for an e-commerce application built with **Go**, **PostgreSQL**, **Goose**, and **sqlc**.

The project follows a layered structure where HTTP handlers communicate with services, services communicate with the generated sqlc repository, and the repository communicates with PostgreSQL.

---

## Tech Stack

* **Go** — Backend programming language
* **PostgreSQL** — Relational database
* **pgx/v5** — PostgreSQL driver for Go
* **Goose** — Database migration tool
* **sqlc** — Generates type-safe Go code from SQL
* **REST API** — HTTP API
* **JSON** — API request/response format

---

# Project Structure

```text
E-Commerce_API_Go/
│
├── cmd/
│   ├── api.go
│   └── main.go
│
├── internal/
│   │
│   ├── adapters/
│   │   └── postgresql/
│   │       │
│   │       ├── migrations/
│   │       │   ├── 00001_create_products.sql
│   │       │   └── 00002_create_orders.sql
│   │       │
│   │       └── sqlc/
│   │           ├── db.go
│   │           ├── models.go
│   │           ├── queries.sql
│   │           └── queries.sql.go
│   │
│   ├── env/
│   │   └── env.go
│   │
│   ├── json/
│   │   └── json.go
│   │
│   ├── orders/
│   │   ├── handlers.go
│   │   ├── service.go
│   │   └── types.go
│   │
│   └── products/
│       ├── handlers.go
│       └── service.go
│
├── picss/
│   ├── flow.png
│   └── workingflow.png
│
├── .env
├── db_conntection.md
├── go.mod
├── go.sum
├── Readme.md
└── sqlc.yaml
```

---

# Architecture

The general application flow is:

```text
                    HTTP Request
                         │
                         ▼
                  ┌─────────────┐
                  │   Handler   │
                  └──────┬──────┘
                         │
                         ▼
                  ┌─────────────┐
                  │   Service   │
                  └──────┬──────┘
                         │
                         ▼
                  ┌─────────────┐
                  │    sqlc     │
                  │  Repository │
                  └──────┬──────┘
                         │
                         ▼
                  ┌─────────────┐
                  │ PostgreSQL  │
                  └─────────────┘
```

For orders:

```text
Client
  │
  │ POST /orders
  ▼
Order Handler
  │
  ▼
Order Service
  │
  ├── Validate customer
  │
  ├── Validate order items
  │
  ├── Find products
  │
  ├── Check stock
  │
  ├── Create order
  │
  ├── Create order items
  │
  ├── Update product stock
  │
  └── Commit transaction
  │
  ▼
PostgreSQL
```

---

# Requirements

Before running the project, install:

* Go 1.21+
* PostgreSQL
* sqlc
* Goose

Docker is **not required** for this project because PostgreSQL is installed directly on Windows.

---

# PostgreSQL Setup

This project uses a local PostgreSQL server.

Current configuration:

```text
Host:      localhost
Port:      5432
User:      postgres
Password:  admin
Database:  ecom-api-go
```

The database name is:

```text
ecom-api-go
```

---

# Create the Database

If the database does not exist, create it from PostgreSQL:

```sql
CREATE DATABASE "ecom-api-go";
```

Then connect:

```powershell
psql -U postgres -d ecom-api-go
```

Enter the PostgreSQL password when requested.

Expected prompt:

```text
ecom-api-go=#
```

---

# Test PostgreSQL

Inside `psql`, run:

```sql
\conninfo
```

This should show information similar to:

```text
Database      | ecom-api-go
Client User   | postgres
Host          | localhost
Server Port   | 5432
```

Check database tables:

```sql
\dt
```

Exit PostgreSQL:

```sql
\q
```

---

# Environment Configuration

Create a `.env` file in the project root:

```env
GOOSE_DBSTRING="host=localhost user=postgres password=admin dbname=ecom-api-go sslmode=disable"
GOOSE_DRIVER=postgres
GOOSE_MIGRATION_DIR=./internal/adapters/postgresql/migrations
```

### Database connection

The important part is:

```text
host=localhost
user=postgres
password=admin
dbname=ecom-api-go
sslmode=disable
```

Make sure `dbname` matches the actual PostgreSQL database.

---

# Installing sqlc

Install sqlc using Go:

```powershell
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```

Verify the installation:

```powershell
sqlc version
```

Example:

```text
v1.31.1
```

sqlc is installed as a binary and does not need to be installed separately inside every project.

---

# Installing Goose

Install Goose:

```powershell
go install github.com/pressly/goose/v3/cmd/goose@latest
```

Verify:

```powershell
goose version
```

---

# Goose Migrations

Database migrations are stored here:

```text
internal/
└── adapters/
    └── postgresql/
        └── migrations/
            ├── 00001_create_products.sql
            └── 00002_create_orders.sql
```

Goose keeps track of applied migrations using:

```text
goose_db_version
```

---

# Migration File Structure

A Goose SQL migration normally follows this structure:

```sql
-- +goose Up

CREATE TABLE example (
    id BIGSERIAL PRIMARY KEY
);

-- +goose Down

DROP TABLE example;
```

The `Up` section creates or changes the database.

The `Down` section reverses the migration.

---

# Run Migrations

Always run Goose from the project root:

```powershell
cd D:\webProjects\E-Commerce_API_Go
```

Then:

```powershell
goose up
```

Example:

```text
OK   00001_create_products.sql
OK   00002_create_orders.sql

goose: successfully migrated database to version: 2
```

---

# Check Migration Status

```powershell
goose status
```

This shows which migrations have already been applied and which are still pending.

---

# Check Current Database Version

```powershell
goose version
```

---

# Create a New Migration

When adding a new database change:

```powershell
goose create add_something sql
```

For example:

```powershell
goose create add_users sql
```

This creates a new migration file.

Then edit the generated SQL file.

After creating the migration:

```powershell
goose up
```

---

# Roll Back a Migration

To roll back the latest migration:

```powershell
goose down
```

Be careful with rollback commands on databases containing important data.

---

# sqlc

sqlc converts SQL queries into type-safe Go code.

The generated files are located here:

```text
internal/adapters/postgresql/sqlc/
```

Example:

```text
sqlc/
├── db.go
├── models.go
├── queries.sql
└── queries.sql.go
```

---

# sqlc Configuration

The project configuration is:

```text
sqlc.yaml
```

This file tells sqlc:

* Where SQL schema files are located
* Where SQL queries are located
* Which PostgreSQL version is being used
* Where generated Go files should be written
* Which Go package name should be used

---

# Generate sqlc Code

After changing SQL queries or schema:

```powershell
sqlc generate
```

Run this from the project root:

```powershell
PS D:\webProjects\E-Commerce_API_Go>
```

Do not manually edit generated files such as:

```text
queries.sql.go
models.go
```

Instead, change the SQL source and run:

```powershell
sqlc generate
```

again.

---

# Example sqlc Query

A list-products query:

```sql
-- name: ListProducts :many
SELECT *
FROM public.products
ORDER BY id ASC;
```

sqlc generates the corresponding Go method.

For example:

```go
products, err := q.ListProducts(ctx)
```

---

# Creating a Product

Example SQL:

```sql
INSERT INTO public.products (
    name,
    price_in_centers,
    quantity
)
VALUES (
    'iPhone 15',
    79900,
    10
);
```

Check the result:

```sql
SELECT *
FROM public.products
ORDER BY id ASC;
```

Prices are stored in cents.

For example:

```text
79900 = $799.00
```

---

# Product Model

The generated product model currently looks similar to:

```go
type Product struct {
    ID             int64
    Name           string
    PriceInCenters int32
    Quantity       int32
    CreatedAt      pgtype.Timestamptz
}
```

---

# Orders

Orders are handled inside:

```text
internal/orders/
├── handlers.go
├── service.go
└── types.go
```

The order service is responsible for business logic such as:

* Validating the customer
* Checking order items
* Finding products
* Checking available stock
* Creating the order
* Creating order items
* Updating product stock
* Committing the transaction

---

# Database Transactions

Order creation uses a PostgreSQL transaction.

General flow:

```go
tx, err := db.Begin(ctx)
```

Then:

```go
qtx := repo.WithTx(tx)
```

All order-related operations use the transaction.

If something fails:

```go
tx.Rollback(ctx)
```

If everything succeeds:

```go
tx.Commit(ctx)
```

This prevents partially created orders.

For example:

```text
Create Order
     │
     ├── Create Order Item
     │
     ├── Update Product Stock
     │
     └── Commit
```

If one operation fails:

```text
Create Order
     │
     ├── Create Order Item
     │
     ├── Update Product Stock ❌
     │
     └── Rollback
```

---

# Error Handling

The orders service defines errors such as:

```go
var (
    ErrProductNotFound = errors.New("product not found")
    ErrProductNoStock  = errors.New("product has not enough stock")
)
```

These errors represent business-level failures.

---

# Running the API

From the project root:

```powershell
go run ./cmd
```

Expected output:

```text
connected to database
server has started at addr :8080
```

The API is available at:

```text
http://localhost:8080
```

---

# Typical Development Workflow

When starting work on the project:

```powershell
cd D:\webProjects\E-Commerce_API_Go
```

Check PostgreSQL is running.

Then check migrations:

```powershell
goose status
```

Apply pending migrations:

```powershell
goose up
```

Generate sqlc code:

```powershell
sqlc generate
```

Run the API:

```powershell
go run ./cmd
```

---

# Recommended Workflow After Changing the Database

When adding a new table or changing an existing table:

```text
1. Create Goose migration
        ↓
2. Write SQL migration
        ↓
3. Run goose up
        ↓
4. Verify PostgreSQL schema
        ↓
5. Update sqlc queries
        ↓
6. Run sqlc generate
        ↓
7. Update Go service
        ↓
8. Update handlers
        ↓
9. Test API
```

Example:

```powershell
goose create add_users sql
```

Then:

```powershell
goose up
```

Then:

```powershell
sqlc generate
```

Then:

```powershell
go run ./cmd
```

---

# Important Rule

Do not manually modify generated sqlc files.

Avoid changing:

```text
internal/adapters/postgresql/sqlc/models.go
internal/adapters/postgresql/sqlc/queries.sql.go
```

Instead modify:

```text
migrations/
queries.sql
sqlc.yaml
```

and regenerate:

```powershell
sqlc generate
```

---

# Docker

Docker is **not required** for this project.

The original project may contain a `docker-compose.yaml` with PostgreSQL configuration, but when PostgreSQL is installed directly on Windows, the application can connect directly to:

```text
localhost:5432
```

Therefore, this project currently uses:

```text
Go API
   │
   ▼
Local PostgreSQL
   │
   ▼
ecom-api-go
```

instead of:

```text
Go API
   │
   ▼
Docker
   │
   ▼
PostgreSQL container
```

---

# Useful PostgreSQL Commands

Connect to the database:

```powershell
psql -U postgres -d ecom-api-go
```

List databases:

```sql
\l
```

List tables:

```sql
\dt
```

Describe a table:

```sql
\d products
```

Show connection information:

```sql
\conninfo
```

Exit:

```sql
\q
```

---

# Useful Go Commands

Download dependencies:

```powershell
go mod download
```

Synchronize dependencies:

```powershell
go mod tidy
```

Run the application:

```powershell
go run ./cmd
```

Build:

```powershell
go build ./...
```

Run tests:

```powershell
go test ./...
```

Format Go code:

```powershell
gofmt -w .
```

---

# Troubleshooting

## PostgreSQL connection failed

Check that PostgreSQL is running.

Test:

```powershell
psql -U postgres -d ecom-api-go
```

Verify:

```text
host=localhost
port=5432
user=postgres
password=admin
dbname=ecom-api-go
```

---

## Database has no tables

Check:

```sql
\dt
```

If only `goose_db_version` exists, run:

```powershell
goose up
```

---

## sqlc method does not exist

For example:

```text
qtx.CreateOrder undefined
```

This usually means sqlc has not generated that method.

Check that the corresponding query exists in:

```text
queries.sql
```

For example:

```sql
-- name: CreateOrder :one
...
```

Then run:

```powershell
sqlc generate
```

---

## sqlc parameter type does not exist

For example:

```text
undefined: repo.CreateOrderItemParams
```

Make sure the SQL query has the correct sqlc name:

```sql
-- name: CreateOrderItem :one
```

Then run:

```powershell
sqlc generate
```

The generated parameter type should then be available.

---

## Service interface error

If you see:

```text
*svc does not implement Service
```

check that the interface method and implementation have exactly the same:

* Method name
* Parameters
* Return values

Example:

```go
type Service interface {
    ListProducts(ctx context.Context) ([]repo.Product, error)
}
```

must match:

```go
func (s *svc) ListProducts(ctx context.Context) ([]repo.Product, error)
```

---

# Current Database

The current development database is:

```text
Database: ecom-api-go
```

Current known tables:

```text
goose_db_version
products
orders
```

The exact schema should always be verified using:

```sql
\dt
```

---

# Development Checklist

Before starting development:

* [ ] PostgreSQL is running
* [ ] `ecom-api-go` database exists
* [ ] `.env` is configured
* [ ] Goose is installed
* [ ] sqlc is installed
* [ ] Goose migrations are applied
* [ ] `goose status` is correct
* [ ] sqlc code is generated
* [ ] Go application starts
* [ ] Database connection succeeds
* [ ] API starts on port `8080`

---

# Quick Start

For an already configured project:

```powershell
# Go to project
cd D:\webProjects\E-Commerce_API_Go

# Check migrations
goose status

# Apply migrations
goose up

# Generate database code
sqlc generate

# Start API
go run ./cmd
```

Expected result:

```text
connected to database
server has started at addr :8080
```

---

# Project Flow

```text
                  ┌───────────────┐
                  │    Client     │
                  └───────┬───────┘
                          │
                          ▼
                  ┌───────────────┐
                  │    Handler    │
                  └───────┬───────┘
                          │
                          ▼
                  ┌───────────────┐
                  │    Service    │
                  └───────┬───────┘
                          │
                          ▼
                  ┌───────────────┐
                  │     sqlc      │
                  └───────┬───────┘
                          │
                          ▼
                  ┌───────────────┐
                  │  PostgreSQL   │
                  └───────────────┘
```

Database changes follow:

```text
Migration SQL
      │
      ▼
   Goose
      │
      ▼
PostgreSQL Schema
      │
      ▼
    sqlc
      │
      ▼
Generated Go Code
      │
      ▼
Go Service
      │
      ▼
HTTP Handler
```

---

# Official Documentation

* sqlc — installation and usage
* Goose — migrations and CLI usage

---

# Notes

This project uses local PostgreSQL instead of Docker.

The database credentials in the current development environment are:

```text
user=postgres
password=admin
database=ecom-api-go
host=localhost
port=5432
```

For production, credentials should **not** be committed to the repository. Use environment variables or a secure secrets manager.

Also make sure `.env` is included in `.gitignore` before pushing the project to GitHub.

Example:

```gitignore
.env
```
