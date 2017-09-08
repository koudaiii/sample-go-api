# API Server

- Simple Rest API using gin(framework) & gorm(orm)
- using PostgreSQL

## Usage

### Local

- install & build

```shell-session
$ make
```

- create database

```shell-session
$ psql
psql (9.6.1)
Type "help" for help.

koudaiii=# CREATE DATABASE sample-go-api
koudaiii-# \q
```

- run

```shell-session
$ AUTOMIGRATE=1 bin/sample-go-api
```

### Docker

- using docker-compose

```shell-session
$ docker-compose build
$ docker-compose up
```

## Endpoint list

### Users Resource

```
GET    /
GET    /users
GET    /users/:id
POST   /users
PUT    /users/:id
DELETE /users/:id
```

server runs at http://localhost:8080
