FROM golang:1.21.6-alpine AS builder
WORKDIR /app
COPY go.mod .
COPY cmd/ .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:3.19.1
ARG VERSION
WORKDIR /root/
COPY --from=builder /app/main .
ENV SIMPLE_SERVICE_VERSION=${VERSION}
EXPOSE 8080
CMD ["./main"]
