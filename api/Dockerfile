FROM golang as build

WORKDIR /go/src/app

COPY vendor .

RUN go get github.com/julienschmidt/httprouter \
google.golang.org/grpc

COPY . .
RUN CGO_ENABLED=0 go build -ldflags "-s -w" -o 

FROM scratch
COPY --from=build /go/src/app /app
ENTRYPOINT ["/app"]