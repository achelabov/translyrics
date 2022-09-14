# Translyrics
The backend of a simple blogging platform built with golang.

## Features
- JWT Authentication
- MongoDB as a database
- Gin-gonic as a web framework

## API
### POST /auth/sign-up
Creates new user <br/>
**Request example**
```
{
    "username": "exampleUsername",
    "email": "example@gmail.com"
    "password": "examplePassword"
}
```
### POST /auth/sign-in
Request to get a JWT token by user credentials <br/>
**Request example**
```
{
    "username": "exampleUsername",
    "password": "examplePassword"
}
```
**Response example**
```
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InRlc3R1c2VyMSIsImVtYWlsIjoidGVzdGVtYWlsMSIsImV4cCI6MTY2MzE4MzYwN30.DWKK1YCYA-T9ayZqLiRnKcfumf0Idb8XGHtcSRTFIU4"
}
```
### GET /api/articles
Returns all articles <br/>
**Response example**
```
[
  {
    "id": "631f87caa64144284cd8a2d9",
    "userId": "5da2d8aae9b63715ddfae856",
    "title": "exampleTitle",
    "text": "exampleText"
  },
  ...
]
```
### POST /api/articles
> Required bearer token

Creates new article <br/>
**Request example**
```
{
  "title": "exampleTitle",
  "text": "exampleText"
}
```
### GET /api/articles/:id
Returns article by id <br/>
**Response example**
```
{
  "id": "631f87caa64144284cd8a2d9",
  "userId": "5da2d8aae9b63715ddfae856",
  "title": "exampleTitle",
  "text": "exampleText"
}
```
### PUT /api/articles/:id
> Required bearer token

Updates article by id <br/>
**Request example**
```
{
  "title": "exampleTitle",
  "text": "exampleText"
}
```
### DELETE /api/articles/:id
> Required bearer token

Deletes article by id