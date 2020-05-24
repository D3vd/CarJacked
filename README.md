# Car Jacked

[![CodeFactor](https://www.codefactor.io/repository/github/r3l3ntl3ss/carjacked/badge?s=c6707e33b8272259ab6ee93d5de857e164fb9b77)](https://www.codefactor.io/repository/github/r3l3ntl3ss/carjacked)

Application for reporting stolen cars

Frontend - [carjacked.now.sh](https://carjacked.now.sh)
Backend - [carjacked.herokuapp.com](https://carjacked.herokuapp.com/)

Check out the API Documentation [here](./Documentation/API.md).

## Stack

- Backend - Golang
- Frontend - React.js
- Database - MongoDB

## Running the App Locally

### Manual setup

Make sure to have MongoDB running at port `27017` or export custom mongoDB URI with `export MONGODB_URI=`

#### Download the source code and install all necessary packages

```bash
  go get github.com/R3l3ntl3ss/CarJacked

  # Move to project directory
  cd $GOPATH/src/github.com/R3l3ntl3ss/CarJacked

  # Install all dependencies
  go get -d -v ./...
  go install -v ./...
```

#### Start the Server

```bash
  # Export Secret for JWT Authentication and Admin Secret for Sign Up
  export SECRET=secret

  # Run the app
  go run .
```

#### Start the Client

Please refer the client [README.md](./client/README.md) to start the client server

### Using Docker Compose

The whole app can be ran locally by with a single command using docker compose as the all parts of the app are containerized.

```bash
  git clone https://github.com/R3l3ntl3ss/CarJacked
  docker-compose up
```

Go to `http://localhost:8000` to access the Frontend. The Backend will be hosted at `http://localhost:8080`

## Libraries and Packages Used

- HTTP Server - [Gin Gonic](https://github.com/gin-gonic/gin)
- MongoDB Driver - [mongodb-driver](https://github.com/mongodb/mongo-go-driver)
- JWT Token Auth - [dgrijalva/jwt-go](https://github/dgrijalva/jwt-go)
