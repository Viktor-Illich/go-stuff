FROM golang:latest

RUN mkdir /files
COPY kv_tcp.go /files
WORKDIR /files

RUN go build -o /files/kv_tcp kv_tcp.go
ENTRYPOINT ["/files/kv_tcp","80"]
