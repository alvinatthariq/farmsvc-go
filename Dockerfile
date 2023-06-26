FROM golang:1.17-alpine

ENV GOPATH /go

RUN mkdir -p "$GOPATH/src/github.com/alvinatthariq/farmsvc-go" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

ADD . ${GOPATH}/src/github.com/alvinatthariq/farmsvc-go/

WORKDIR ${GOPATH}/src/github.com/alvinatthariq/farmsvc-go

COPY go.mod go.sum ./

RUN go get ./...

COPY *.go *.json ./

RUN apk update && apk add --no-cache git

RUN CGO_ENABLED=0 GOOS=linux go build -tags dynamic -o farmsvc-go

EXPOSE 8080

ENTRYPOINT ["/go/src/github.com/alvinatthariq/farmsvc-go/farmsvc-go"]