package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)


type Response struct {
	StatusCode   int
	ErrorMessage string
	Data         interface{}
	Timestamp    time.Time
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!\n")
	fmt.Println("Endpoint hit: homePage\n")
}

func getRequestBody(r *http.Request) (interface{}, error) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Cannot read request body. Error: %s", err)
		return Response{
			Timestamp:    time.Now(),
			ErrorMessage: err.Error(),
			Data:         "",
			StatusCode:   http.StatusBadRequest,
		}, err
	}

	body := &Response{
		Timestamp:    time.Now(),
		ErrorMessage: "",
		Data:         "",
		StatusCode:   http.StatusBadRequest,
	}

	err = json.Unmarshal(requestBody, requestBody)
	if err != nil {
		log.Printf("Cannot unmarshal request body. Error: %s", err)
		return Response{
			Timestamp:    time.Now(),
			ErrorMessage: err.Error(),
			Data:         "",
			StatusCode:   http.StatusBadRequest,
		}, err
	}

	return *body, nil
}

func sendResponse(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonResponse, err := json.Marshal(data)
	if err != nil {
		log.Printf("Cannot marshal message. Error: %s", err)
		return err
	}
	json.NewEncoder(w).Encode(Response{
		Timestamp:    time.Now(),
		ErrorMessage: "",
		Data:         string(jsonResponse),
		StatusCode:   http.StatusOK,
	})
	return err
}
