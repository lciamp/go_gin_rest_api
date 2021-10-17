FROM golang:1.17.2-alpine3.14

MAINTAINER lou ciampanelli <l.ciamp@me.com>

WORKDIR /app

# download go modules
COPY go.mod .
COPY go.sum .
RUN go mod download

# copy main.go
COPY *.go ./

# build
RUN go build -o /docker-main

# open port
EXPOSE 8080

# run
CMD [ "/docker-main" ]

# docker build -t docker-f1 .
# docker run -d docker-f1:latest