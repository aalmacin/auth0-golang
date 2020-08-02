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
