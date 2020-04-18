package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"google.golang.org/grpc"

	tf_core_framework "tensorflow/core/framework/tensor_go_proto"
	tf_core_framework_shape "tensorflow/core/framework/tensor_shape_go_proto"
	tf_core_framework_types "tensorflow/core/framework/types_go_proto"

	pb "tensorflow_serving/apis"
)

func main() {
	fmt.Printf("Hello, World!\n")

	r := httprouter.New()

	r.POST("/classify", classifyHandler)
	log.Fatal(http.ListenAndServe(":8080", r))
}

func classifyHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	fileUrl := "https://upload.wikimedia.org/wikipedia/commons/4/4c/Push_van_cat.jpg"
	resp, err := http.Get(fileUrl)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	imageBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	// // First Read image
	// imageFile, header, err := r.FormFile("image")

	// // Contains filename and extension
	// imageName := strings.Split(header.Filename, ".")
	// if err != nil {
	// 	responseError(w, "Could not read image", http.StatusBadRequest)
	// 	return
	// }
	// defer imageFile.Close()

	// var imageBuffer bytes.Buffer
	// // Copy image data to a buffer
	// io.Copy(&imageBuffer, imageFile)

	// Create Request to tfServer
	var tfServer string = "locahost:8500"

	request := &pb.PredictRequest{
		ModelSpec: &pb.ModelSpec{
			Name:          "resnet",
			SignatureName: "serving_default",
		},
		Inputs: map[string]*tf_core_framework.TensorProto{
			"image_bytes": &tf_core_framework.TensorProto{
				Dtype: tf_core_framework_types.DataType_DT_STRING,
				TensorShape: &tf_core_framework_shape.TensorShapeProto{
					Dim: []*tf_core_framework_shape.TensorShapeProto_Dim{
						&tf_core_framework_shape.TensorShapeProto_Dim{
							Size: int64(1),
						},
					},
				},
				StringVal: [][]byte{imageBytes},
			},
		},
	}

	// Connect to server
	conn, err := grpc.Dial(tfServer, grpc.WithInsecure())
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

	fmt.Printf("classes: %v", result.Outputs["classes"].Int64Val)
}
