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
# After that, try to hit the endpoints in http client (Postman, Insomnia, cURL, etc.)

# Stop the application
docker compose down
```

### Tool Used

---
* All libraries listed in `go.mod`


