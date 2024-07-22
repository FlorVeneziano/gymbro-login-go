FROM golang:1.22
WORKDIR /gymbro-login
COPY /src/go.mod .
COPY /src/go.sum .
RUN /src/go mod download
COPY . .
RUN go build -o ./src/out/dist .
CMD ./src/out/dist
