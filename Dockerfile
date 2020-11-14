FROM golang:1.15.3-alpine

RUN mkdir -p /usr/service
WORKDIR /usr/service

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -mod=readonly -v -o bin/service

CMD ["bin/service"]