FROM golang:latest

RUN mkdir /app

WORKDIR /app

COPY hello-world.go /app

RUN go build -o hello-world

ENTRYPOINT ["./hello-world"]
