# Links

[Deploying a Golang RESTFUL API with Gin, SQLC and PostreSQL](https://dev.to/geoff89/deploying-a-golang-restful-api-with-gin-sqlc-and-postgresql-1lbl)

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

## Commands

**create and run migration**

```
migrate create -ext sql -dir db/migration -seq init_schema
```

Update the sql files then

```
migrate -path db/migration -database "postgresql://root:secret@localhost:5432/eldie?sslmode=disable" -verbose up
```