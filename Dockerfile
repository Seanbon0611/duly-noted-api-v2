FROM golang:latest

LABEL base.name="dulynoted"

WORKDIR /app

COPY . .

RUN go build -o main .


CMD ["app"]