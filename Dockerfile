FROM golang:1.22
WORKDIR /gymbro-login
COPY src/go.mod .
COPY src/go.sum .
RUN go mod download
COPY src/ .
RUN go build -o ./out/dist .
CMD ["./out/dist"]
