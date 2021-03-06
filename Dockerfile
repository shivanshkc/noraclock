FROM golang:1.15-alpine as builder

# Create and change to the 'code' directory.
WORKDIR /code

# Build Application Native Binary.
# -mod=readonly ensures immutable go.mod and go.sum in container builds.
COPY . .
RUN GOOS=linux go build -tags musl -mod=readonly -v -o bin/application

FROM alpine:3

# Create and change to the the bin directory.
WORKDIR /usr/bin

# Copy the files to the production image from the builder stage.
COPY --from=builder /code/bin .

# Run the web service on container startup.
CMD ["application"]