FROM golang:1.15-alpine AS builder

ENV GO111MODULE=on

WORKDIR /src
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o httpserver

FROM scratch
COPY --from=builder /src/httpserver /app/
ENTRYPOINT ["/app/httpserver"]