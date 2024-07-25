FROM golang:1.22-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

WORKDIR /app/cmd

RUN go build -o /app/main .

RUN ls -l /app

WORKDIR /app

EXPOSE 8080

CMD ["/app/main"]
