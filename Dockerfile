# multi stage build to reduce size
FROM golang:1.13 as builder
WORKDIR /workspace
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download 

COPY main.go main.go 
COPY handlers/ handlers/
COPY services/ services/
COPY types/ types/  

RUN ls -al

RUN go env 
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -o golang-microservice main.go
FROM alpine
WORKDIR /
COPY --from=builder /workspace/golang-microservice .
ENTRYPOINT ["/golang-microservice"]