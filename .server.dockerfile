FROM golang:1.21-alpine

RUN apk add --no-cache git

RUN mkdir /grpc-app
ADD ./server /grpc-app
WORKDIR /grpc-app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ./out/app ./server.go

EXPOSE 5001

CMD ["./out/app"]