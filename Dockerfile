## We specify the base image we need for our
## go application
FROM golang:1.15-alpine
RUN mkdir /app
ADD . /app
WORKDIR /app
## Add this go mod download command to pull in any dependencies
RUN go mod download
## Our project will now successfully build with the necessary go libraries included.
RUN go build -o main ./cmd
## Our start command which kicks off
## our newly created binary executable
CMD ["/app/main"]