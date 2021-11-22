package handler

import (
	"context"
	"encoding/json"
	"fmt"
	calculator "go_server/service"
	"go_server/service/dto"
	"log"
	"net/http"
)

func Fib(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var data dto.FibRequest
	err := decoder.Decode(&data)
	if err != nil {
		log.Fatalf("decoding error: %v", err)
	}
	res, _ := calculator.Fib(context.Background(), &data)
	fmt.Fprintf(w, "%d", res)
}

func Multiple(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var data dto.MultipleRequest
	err := decoder.Decode(&data)
	if err != nil {
		log.Fatalf("decoding error: %v", err)
	}

	res, _ := calculator.Multiple(context.Background(), &data)
	fmt.Fprintf(w, "%d", res)
}

func Add(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var data dto.AddRequest
	err := decoder.Decode(&data)
	if err != nil {
		log.Fatalf("decoding error: %v", err)
	}

	res, _ := calculator.Add(context.Background(), &data)
	fmt.Fprintf(w, "%d", res)
}
