# Go Fication
An API Boilerplate written in Golang with chi-route and Gorm. Write restful API with fast development and developer friendly.

## Table of Contents

- [Architecture](#architecture)
- [Boilerplate Structure](#boilerplate-structure)
- [Configuration Manage](#configuration-manage)
  - [ENV Manage](#env-manage)
  - [Server Configuration](#server-configuration)
  - [Database Configuration](#database-configuration)
  - [PgAdmin](#pg-admin)
- [Installation](#installation)
  - [Local Setup Instruction](#local-setup-instruction)
  - [Develop Application in Docker with Live Reload](#develop-application-in-docker-with-live-reload)
- [Let's Build an API](#lets-build-an-api)
- [Code Examples](#code-examples)
- [Deployment](#deployment)
  - [Container Development Build](#container-development-build)
  - [Container Production Build and Up](#container-production-build-and-up)
- [Useful Commands](#useful-commands)
- [ENV YAML Configure](#env-yaml-configure)
- [Use Packages](#use-packages)
### Architecture
#### In this project use 3 layer architecture
 - Models
 - Repository
 - Controllers
- More About [Uncle Bob’s Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
### Boilerplate Structure
<pre>├── <font color="#3465A4"><b>config</b></font>
├── <font color="#3465A4"><b>controllers</b></font>
├── <font color="#3465A4"><b>helpers</b></font>
├── <font color="#3465A4"><b>infra</b></font>
│   ├── <font color="#3465A4">database</font>
│   └── <font color="#3465A4">logger</font>
├── <font color="#3465A4"><b>migrations</b></font>
├── <font color="#3465A4"><b>models</b></font>
├── <font color="#3465A4"><b>repository</b></font>
├── <font color="#3465A4"><b>routers</b></font>
│   ├── <font color="#3465A4">middlewares</font>
</pre>
### Configuration Manage
#### ENV Manage

- Default ENV Configuration Manage from `.env`. sample file `.env.example`
```text
# Server Configuration
SECRET=h9wt*pasj6796j##w(w8=xaje8tpi6h*r&hzgrz065u&ed+k2)
DEBUG=True # `False` in Production
ALLOWED_HOSTS=0.0.0.0
SERVER_HOST=0.0.0.0
SERVER_PORT=8000

# Database Configuration
MASTER_DB_NAME=test_pg_go
MASTER_DB_USER=mamun
MASTER_DB_PASSWORD=123
MASTER_DB_HOST=postgres_db
MASTER_DB_PORT=5432
MASTER_DB_LOG_MODE=True # `False` in Production
MASTER_SSL_MODE=disable

REPLICA_DB_NAME=test_pg_go
REPLICA_DB_USER=mamun
REPLICA_DB_PASSWORD=123
REPLICA_DB_HOST=localhost
REPLICA_DB_PORT=5432
REPLICA_DB_LOG_MODE=True # `False` in Production
REPLICA_SSL_MODE=disable
```
- Server `DEBUG` set `False` in Production
- Database Logger `MASTER_DB_LOG_MODE` and `REPLICA_DB_LOG_MODE`  set `False` in production
- If ENV Manage from YAML file add a config.yml file and configuration [db.go](pkg/config/db.go) and [server.go](pkg/config/server.go). See More [ENV YAML Configure](#env-yaml-configure)

#### Server Configuration
- Use [chi](https://github.com/go-chi/chi) Route

#### Database Configuration
- Use [GORM](https://github.com/go-gorm/gorm) as an ORM
- Use database `MASTER_DB_HOST` value set as `localhost` for local development, and use `postgres_db` for docker development 
#### PG Admin
- Check  PG Admin on [http://0.0.0.0:5050/browser/](http://0.0.0.0:5050/browser/)
- Login with Credential Email `admin@admin.com` Password `root`
- Connect Database Host as `postgres_db`, DB Username and Password as per `.env` set
- Note: if not configure `.env`, default Username `mamun` and password `123`

### Installation
#### Local Setup Instruction
Follow these steps:
- Copy [.env.example](.env.example) as `.env` and configure necessary values
- To add all dependencies for a package in your module `go get .` in the current directory
- Locally run `go run main.go` or `go build main.go` and run `./main`
- Check Application health available on [0.0.0.0:8000/health](http://0.0.0.0:8000/health)

#### Develop Application in Docker with Live Reload
Follow these steps:
- Make sure install the latest version of docker and docker-compose
- Docker Installation for your desire OS https://docs.docker.com/engine/install/ubuntu/
- Docker Composer Installation https://docs.docker.com/compose/install/
- Run and Develop `make dev`
- Check Application health available on [0.0.0.0:8000/health](http://0.0.0.0:8000/health)

### Let's Build an API

1. [models](models) folder add a new file name `example_model.go`

```go
package models

import (
	"time"
)

type Example struct {
	Id        int        `json:"id"`
	Data      string     `json:"data" binding:"required"`
	CreatedAt *time.Time `json:"created_at,string,omitempty"`
	UpdatedAt *time.Time `json:"updated_at_at,string,omitempty"`
}
// TableName is Database Table Name of this model
func (e *Example) TableName() string {
	return "examples"
}
```
2. Add Model to [migration](migrations/migrations.go)

```go
package migrations

import (
	"go-fication/models"
	"go-fication/infra/database"
)

func Migrate() {
	var migrationModels = []interface{}{&models.Example{}}
	err := database.GetDB().AutoMigrate(migrationModels...)
	if err != nil {
		return
	}
}
```
3. [repository](repository) folder add a file `example_repo.go`
```go
package repository

import (
	"go-fication/helpers/pagination"
	"go-fication/models"
)

type ExampleRepo interface {
	GetExamples(limit, offset int64) (res interface{}, err error)
	CreateExample(exp *models.Example) error
}

func (r *GormRepository) GetExamples(limit, offset int64) (res interface{}, err error) {
	var example []*models.Example
	res = pagination.Paginate(&pagination.Param{
		DB:      r.db,
		Limit:   limit,
		Offset:  offset,
		OrderBy: "id ASC",
	}, &example)
	return
}
func (r *GormRepository) GetExamplesList() (exp []*models.Example, err error) {
	err = r.db.Database.Find(&exp).Error
	return
}

func (r *GormRepository) CreateExample(exp *models.Example) (err error) {
	err = r.db.Database.Create(exp).Error
	return
}
```
4. [controller](controllers) folder add a file `example_controller.go`
- Create API Endpoint 
- Use any syntax of GORM after `base.DB`, this is wrapper of `*gorm.DB`

```go
package controllers

import (
  "encoding/json"
  "go-fication/models"
  "go-fication/repository"
  "net/http"
  "strconv"
)

type ExampleHandler struct {
  repo repository.ExampleRepo
}

func NewExampleHandler(repo repository.ExampleRepo) *ExampleHandler {
  return &ExampleHandler{
    repo: repo,
  }
}
func (h *ExampleHandler) GetData(w http.ResponseWriter, request *http.Request) {
  q := request.URL.Query()
  limit, _ := strconv.Atoi(q.Get("limit"))
  offset, _ := strconv.Atoi(q.Get("offset"))

  data, err := h.repo.GetExamples(int64(limit), int64(offset))
  if err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    return
  }
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(&data)
}

func (h *ExampleHandler) CreateData(w http.ResponseWriter, request *http.Request) {
  example := new(models.Example)
  err := json.NewDecoder(request.Body).Decode(&example)
  if err != nil {
    http.Error(w, err.Error(), http.StatusBadRequest)
    return
  }
  err = h.repo.CreateExample(example)
  if err != nil {
    return
  }
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(example)
}
```
4. [routers](routers) folder add a file `example.go`
```go
package routers

import (
  "github.com/go-chi/chi/v5"
  "go-fication/controllers"
  "go-fication/infra/database"
  "go-fication/repository"
)

func ExamplesRoutes(router *chi.Mux, db *database.DB) {
  repo := repository.NewGormRepository(db)
  exampleCtrl := controllers.NewExampleHandler(repo)
  router.Group(func(r chi.Router) {
    r.Get("/test", exampleCtrl.GetData)
    r.Post("/test", exampleCtrl.CreateData)

  })
}

```
5. Finally, register routes to [index.go](routers/index.go)
```go
package routers

import (
  "github.com/go-chi/chi/v5"
  "net/http"
)

//RegisterRoutes add all routing list here automatically get main router
func RegisterRoutes(router *chi.Mux) {
  router.Get("/", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("\"live\": \"ok\""))
  })
  //Add All route
  ExamplesRoutes(router)
}
```
- Congratulation, your new endpoint `0.0.0.0:8000/v1/example/`
### Code Examples
- [go-fication-examples](https://github.com/akmamun/go-fication-examples)

### Deployment
#### Container Development Build
- Run `make build`

#### Container Production Build and Up
- Run `make production`

#### ENV Yaml Configure
```yaml
database:
  driver: "postgres"
  dbname: "test_pg_go"
  username: "mamun"
  password: "123"
  host: "postgres_db" # use `localhost` for local development
  port: "5432"
  ssl_mode: disable
  log_mode: false

server:
  host: "0.0.0.0"
  port: "8000"
  secret: "secret"
  allow_hosts: "localhost"
  debug: false #use `false` in production
  request:
    timeout: 100
```
- [Server Config](pkg/config/server.go)
```go
func ServerConfig() string {
viper.SetDefault("server.host", "0.0.0.0")
viper.SetDefault("server.port", "8000")
appServer := fmt.Sprintf("%s:%s", viper.GetString("server.host"), viper.GetString("server.port"))
return appServer
}
```
- [DB Config](pkg/config/db.go)
```go
func DbConfiguration() string {
	
dbname := viper.GetString("database.dbname")
username := viper.GetString("database.username")
password := viper.GetString("database.password")
host := viper.GetString("database.host")
port := viper.GetString("database.port")
sslMode := viper.GetString("database.ssl_mode")

dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
host, username, password, dbname, port, sslMode)
return dsn
}
```

### Useful Commands

- `make dev`: make dev for development work
- `make build`: make build container
- `make production`: docker production build and up
- `clean`: clean for all clear docker images

### Use Packages
- [Viper](https://github.com/spf13/viper) - Go configuration with fangs.
- [Gorm](https://github.com/go-gorm/gorm) - The fantastic ORM library for Golang
- [Logger](https://github.com/rs/zerolog) - Zero Allocation JSON Logger
- [Air](https://github.com/cosmtrek/air) - Live reload for Go apps (Docker Development)


## Migration with golang-migrate
## Download migrate to local directory
```bash
curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xz -- migrate
```

## Migrate (local)
```bash
./migrate -database ${MASTER_DB_URL} -path infra/database/migrations up
```

## Migrate (in debug container)
```bash
/usr/local/bin/migrate
```

## Create a new migration
```bash
./migrate create -ext sql -dir infra/database/migrations -seq name_of_migration
```
