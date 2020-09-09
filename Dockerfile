FROM golang:1.14.4

# # All these steps will be cached
RUN mkdir /api
WORKDIR /api
# Get dependencies first
COPY go.mod .
COPY go.sum .
RUN go mod download

# Then copy all source, then build
COPY . .
RUN go build -o app main.go

# Then run
CMD ./app prod