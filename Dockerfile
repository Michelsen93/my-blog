FROM golang:1.23.1-alpine3.20

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /my-blog

EXPOSE 8080

CMD ["/my-blog"]
