# syntax=docker/dockerfile:1

FROM golang:1.19-alpine

WORKDIR /receiver

COPY go.mod ./
COPY go.sum ./

RUN go mod tidy

#RUN go get github.com/gin-gonic/gin

#RUN go mod download

COPY *.go ./

RUN go get .

RUN go build -o /docker-gs-ping

EXPOSE 8080

CMD [ "/docker-gs-ping" ]
