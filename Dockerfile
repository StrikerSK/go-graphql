FROM golang:1.16.4-alpine3.13 as build

# allows app_env to be set during build (defaults to empty string)
ARG app_env

# sets an environment variable to app_env argument, this way the variable     will persist in the container for use in code
ENV APP_ENV $app_env

COPY . /src/go

WORKDIR /src/go

# install all dependencies
RUN go get ./...

# build the binary
RUN go build

# Put back once we have an application
CMD ["go-graphql"]

EXPOSE 5000