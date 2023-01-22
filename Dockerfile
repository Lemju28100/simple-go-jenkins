ARG GO_VERSION="1.19.5"
ARG BASE_IMAGE="golang:${GO_VERSION}-alpine"

FROM ${BASE_IMAGE} AS builder

    # Change working directory to /go/src/app
    WORKDIR /go/src/app

    # Copy the current directory contents into the container at /go/src/app
    COPY . .

    # Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed
    RUN go mod download

    # Build the Go app
    RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM scratch

    WORKDIR /app

    # Copy the Pre-built binary file from the previous stage
    COPY --from=builder /go/src/app/app ./app

    COPY --from=builder /go/src/app/index.html index.html
    COPY --from=builder /go/src/app/assets assets
    COPY --from=builder /go/src/app/.env .env

    # Expose port 8080 to the outside world
    EXPOSE 3000

    CMD ["./app"]