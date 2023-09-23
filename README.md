# Go Online Courses API

---
### Description

---
This is a RESTful API of `go-online-courses-api` project, 
created using: GoLang, Gin Framework, MySQL, JWT Authentication,
Midtrans Payment Gateway, FFmpeg, as well as implement `clean architecture` and `integration test`.

### Features

---
1. API Documentation
2. Authentication and Authorization using JWT
3. Midtrans Payment Gateway
4. Integration Testing

### Database Design

---
![DB_DESIGN](https://ik.imagekit.io/mazzlookman/go_pzn_restful_api_diagram.png?updatedAt=1695427800586)

### How To Run This Project ?

---
> Make sure your computer has `Git` and `Docker` are installed.

Please follow the steps below:

```
# Open the terminal, then clone this repository
git clone https://github.com/mazzlookman/go-online-courses-api

# Move to project
cd ./go-online-courses-api

# Pulling all the required docker images and create the container and network
docker compose create

# Run the application
docker compose start

# note: it takes 30-50 seconds waiting for the server to connect to the database
# After that, you can start to hit the endpoints using http client (Postman, Insomnia, cURL, etc.)

# Stop the application
docker compose down
```

### API Documentation

---

You can view the API Documentation at `/apidocs` directory. Enter the `/apidocs` directory, 
then there is API Documentation for each endpoint. In this project, it's using `openapi version 3.0.3` 
with `json` file format. If you want to get the UI display of this API Documentation, you must install the `swagger-ui plugin` for your IDE.
> Which I know:
> * Jetbrains: is automatically installed
> * VSCode: it's "OpenAPI (Swagger) Editor". 

### Tool Used

---
* All libraries listed on `go.mod`


