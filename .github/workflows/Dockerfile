FROM golang:1.22
WORKDIR /gymbro-login
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o ./out/dist .
CMD ./out/dist
