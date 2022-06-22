FROM golang:latest
LABEL maintainer="Superm4n"
WORKDIR /app
COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 8080
ENTRYPOINT ["./main"]