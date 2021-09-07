FROM golang:latest

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN GO111MODULE=on go build -v -o ./fizzbuzz *.go

EXPOSE 8080

CMD ["./fizzbuzz"]