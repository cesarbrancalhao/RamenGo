# RamenGo
RamenGo is a food app built in [Gin](https://github.com/gin-gonic/gin).

## Table of Contents

1. [Getting Started](#start)
2. [Tech Stacks](#stacks)

### <a name="start">Getting Started</a>

Setup your docker container. Go to database/ and run:

```bash
docker compose up -d
```

Go to the .env file and insert your x-api-key key.

Run the development server:

```bash
go run api/cmd/main.go
```

Use [http://localhost:8080](http://localhost:8080) in your API tool to test.

### <a name="stacks">Stacks</a>

- [Gin](https://github.com/gin-gonic/gin)
- [Go](https://go.dev/)
- [GORM](https://gorm.io/index.html/)
- [Docker](https://www.docker.com/)
- [MySQL](https://www.mysql.com/)
- [JavaScript](https://www.javascript.com/)
- [CSS](https://developer.mozilla.org/en-US/docs/Web/CSS)