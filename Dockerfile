FROM golang:1.14.4

COPY . .
RUN go build -o app main.go
CMD ./app