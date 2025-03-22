package main

import (
	"log"
	"net/http"
)

const (
	port = ":8080"
)

func main() {
	http.HandleFunc("/validatedeploy", ValidateDeploymentName)

	log.Fatal(http.ListenAndServe(port, nil))
}
