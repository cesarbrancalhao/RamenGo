# RamenGo
RamenGo is a food app built in [Gin](https://github.com/gin-gonic/gin) for [Red Ventures'](https://redventures.com.br/) techincal test.

![Alt text](client/assets/images/capa.jpg?raw=true)

## Table of Contents

1. [Getting Started](#start)
2. [Tech Stacks](#stacks)


## <a name="start">Getting Started</a>
### Local:
Set up your docker container.

```bash
npm run compose
```

After the initialization, migrate the models.

```bash
npm run migrate
```

Create a .env file and insert your x-api-key. Check [.env.example](api/.env.example) for examples.

Run the development server:

```bash
npm run api
```

Use [http://localhost:8080](http://localhost:8080) in your API tool to test the API locally.


### Cloud (for Red Ventures team):
Use [this endpoint](https://ramengo-back.onrender.com) to access the API.
WARNING: due to Render's nature, instances can spin down and delay requests by 50 seconds or more, in case you're experiencing timeouts, try again a minute or two later.

## <a name="stacks">Stacks</a>

- [Gin](https://github.com/gin-gonic/gin)
- [Go](https://go.dev/)
- [GORM](https://gorm.io/index.html/)
- [Docker](https://www.docker.com/)
- [MySQL](https://www.mysql.com/)
- [JavaScript](https://www.javascript.com/)
- [CSS](https://developer.mozilla.org/en-US/docs/Web/CSS)
- [Aiven](https://console.aiven.io/)
- [Render](https://dashboard.render.com/)