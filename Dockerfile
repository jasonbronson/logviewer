FROM golang:1.16-alpine

RUN apk update && apk add git bash build-base curl

RUN mkdir -p /app
WORKDIR /app

ADD . /app

ENV GO111MODULE=on

#development only
RUN ["go", "get", "github.com/githubnemo/CompileDaemon"]

#production
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o ./dist/api /app/*.go

#cron job
#RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o ./dist/cron /app/transport/cron.go

EXPOSE 8000

CMD ["/dist/api"]
