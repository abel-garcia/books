# books
Challenge 

Arquitecture project [DDD](https://en.wikipedia.org/wiki/Domain-driven_design)

![image](https://bozdag.dev/common/onion-architecture.png)


set up environment variables
```bash
export BOOKS_HTTPSERVER=8989
export BOOKS_STORAGE_USER=books
export BOOKS_STORAGE_PASSWORD=123456
export BOOKS_STORAGE_HOST=db
export BOOKS_STORAGE_PORT=5432
export BOOKS_STORAGE_NAME=books
```

build image docker compose
```bash
docker compose build
```

run containers
```bash 
docker compose up -d
```

API documentation on:
```http://localhost:8989/swagger/index.html```

Note: BearerAuth required

example test:
```bash
go test ./...
```