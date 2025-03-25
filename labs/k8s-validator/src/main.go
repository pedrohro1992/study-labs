package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/validatedeploy", ValidateDeploymentName)
	http.HandleFunc("/containername", ValidateContainerName)

	cert := "/etc/certs/webhook/tls.crt"
	key := "/etc/certs/webhook/tls.key"
	log.Fatalln(http.ListenAndServeTLS(":8443", cert, key, nil))
	// log.Fatalln(http.ListenAndServe(":8080", nil))
}
