package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
)

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
}

func responseError(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

func responseJson(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "appication/json")
	json.NewEncoder(w).Encode(data)
}
