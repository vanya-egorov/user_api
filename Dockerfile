FROM golang:1.23

WORKDIR /app

ENV GO111MODULE=on \
    CGO_ENABLED=0

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o server cmd/main.go

CMD ["./server"]