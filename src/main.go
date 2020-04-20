package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"google.golang.org/grpc"

	tf_core_framework "tensorflow/core/framework/tensor_go_proto"
	tf_core_framework_shape "tensorflow/core/framework/tensor_shape_go_proto"
	tf_core_framework_types "tensorflow/core/framework/types_go_proto"

	pb "tensorflow_serving/apis"
)

type payload struct {
	URL string `json:"URL"`
}

type ClassifyResult struct {
	Label string `json:"label"`
}

var tfServer string = "tfserver:8500"

func main() {
	handleRequests()
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func classifyHandler(w http.ResponseWriter, r *http.Request) {

	var imgUrl payload
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Enter image URL")
	}

	json.Unmarshal(reqBody, &imgUrl)
	w.WriteHeader(http.StatusCreated)

	resp, err := http.Get(imgUrl.URL)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	imageBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

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
	log.Printf("Connecting to tfServer at %v", tfServer)
	if err != nil {
		log.Fatalf("Cannot connect to server %v", err)
		responseError(w, "Cannot connect to tfSever", http.StatusInternalServerError)
	}
	defer conn.Close()

	stub := pb.NewPredictionServiceClient(conn)

	result, err := stub.Predict(context.Background(), request)
	if err != nil {
		responseError(w, "Could not run prediction", http.StatusInternalServerError)
	}

	result_class := result.Outputs["classes"].Int64Val[0]
	predictidx := int(result_class) - 1 // return int

	// get the classes
	classes, err := getClassName()
	if err != nil {
		log.Printf("Error: %s", err)
	}

	log.Printf("\nclasses: %v", predictidx)
	log.Printf("\nProbably %v", classes[predictidx][1])
	responseJson(w, ClassifyResult{Label: classes[predictidx][1]})
}
