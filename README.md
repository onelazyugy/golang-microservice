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
- `$ docker run -p 3000:3000 golang-microservice`

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