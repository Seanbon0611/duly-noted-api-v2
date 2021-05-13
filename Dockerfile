FROM golang:latest

LABEL base.name="dulynoted"

WORKDIR /app

COPY go.mod .

COPY go.sum .

COPY . .

RUN go build

CMD ["./duly-noted-api-v2"]
