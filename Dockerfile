FROM golang:1.15.3-alpine

RUN mkdir -p /usr/service
WORKDIR /usr/service

COPY . .
RUN go build -o bin/service ./main.go

CMD ["bin/service"]