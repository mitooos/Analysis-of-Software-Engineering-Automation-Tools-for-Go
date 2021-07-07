



FROM golang:1.15.8-alpine3.13

WORKDIR /app

ENV GO111MODULE=on 

COPY . .

RUN go build src/main.go

EXPOSE [8080]

CMD ["./main"]

