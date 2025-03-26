package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/validatedeploy", ValidateLimitsWithName)
	http.HandleFunc("/containername", ValidateContainerName)

	cert := "/etc/certs/webhook/tls.crt"
	key := "/etc/certs/webhook/tls.key"
	httpsPort := ":9443"

	log.Fatalln(http.ListenAndServeTLS(httpsPort, cert, key, nil))
	// log.Fatalln(http.ListenAndServe(":8080", nil))

}
