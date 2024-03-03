FROM golang:1.21-alpine as build

COPY . /app

WORKDIR /app

RUN go mod tidy
RUN go build -o main .

FROM alpine:3.12

COPY --from=build /app/main /app
RUN chmod +x /app

CMD ["./app"]

