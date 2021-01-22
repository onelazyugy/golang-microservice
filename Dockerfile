FROM golang:latest as builder
RUN mkdir /app
WORKDIR /app
COPY . ./
RUN go mod download
# RUN make test
ARG version=dev
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags "-X main.version=$version" -o golang-microservice -v ./main.go

FROM alpine
EXPOSE 3000
ENV PORT=7777
COPY --from=builder /app/golang-microservice /