FROM golang:latest
WORKDIR /app
ADD . .
RUN go build -o app cmd/app.go
EXPOSE 8080
CMD ["./app"]