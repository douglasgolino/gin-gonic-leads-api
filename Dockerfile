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

FROM scratch
EXPOSE 8000
COPY --from=builder /app/server /server
ENTRYPOINT [ "/server" ]
