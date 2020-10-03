FROM golang:1.7.3
WORKDIR /go/src/github.com/mischal/releasing-test/
RUN go get -u github.com/gorilla/mux
COPY main.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o hello-world .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /go/src/github.com/mischal/releasing-test/hello-world .
CMD ["./hello-world"]