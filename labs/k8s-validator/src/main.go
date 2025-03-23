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

	// cert := "/etc/certs/webhook/tls.crt"
	// key := "/etc/certs/webhook/tls.key"
	// log.Fatalln(http.ListenAndServeTLS(":8443", cert, key, nil))
	log.Fatalln(http.ListenAndServe(port, nil))
}
