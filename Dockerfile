# Use the official Golang image as the base image
FROM golang:1.23.5-alpine3.21
WORKDIR /url-shorter
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o url-shorter .
EXPOSE 8000
ENTRYPOINT [ "./url-shorter" ]    