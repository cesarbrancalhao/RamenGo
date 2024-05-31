# RamenGo
RamenGo is a food app built in [Gin](https://github.com/gin-gonic/gin) for [Red Ventures'](https://redventures.com.br/) techincal test.

## Table of Contents

1. [Getting Started](#start)
2. [Tech Stacks](#stacks)

### <a name="start">Getting Started</a>

Set up your docker container.

```bash
./scripts/compose.sh
```
After the initialization, migrate the models.

```bash
./scripts/migrate.sh
```

Create a .env file and insert your x-api-key. Check [.env.example](api/.env.example) for examples.

Run the development server:

```bash
./scripts/run.sh
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