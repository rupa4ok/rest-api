# Rest Api go gin + gorm
It is a just simple RESTful API with Go using:
1. **Gin Framework**
2. **Gorm**

This Gin GORM API Boilerplate with MVC paradigm.

## Installation & Run

Setting DB in main.go
```go
make init
```

After build: localhost:8080

Db connection: .env

Default server port in dockerfile: ENV PORT 8080, external port in docker-compose.yml.

## Structure
```
├── docker
│   └── app 
│       └── Dockerfile
├── src
│    ├── Models
│    │   ├── Book.go // Book models
│    |	 ├── Scheme.go // Book struct and tabel
│    ├── Config
│    │   └── Database.go // Global DB
│    ├── Controllers
│    │   └── Book.go // Book Controller
│    ├── ApiHelpers
│    │   └── Response.go // response function
│    └── Routers
│        └── Routers.go // Routers
└── main.go
```

## API

#### /book
* `GET` : Get all book
* `POST` : Create a new book

#### /book/:id
* `GET` : Get a book
* `PUT` : Update a book
* `DELETE` : Delete a book

#Post Params
```
{
	"author": "Op Super John Doe Bilw",
	"name": "Implementation Golang",
	"category": "Knowledge"
}
