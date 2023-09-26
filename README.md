# Go Online Courses API

---
### Description

---
This project is a RESTful API of online courses website inspired by [kelas.programmerzamannow.com](https://kelas.programmerzamannow.com/). 
Here it is applied in a project "Go Online Courses API", which is made using Go-Lang and Gin Framework.
This project implements `Clean Code`, `Integration Testing`, and `API Documentation`.

### Libraries and Tools Used

---
1. GORM: [gorm.io/gorm](https://github.com/go-gorm/gorm)
2. Midtrans Payment Gateway: [github.com/veritrans/go-midtrans](https://github.com/veritrans/go-midtrans)
3. JWT Authentication: [github.com/golang-jwt/jwt/v4](https://github.com/golang-jwt/jwt)
4. Integration Testing: [github.com/stretchr/testify](https://github.com/stretchr/testify)
5. Request Input Validation: [github.com/go-playground/validator/v10](https://github.com/go-playground/validator)
6. FFmpeg: [gopkg.in/vansante/go-ffprobe.v2](https://github.com/vansante/go-ffprobe/tree/v2.1.1)
7. Password Hashing (Bcrypt): [golang.org/x/crypto](https://pkg.go.dev/golang.org/x/crypto)
8. HTTP Tunneling: [localtunnel.me](https://github.com/localtunnel/localtunnel)
> The full details can be seen on `./go.mod`

### Database Design

---
![DB_DESIGN](https://ik.imagekit.io/mazzlookman/go_pzn_restful_api_diagram.png?updatedAt=1695427800586)
Table Relationship:
* `users` table has `many to many` relationship with `courses` table.
* `authors` table has `one to many` relationship with `courses` table.
* `categories` table has `many to many` relationship with `courses` table.
* `courses` table has `one to many` relationship with `lesson_titles` table.
* `lesson_titles` table has `one to many` relationship with `lesson_contents` table.
* `transactions` table is `belongs to` `courses` table and `users` table.


### How To Run This Project ?

---
> Make sure your computer has `Git` and `Docker` are installed.

Please follow the steps below:

```
# Open the terminal, then clone this repository
git clone https://github.com/mazzlookman/go-online-courses-api

# Move to project
cd ./go-online-courses-api

# Run application
docker compose up -d

# note: it takes 20-50 seconds waiting for the server to connect to the database
# After that, you can start to hit the endpoints using http client (Postman, Insomnia, cURL, etc.)

# Stop application
docker compose down

# If you want to delete application image and volume
docker image rm goc-api
docker volume prune
```

### API Documentation

---

You can view the API Documentation in this repository at `./docs` directory. Enter into `./docs` directory, 
then there is API Documentation for each endpoint. In this project, it's using `openapi version 3.0.3` 
with `json` file format. If you want to get the UI display of this API Documentation, you must install the `swagger-ui plugin` for your IDE.
> Which I know:
> * Jetbrains: is automatically installed
> * VSCode: it's "OpenAPI (Swagger) Editor". 

### How To Contribute?

---
You can contribute by adding features or endpoints or creating the frontend. 
So, let's start fork this repository and feel free to pull request.


