# matchoshop
E-commerce for men's products

## Requirements

- [Golang](https://golang.org/) as main programming language.
- [Go Module](https://go.dev/blog/using-go-modules) for package management.
- [Goose](https://github.com/steinbacher/goose/) as migration tool.
- [Postgresql](https://www.postgresql.org/) as database driver.
- [Docker-compose](https://docs.docker.com/compose/) for running database container locally.

## Setup
Prepare necessary environment by rename .env.example to .env

```bash
HOST=
PORT=9000
DATABASE_URL=postgres://postgres:mypass@localhost:8010/matchoshop
DB_SSL_MODE=disable
TIMEZONE=Asia/Jakarta
```

Export database environment for migration config
```bash
export DATABASE_URL=postgres://postgres:mypass@localhost:8010/matchoshop
export DB_SSL_MODE=disable
```

Run database container

```bash
docker-compose up
```

## Run the App

Get Go packages

```bash
go get .
```

Update Go package vendor

```bash
go mod vendor
```

Build the app

```bash
go build -o bin/matchoshop -v .
```

Run the App

```bash
./bin/matchoshop
```

## Migration

Create new migration
```bash
goose create AddSomeColumns
```

Up migration
```bash
goose up
```

Down migration
```bash
goose down
```

Check migration status
```bash
goose status
```
