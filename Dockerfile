FROM golang:latest
WORKDIR /app
ADD . .
RUN go build -o app
EXPOSE 8080
CMD [./app]