FROM golang:1.23.5-alpine

WORKDIR /app

COPY go.mod go.sum ./

COPY . .

RUN go build -o BOOKSTORE ./cmd/main/main.go

EXPOSE 8080

CMD [ "./BOOKSTORE" ]
