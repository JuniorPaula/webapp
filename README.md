# Web Application - Login and Image Upload

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)
![MySQL](https://img.shields.io/badge/mysql-%2300f.svg?style=for-the-badge&logo=mysql&logoColor=white)

This is a project of a simple web application in Golang that aims to test the login component, database integration, and image upload functionality. The application uses the `mux` package for route handling, MySQL with Docker and Docker Compose to set up the database, and the `Session` library for storing user data states.

## Key Features

-   **Login**: The application includes a login system that allows users to authenticate with their credentials and access protected resources.
-   **Image Upload**: Users can upload images to the application, which are then stored on the server.
- **API REST**: The application include an api rest to login and authentication the users.

## Integration Testing for Image Upload

A significant emphasis has been placed on integration testing for the image upload functionality. Integration tests ensure that the image upload process functions correctly, from user interaction to the proper storage of the image on the server.

Integration tests can be executed using Golang's standard testing library. Make sure you have your development environment properly configured.

## Dependencies

-   Golang
-   Docker
-   Docker Compose

Make sure you have Golang, Docker, and Docker Compose installed on your machine before running the application.

## Running the Application

1.  Clone this repository to your local machine.
2.  Navigate to the project's root directory.
3.  Run the command `docker-compose up -d` to start the MySQL container.
4.  Ensure that the MySQL container is running correctly.
5.  Run the command `go run ./cmd/web` to start the web application.
6.  Access the application in your browser at `http://localhost:8000`.

## Testing Image Upload

Image upload is a key functionality of the application. To test this feature, follow these steps:

1.  Access the application in your browser.
2.  Authenticate yourself with your user credentials.
3.  Find the option to upload images in the application interface.
4.  Select an image from your computer and upload it.
5.  Verify that the image has been successfully uploaded and is visible in the application.

## Tests
### Unit tests
to run the unit test:
`$ cd ./cmd/web` 
`$ go test -v .`
```
=== RUN   TestForm_Has
--- PASS: TestForm_Has (0.00s)
=== RUN   TestForm_Require
--- PASS: TestForm_Require (0.00s)
=== RUN   TestForm_Check
--- PASS: TestForm_Check (0.00s)
=== RUN   TestForm_ErrorGet
--- PASS: TestForm_ErrorGet (0.00s)
=== RUN   Test_application_handlers
--- PASS: Test_application_handlers (0.03s)
=== RUN   TestAppHome
--- PASS: TestAppHome (0.00s)
=== RUN   TestApp_renderWithBadTemplate
--- PASS: TestApp_renderWithBadTemplate (0.00s)
=== RUN   Test_app_Login
--- PASS: Test_app_Login (1.87s)
=== RUN   Test_app_UploadFiles
--- PASS: Test_app_UploadFiles (0.00s)
=== RUN   Test_app_UploadProfilePic
--- PASS: Test_app_UploadProfilePic (0.00s)
=== RUN   Test_application_addIPToContext
    middleware_test.go:37: 192.0.2.1
    middleware_test.go:37: unknown
    middleware_test.go:37: 192.3.2.1
    middleware_test.go:37: hello
--- PASS: Test_application_addIPToContext (0.00s)
=== RUN   Test_application_ipFromContext
--- PASS: Test_application_ipFromContext (0.00s)
=== RUN   Test_app_auth
--- PASS: Test_app_auth (0.00s)
=== RUN   Test_application_routes
--- PASS: Test_application_routes (0.00s)
PASS
ok      webapp/cmd/web  1.918s
```

### Integration tests
To run integration tests
`$ cd pkg/repository/dbrepo`
`$ go test -v -tags=integration .`
```
=== RUN   Test_pingDB
--- PASS: Test_pingDB (0.00s)
=== RUN   TestPostgresDBRepoInsertUser
--- PASS: TestPostgresDBRepoInsertUser (0.27s)
=== RUN   TestPostgresDBRepoAllUsers
--- PASS: TestPostgresDBRepoAllUsers (0.26s)
=== RUN   TestPostgresDbRepoGetUser
--- PASS: TestPostgresDbRepoGetUser (0.00s)
=== RUN   TestPostgresDbRepoGetUserByEmail
--- PASS: TestPostgresDbRepoGetUserByEmail (0.00s)
=== RUN   TestPostgresDBRepoUpdateUser
--- PASS: TestPostgresDBRepoUpdateUser (0.00s)
=== RUN   TestPostgresDbRepoDeleteUser
--- PASS: TestPostgresDbRepoDeleteUser (0.00s)
=== RUN   TestPostgresDBRepoResetPassword
--- PASS: TestPostgresDBRepoResetPassword (0.52s)
=== RUN   TestPostgresDBRepoInsertUserImage
--- PASS: TestPostgresDBRepoInsertUserImage (0.00s)
PASS
ok      webapp/pkg/repository/dbrepo 
```

## API Rest
this is a feature that was added later. consists of a rest api to authenticate and keep the user logged in through the refresh token

### Quicky start
```
$ go run ./cmd/api
```

### Routes
**WEB**
authentication
`[POST] /web/auth`
refresh token
`[GET] /web/refresh_token`
logout
`[GET] /web/logout`

**Authentication**
auth
`[POST] /auth`
refresh token
`[POST] /refresh-token`

**Users**
get all
`[GET] /users/`
get by id
`[GET] /users/{id}`
delete
`[DELETE] /users/{id}`
insert
`[PUT] /users/`
update
`[PATCH] /users/`

### Integration tests
```
$ cd ./cmd/api
$ go test . -v
```
```
=== RUN   Test_app_authenticate
--- PASS: Test_app_authenticate (2.05s)
=== RUN   Test_app_refresh
--- PASS: Test_app_refresh (0.00s)
=== RUN   Test_app_userHandlers
--- PASS: Test_app_userHandlers (0.00s)
=== RUN   Test_app_refreshUsingCookie
--- PASS: Test_app_refreshUsingCookie (0.00s)
=== RUN   Test_app_deleteRefreshCookie
--- PASS: Test_app_deleteRefreshCookie (0.00s)
=== RUN   Test_app_enableCORS
--- PASS: Test_app_enableCORS (0.00s)
=== RUN   Test_app_authRequire
--- PASS: Test_app_authRequire (0.00s)
=== RUN   Test_app_routes
--- PASS: Test_app_routes (0.00s)
=== RUN   Test_app_getTokenFromHeaderAndVerify
--- PASS: Test_app_getTokenFromHeaderAndVerify (0.00s)
PASS
ok      webapp/cmd/api  (cached)
```