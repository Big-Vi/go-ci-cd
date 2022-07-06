## Build
FROM golang:1.16-buster AS build

WORKDIR /app

# Download Go Modules
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy all go files using wildcard
COPY *.go ./

# build binary file
RUN go build -o /go-ci-cd


## Deploy
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /go-ci-cd /go-ci-cd

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/go-ci-cd"]