
FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o app .


FROM alpine:3.20
WORKDIR /app
COPY --from=builder /app/app .
COPY --from=builder /app/static ./static
EXPOSE 8080
CMD ["./app"]