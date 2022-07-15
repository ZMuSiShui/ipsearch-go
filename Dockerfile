FROM golang:1.17 as builder
WORKDIR /root/
ADD . /root/
RUN go mod tidy && \
    CGO_ENABLE=0 GOOS=linux go build -ldflags '-s -w --extldflags "-static -fpic"' -tags netgo -o app app.go
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /root/cmd/app .
COPY ./web/ web/
EXPOSE 8080
CMD [ "./app" ]
