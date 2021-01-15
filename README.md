# initialize module and download dependencies
- make sure you are in your GOPATH directory
- `$ go mod init github.com/onelzyugy/projects/learn-http`
- `$ go get github.com/gorilla/mux`

# build
- ```$ go build```
- ```$ ./learn-http``` to run
- http://localhost:8181/retrieve-todo

# dockerize 
- `$ docker build . -t learn-http`
- `$ docker images`
- `$ docker run -p 3000:3000 learn-http`