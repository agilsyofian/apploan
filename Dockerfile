FROM golang:alpine

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

RUN rm -rf app.env
ADD dockerize.env app.env 

RUN go mod tidy

RUN go build -o main

EXPOSE 8080
ENTRYPOINT ["/app/main"]
