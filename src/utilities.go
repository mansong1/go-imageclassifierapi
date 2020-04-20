package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage)
	router.HandleFunc("/classify", classifyHandler).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func responseError(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Accept", "application/json")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

func responseJson(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "appication/json")
	json.NewEncoder(w).Encode(data)
}

func getClassName() (map[int][]string, error) {
	// open imagnet_class_index file
	reader, err := os.Open("/app/conf/imagenet_class_index.json")
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	// read Json categories
	catJSON, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	// unmarshal into map of int to array of string
	var classes map[int][]string
	err = json.Unmarshal(catJSON, &classes)
	if err != nil {
		return nil, err
	}
	return classes, nil
}
