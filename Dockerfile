FROM golang:alpine

LABEL maintainer="Joseph Paz <josephpaz@martip07.me>"

RUN apk add --no-cache git

WORKDIR /go/src/app

COPY . .

RUN go get -d -v ./...

RUN GOOS=linux GOARCH=amd64 go build -o apigo main.go

EXPOSE 8088

CMD ["./apigo"]