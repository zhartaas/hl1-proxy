FROM golang:1.21-alpine

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o ./hl1-proxy ./cmd/web

CMD [ "./hl1-proxy" ]