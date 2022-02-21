FROM golang:1.17 as builder
RUN mkdir -p /go/src/
WORKDIR /go/src/
ADD . /go/src/
RUN CGO_ENABLE=0 GOOS=linux go build -o app cmd/app.go
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/app .
COPY ./web/ web/
RUN ls
EXPOSE 8080
CMD [ "./app" ]
