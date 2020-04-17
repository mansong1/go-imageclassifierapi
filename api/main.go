package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"google.golang.org/grpc"

	pb "tensorflow_serving/apis"
	tf_core_framework "github.com/tensorflow/tensorflow/tensorflow/go/core/framework"
)

type ClassifyResult struct {
	FileName string `json: "Filename"`
	Label    string `json:labels`
}

func main() {
	fmt.Printf("Hello, World!\n")

	r := httprouter.New()
	r.POST("/classify", classifyHandler)
	log.Fatal(http.ListenAndServe(":8080", r))
}

func classifyHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// First Read image
	imageFile, header, err := r.FormFile("image")

	// Contains filename and extension
	imageName := strings.Split(header.Filename, ".")
	if err != nil {
		responseError(w, "Could not read image", http.StatusBadRequest)
		return
	}
	defer imageFile.Close()

	var imageBuffer bytes.Buffer
	// Copy image data to a buffer
	io.Copy(&imageBuffer, imageFile)

	// Create Request to tfServer
	var tfServer string = "locahost:8500"

	request := &pb.PredictRequest{
		ModelSpec: &pb.ModelSpec{
			Name:          "resnet",
			SignatureName: "serving_default",
		},
		Inputs: map[string]*tf_core_framework.TensorProto{
			"image_bytes": &tf_core_framework.TensorProto{
				Dtype: tf_core_framework.DataType_DT_STRING,
				TensorShape: &tf_core_framework.TensorShapeProto{
					Dim: []*tf_core_framework.TensorShapeProto_Dim{
						&tf_core_framework.TensorShapeProto_Dim{
							Size: int64(1),
						},
					},
				},
				StringVal: [][]byte{imageBytes},
			},
		},
	}

	// Connect to server
	conn, err := grpc.Dial(*tfServer, grpc.WithInsecure())
	if err != nil {
		responseError(w, "Cannot connect to tfSever", http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	stub := pb.NewPredictionServiceClient(conn)

	result, err := stub.Predict(context.Background(), request)
	if err != nil {
		responseError(w, "Could not run prediction", http.StatusInternalServerError)
	}

	responseJson(w, ClassifyResult {
		FileName: header.FileName, Label: resp.Outputs["classes"].Int64Val
	})
}
