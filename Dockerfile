
FROM widnyana/go-builder:1.22-alpine



LABEL maintainer="Selman TUNC <samtunc@yahoo.com>"
LABEL version="1.0"
LABEL project="Go and ReactJS"
LABEL description="This docker image is for handling web server requests."
ENV GOPATH /go
RUN mkdir -p "$GOPATH/src/reading" "$GOPATH/bin" && chmod -R 777 "$GOPATH"





WORKDIR /go/src/reading

RUN apk update \ 
    && apk add --no-cache \
    ca-certificates \
    curl \
    htop \
    nano \
    tzdata \
    && update-ca-certificates

RUN go install github.com/cosmtrek/air@latest

RUN pwd


# Make the source code path
RUN mkdir -p /go/src/reading

# Add all source code
ADD . /go/src/reading

RUN go mod tidy

# Run the Go installer
# RUN go install github.com/username/repository

# Indicate the binary as our entrypoint
ENTRYPOINT /go/bin/reading

# Expose your port
EXPOSE 3000

