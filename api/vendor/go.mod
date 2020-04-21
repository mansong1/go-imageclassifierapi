module github.com/figroc/tensorflow-serving-client/v2

go 1.12

require (
	github.com/golang/protobuf v1.4.0-rc.4.0.20200313231945-b860323f09d0
	google.golang.org/grpc v1.26.0
	google.golang.org/protobuf v1.21.0
)

replace tensorflow_serving => ./tensorflow_serving

replace tensorflow => ./tensorflow
