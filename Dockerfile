FROM golang:1.22-alpine3.19 AS builder
RUN apk add --no-cache git upx
WORKDIR /app
COPY ["go.mod", "go.sum", "./"]
RUN go mod download -x
COPY . .
RUN go build -o app -v ./cmd
RUN upx app

#final stage
FROM alpine:latest
LABEL Name=back-qr-code Version=0.0.2
RUN apk update && apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /app .
ENTRYPOINT ["./app"]
