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