FROM golang:latest

LABEL base.name="dulynoted"

WORKDIR /app

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...


EXPOSE 3001

ENTRYPOINT [ "./main" ]
