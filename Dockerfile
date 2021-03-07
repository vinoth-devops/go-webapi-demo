FROM golang:1.14 as builder
#set work directory and copy source code
WORKDIR /go/src/app
COPY ./web.go .
#Install gorilla mux and build go code
RUN go get -u github.com/gorilla/mux \
    && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

#create web image with go binary
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/app/app .
EXPOSE 8080
CMD ["./app"]
