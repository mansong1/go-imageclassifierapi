FROM golang as build
EXPOSE 8080

RUN  mkdir -p /go/src \
  && mkdir -p /go/bin \
  && mkdir -p /go/pkg
ENV GOPATH=/go
ENV PATH=$GOPATH/bin:$PATH

RUN mkdir -p $GOPATH/src/app
ADD . $GOPATH/src/app

RUN go get -u google.golang.org/grpc \
github.com/julienschmidt/httprouter \
google.golang.org/protobuf

WORKDIR /go/src/app

RUN go build -o api .

CMD [ "/go/src/app/api" ]