# API Documentation

Check out the frontend documentation [here](./Frontend.md).

The API router is divided into three controller groups

- [Case](#case)

  - POST /case
  - GET /case/:id

- [Auth](#auth)

  - POST /auth/login
  - POST /auth/signUp

- [Admin](#admin)
  - GET /admin/getCase
  - GET /admin/solveCase
  - GET /admin/checkForNewCase

## Case

Routes for creating and getting information about a Case

### POST /case

Creates a case where the information about the case is part of the JSON body.

**Body :**

```json
{
  "user": {
    "name": "Dev Daksan",
    "email": "devd@gmail.com",
    "phone": "+91 84953 88953"
  },
  "car": {
    "color": "Red",
    "regNo": "TN 09 L 4508",
    "model": "Honda Civic",
    "lastSeen": "2019-06-02",
    "location": "Mandaveli",
    "description": "A huge dent on the front passenger door"
  }
}
```

**Response :**

Responds with `assigned: true` if there is a free officer to take the case and automatically assigns it to them.

```json
{
  "assigned": true,
  "code": 200,
  "id": "5ec9e65778b6cedc8d9dc68c",
  "message": "Successfully Added Case to DB",
  "officer": {
    "_id": "5ec9e0d978b6cedc8d9dc688",
    "userID": "5ec9e0d978b6cedc8d9dc689",
    "name": "William Tracy",
    "assigned": false
  }
}
```

Responds with `assigned: false` if there is no free officer to take the case.

```json
{
  "assigned": false,
  "code": 200,
  "id": "5ec9e65778b6cedc8d9dc68c",
  "message": "Successfully Added Case to DB",
  "officer": {
    "_id": "000000000000000000000000",
    "userID": "000000000000000000000000",
    "name": "",
    "assigned": false
  }
}
```

### GET /case/:id

Gets details about a case where `id` param is the ID of the case

```json
{
  "_id": "5ec9e65778b6cedc8d9dc68c",
  "officer": "5ec9e0d978b6cedc8d9dc688",
  "user": {
    "name": "Dev Daksan",
    "email": "devd@gmail.com",
    "phone": "+91 84953 88953"
  },
  "car": {
    "color": "Red",
    "regNo": "TN 09 L 4508",
    "model": "Honda",
    "lastSeen": "2019-06-02",
    "location": "Mandaveli",
    "image": "",
    "description": "A huge dent on the front passenger door"
  },
  "active": true,
  "assigned": false
}
```

## Auth

Used for Officer Login and Sign Up

### POST /auth/signUp

Creates an Officer account of successful authentication and also check for any unassigned cases and automatically assigns it to the officer.

**Body :**

```json
{
  "username": "will009",
  "password": "william1234",
  "name": "William Tracy",
  "email": "william.tracy@gmail.com",
  "secret": "secret"
}
```

**Response :**

Validates if the secret matches the API admin secret which is set as environment variable then creates the account.

```json
{
  "code": 200,
  "userID": "5ec9e0d978b6cedc8d9dc689"
}
```

If the secret does not match then it will throw an unauthorized error.

```json
{
  "code": 401,
  "message": "The entered admin secret is incorrect."
}
```

### POST /auth/login

Logs in the user if the provided credentials match the user in the DB.

**Body :**

Contains the username and password of the officer

```json
{
  "username": "will009",
  "password": "william1234"
}
```

**Response :**

The API uses JWT for authentication, so on successful login the response consists of the JWT Auth Token that needs to be a part of the authorization header while performing any admin operations.

```json
{
  "code": 200,
  "message": "Successfully authenticated user",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1OTAyOTU4NTIsInVzZXJJRCI6IjVlYzllMGQ5NzhiNmNlZGM4ZDlkYzY4OCJ9.I8-Lv-H8jni68T0u0M8AiM6tAekDCDk5soaRKvQOekE"
}
```

## Admin

Perform admin operations including getting getting information about a case and solving the case.
These routes are protected under JWT Auth. So in order to access these routes you will need to attach the JWT Auth Token, that was delivered during login, in the Authorization header as a Bearer Token.

### GET /admin/getCase

Gets details of the case that is currently assigned to the officer.

**Response :**

```json
{
  "case": {
    "_id": "5ec9e65778b6cedc8d9dc68c",
    "officer": "5ec9e0d978b6cedc8d9dc688",
    "user": {
      "name": "Dev Daksan",
      "email": "devd@gmail.com",
      "phone": "+91 84953 88953"
    },
    "car": {
      "color": "Red",
      "regNo": "TN 09 L 4508",
      "model": "Honda Civic",
      "lastSeen": "2019-06-02",
      "location": "Mandaveli",
      "image": "",
      "description": "A huge dent on the front passenger door"
    },
    "active": true,
    "assigned": true
  },
  "code": 200
}
```

Returns null if the officer as no case assigned at the moment

```json
{
  "case": null,
  "code": 200
}
```

### GET /admin/solveCase

Solves the case that is currently assigned to the officer. And then check if there are any unassigned cases and automatically assigns it to the officer.

**Response :**

```json
{
  "case": null,
  "code": 200,
  "message": "Successfully updated case"
}
```

If an unassigned case was found then it will return the details of the unassigned case.

```json
{
  "case": {
    "_id": "5ec9ec9d78b6cedc8d9dc68f",
    "officer": "5ec9e0d978b6cedc8d9dc688",
    "user": {
      "name": "Dev Daksan",
      "email": "devd@gmail.com",
      "phone": "+91 84953 88953"
    },
    "car": {
      "color": "Red",
      "regNo": "TN 09 L 4508",
      "model": "Honda Civic",
      "lastSeen": "2019-06-02",
      "location": "Mandaveli",
      "image": "",
      "description": "A huge dent on the front passenger door"
    },
    "active": true,
    "assigned": false
  },
  "code": 200,
  "message": "Successfully updated case"
}
```

### GET /admin/checkForNewCase

Checks for new case and automatically assigns them to the officer.

**Response :**

```json
{
  "case": {
    "_id": "5ec9e21478b6cedc8d9dc68b",
    "officer": "5ec9e0d978b6cedc8d9dc688",
    "user": {
      "name": "Dev Daksan",
      "email": "devd@gmail.com",
      "phone": "+91 84953 88953"
    },
    "car": {
      "color": "Red",
      "regNo": "TN 09 L 4508",
      "model": "Honda",
      "lastSeen": "2019-06-02",
      "location": "Mandaveli",
      "image": "",
      "description": "A huge dent on the front passenger door"
    },
    "active": true,
    "assigned": true
  },
  "code": 200
}
```

Returns with null if there no unassigned case in the DB

```json
{
  "case": null,
  "code": 200
}
```
