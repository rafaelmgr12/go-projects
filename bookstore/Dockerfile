FROM golang:1.18 as builder

COPY go.mod go.sum /go/src/github.com/rafaelmgr12/bookstore/
WORKDIR /go/src/github.com/rafaelmgr12/bookstore
RUN go mod download
COPY . /go/src/github.com/rafaelmgr12/bookstore
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/bookstore github.com/rafaelmgr12/bookstore/cmd/main


FROM alpine

RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder /go/src/github.com/rafaelmgr12/bookstore/build/bookstore /usr/bin/bookstore

EXPOSE 9010 9010

ENTRYPOINT ["/usr/bin/bookstore"]