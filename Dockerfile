FROM golang:alpine as builder
WORKDIR /app
COPY . .
CMD go build -o hello

FROM alpine
COPY --from=builder /app/hello .
CMD ./hello

