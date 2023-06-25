FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go *.json ./

RUN CGO_ENABLED=0 GOOS=linux go build -tags dynamic .

EXPOSE 8080

CMD ["/farmsvc-go"]