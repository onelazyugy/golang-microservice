# initialize module and download dependencies
- make sure you are in your GOPATH directory
- `$ go mod init github.com/onelzyugy/projects/golang-microservice`
- `$ go get github.com/gorilla/mux`

# build
- ```$ go build```
- ```$ ./golang-microservice``` to run
- http://localhost:8181/retrieve-todo

# dockerize 
- `$ docker build . -t golang-microservice`
- `$ docker images`
- `$ docker run -p {port}:{port} golang-microservice /golang-microservice`
- if you set ```ENV PORT={port #}``` in Dockerfile, then you need to run `docker run` cmd below
- `$ docker run -p {port #}:{port #} golang-microservice /golang-microservice`
- if you don't set a port number for ```ENV PORT={port #}``` in Dockerfile, then you need to run `docker run` cmd below
- `$ docker run -p 8080:8080 golang-microservice /golang-microservice` (8080 is hard coded in the main.go)

# sample api
- GET: http://localhost:8181/retrieve-todo
- POST: http://localhost:8181/add-todo
- `{
    "itemId": 1,
    "itemName": "my first item name",
    "description": "take out the trash"
}`

# examples
- https://thenewstack.io/make-a-restful-json-api-go/
- https://github.com/gorilla/mux

# Swagger doc
- https://goswagger.io/install.html to install swaggger cli
- `$ swagger generate spec -o ./swagger.yaml`
- `$ swagger serve ./swagger.yaml`
- `$ swagger serve -F swagger ./swagger.yaml`
- examples: 
- https://www.ribice.ba/swagger-golang/, 
- https://github.com/ribice/golang-swaggerui-example/blob/master/cmd/swagger/model.gohttps://github.com/ribice/golang-swaggerui-example/blob/master/cmd/swagger/model.go