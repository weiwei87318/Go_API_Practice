FROM golang

WORKDIR /api

COPY . .
RUN go mod download
RUN go build -o app
EXPOSE 8080
ENTRYPOINT ["./app"]

