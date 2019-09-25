FROM golang:1.13.0-alpine3.10 as builder

# Create directories
RUN mkdir /go/src/app
ADD . /go/src/app
WORKDIR /go/src/app

# Install go dep and install dependencies
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure

EXPOSE 8000

# Build the app
RUN go build -o /src/main .
CMD ["./main"]