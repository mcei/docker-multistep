FROM golang:alpine as compiler
WORKDIR /app
COPY ./main.go .
RUN go build main.go

FROM alpine
COPY --from=compiler /app/main .
CMD ./main

