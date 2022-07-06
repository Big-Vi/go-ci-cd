## Build
FROM golang:1.18-alpine as development

WORKDIR /app

# Cache and install dependencies
COPY go.mod go.sum ./
RUN go mod download

RUN go install github.com/cosmtrek/air@latest
RUN air init

# Copy app files
COPY . .

EXPOSE 8000

CMD air
