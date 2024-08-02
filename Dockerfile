FROM golang:1.22-alpine3.20

RUN mkdir /app

ADD . /app

WORKDIR /app

ENV TZ="Asia/Tashkent"

RUN go build -o main ./cmd/app/main.go

RUN go mod tidy && go mod vendor

CMD ["./main"]