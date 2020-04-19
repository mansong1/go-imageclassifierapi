FROM golang:alpine AS builder
RUN  mkdir -p /go/src \
  && mkdir -p /go/bin \
  && mkdir -p /go/pkg
ENV GOPATH=/go
ENV PATH=$GOPATH/bin:$PATH
RUN mkdir -p $GOPATH/src/app
ADD . $GOPATH/src/app
RUN apk update && \
apk add --no-cache git ca-certificates && \
update-ca-certificates
WORKDIR $GOPATH/src/app
RUN go get github.com/gorilla/mux \
google.golang.org/grpc
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o api .

FROM scratch
EXPOSE 8080
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/src/app /api/
WORKDIR /api
CMD ["./api"]