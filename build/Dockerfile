FROM golang:1.22.1-alpine AS builder
WORKDIR /app
COPY cmd .
COPY go.mod .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:3.19.1
ARG VERSION
WORKDIR /root/
COPY --from=builder /app/main .
ENV SIMPLE_SERVICE_VERSION=${VERSION}
EXPOSE 8080
CMD ["./main"]
