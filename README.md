# RamenGo
RamenGo is a food app built in [Gin](https://github.com/gin-gonic/gin) for [Red Ventures'](https://redventures.com.br/) techincal test.

### Check it live [here](https://ramengo-front.vercel.app/)!

<div align="center">
    
![Alt text](./client/src/assets/prints/print-go.png?raw=true)

</div>

<br/>

## Table of Contents

1. [Getting Started](#start)
    - Client
    - Api
    - Cloud
2. [Tech Stacks](#stacks)
3. [Features](#features)

## <a name="start">Getting Started</a>
### Client:

<details>
<summary> How to run client </summary>
    
Install NPM packages.

```sh
npm i
```

Create a .env file on the root folder and insert your x-api-key. Check [.env.example](./.env.example) for examples.

*OBS: the client has placeholding objects to make sure it works if the server/connection is down, if you don't want this feature, you can go to [placeholders file](./client/src/scripts/config/placeholders.js) and comment/delete them.*
<br/>
*If you don't set an API_KEY on the .env file, it will use the placeholders by default*

Run the bundler.

```sh
npm run webpack
# If you dont have webpack installed you may need to run 'npm install -g webpack' first
```

I kept the Sass dist file for ease of use but you can start it with ```npm run sass```

</details>

### API:

<details>
<summary> How set up API </summary>

Install NPM packages.

```sh
npm i
```
    
Set up your docker container (you need to have Docker Desktop instlled).

```sh
npm run compose
```

After the initialization, migrate the models (this will insert the data too, change the script if you don't want it).

```sh
npm run migrate
```

Create a .env file **IN THE API FOLDER** and insert your x-api-key. Check [.env.api](./api/.env.example) for examples.
Run the development server:

```sh
npm run api
```

You can use [http://localhost:8080](http://localhost:8080) in your API tool to test the API locally.

</details>

### Cloud:

<details>
<summary> How use the webservice</summary>

Use [https://ramengo-back.onrender.com](https://ramengo-back.onrender.com) to access the API. 
<br/><br/>
**WARNING**: Due to Render's nature, instances can spin down and delay requests by 50 seconds or more, in case you're experiencing timeouts, try again a minute or two later.
<br/><br/>
*Note: Some functionalities are intentionally disabled for safety purposes. Use the software on localhost to access all features*

</details>

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
- [Vercel](https://vercel.com/)

## <a name="features">Features</a>

- **Mobile Responsible Design**: The webpages are designed to be visually pleasing and responsive, ensuring an optimal UX across any devices.

- **Clean Architecture**: Using modular and distinctive layers each with it's own responsibility, the application ensures maintainability and scalability, as commponents become flexible and easy to test.

- **API-First Development**:  The application is designed around RESTful principles and the API structure, which ensures a robust app, improving efficiency.

- **Environment Configuration**: Supports environment-specific configurations through environmental variables, allowing for flexible deployment across different environments without hardcoding sensitive information.

- **Continuous Integration/Deployment**: Sets up CI/CD pipelines using GitHub and Render, automating the build, test, and deployment processes for faster and safer releases.

- **Database Abstraction Layer**: Utilizes GORM library for abstracting database operations and simplifying CRUD operations while offering powerful query capabilities and migrations.

<br/>

![Alt text](./client/src/assets/prints/print-main.png?raw=true)
