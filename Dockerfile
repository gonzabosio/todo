FROM golang:1.22

WORKDIR /app

COPY go.* /app/

RUN go mod download

COPY . .

RUN go build -o main main.go

EXPOSE 3000

CMD [ "./main" ]