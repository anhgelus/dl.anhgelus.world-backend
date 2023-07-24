FROM golang:1.20-alpine

WORKDIR /app

COPY . .

RUN mkdir /data

RUN go build -o app .

CMD ./app

EXPOSE 80
