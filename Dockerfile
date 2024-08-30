# FROM golang:1.22.2 AS builder
# WORKDIR /app
# COPY . .
# RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server ./cmd

# FROM scratch
# EXPOSE 8000
# COPY --from=builder /app/cmd/server /server
# ENTRYPOINT [ "/server" ]

FROM golang:1.22.2 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server ./cmd

FROM alpine:latest
RUN apk --no-cache add ca-certificates bash
COPY --from=builder /app/server /server
COPY wait-for-it.sh /wait-for-it.sh
RUN chmod +x /wait-for-it.sh
EXPOSE 8000
ENTRYPOINT ["/wait-for-it.sh", "mkt-db:3306", "--", "/server"]
