FROM golang:1.19-buster

ENV GOPATH=/

COPY ./ ./

RUN go mod download

RUN go build -o file ./cmd/main.go

CMD ["./file"]