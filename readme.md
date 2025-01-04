# Links

[Deploying a Golang RESTFUL API with Gin, SQLC and PostreSQL](https://dev.to/geoff89/deploying-a-golang-restful-api-with-gin-sqlc-and-postgresql-1lbl)
[Food List with PH values](https://www.webpal.org/SAFE/aaarecovery/2_food_storage/Processing/lacf-phs.htm)

## Libraries

```
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
go install github.com/cosmtrek/air@latest
go get github.com/lib/pq
go get github.com/spf13/viper
go get -u github.com/gin-gonic/gin
https://github.com/golang-migrate/migrate/tree/master/cmd/migrate
```

* `sqlc`: generate go code from sql queries (type-safety)
* `air`: live reloading
* `pq`: `postgres driver`
* `viper`: read configuration 
* `gin`: web framework
* `migrate`: helps with database migrations

## sqlc

Schema doc is available [here](https://docs.sqlc.dev/en/stable/reference/config.html)

## Commands

**create and run migration**

Create migration files

```
migrate create -ext sql -dir db/migration -seq init_schema
```

Update the sql files then run below code to update the database

```
migrate -path server/db/migration -database "postgresql://root:secret@localhost:5432/eldie?sslmode=disable" -verbose up
```

Next, `init` and `configure` `sqlc`:

```
sqlc init # creates a `sqlc.yaml` file, which needs to be populated
```

```yaml
version: "1"
packages:
  - name: "db"
    path: "./db/sqlc"
    queries: "./db/query/"
    schema: "./db/migration/"
    engine: "postgresql"
    emit_json_tags: true
    emit_prepared_queries: true
    emit_interface: false
    emit_exact_table_names: false
    emit_empty_slices: true
```

Create file `db/query/contact.sql` which contain the queries for
which `go` code should be generated.

```sql
-- name: CreateContact :one
INSERT INTO contacts(
    first_name,
    last_name,
    phone_number,
    street,
    created_at,
    updated_at
) VALUES (
    $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: GetContactById :one
SELECT * FROM contacts
WHERE contact_id = $1 LIMIT 1;

-- name: ListContacts :many
SELECT * FROM contacts
ORDER BY contact_id
LIMIT $1
OFFSET $2;

-- name: UpdateContact :one
UPDATE contacts
SET
first_name = coalesce(sqlc.narg('first_name'), first_name),
last_name = coalesce(sqlc.narg('last_name'), last_name),
phone_number = coalesce(sqlc.narg('phone_number'), phone_number),
street = coalesce(sqlc.narg('street'), street),
updated_at = coalesce(sqlc.narg('updated_at'), updated_at)
WHERE contact_id = sqlc.arg('contact_id')
RETURNING *;

-- name: DeleteContact :exec
DELETE FROM contacts
WHERE contact_id = $1;
```

Last, generate code `go` code:

```
sqlc generate
```