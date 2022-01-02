FROM golang:latest
WORKDIR /app
ADD . .
RUN go build -o app
EXPOSE 3000
CMD [./app]