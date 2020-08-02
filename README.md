# Auth0 - Golang Setup

## Create `go.mod` file

```sh
  touch go.mod
```

Update `go.mod` file

```go
  module github.com/aalmacin/auth0-golang

  go 1.23
```

## Add dependencies

```sh
  go get -d github.com/auth0/go-jwt-middleware
  go get -d github.com/dgrijalva/jwt-go
  go get -d github.com/codegangsta/negroni
  go get -d github.com/gorilla/mux
```

## Setup Middleware

### Setup `JwtMiddleware.go`

```go
package main

import (
	"net/http"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
)

func JwtMiddleware() http.HandlerFunc {
	jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {},
		SigningMethod:       jwt.SigningMethodRS256,
	})
}
```
