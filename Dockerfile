FROM golang:latest

LABEL base.name="dulynoted"

WORKDIR /app

COPY . .

RUN go build -o main .

EXPOSE 3001

ENTRYPOINT [ "./main" ]