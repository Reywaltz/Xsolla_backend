FROM golang:latest

WORKDIR /go/src/github.com/Reywaltz/xsolla_backend

COPY . .

RUN go mod download

RUN go build -o ./bin/ ./cmd/item-api/main.go

CMD [ "./bin/main" ]