package main

import (
	"log"
	"net/http"
)

const (
	httpsPort = ":8443"
	cert      = "/etc/certs/webhook/tls.crt"
	key       = "/etc/certs/webhook/tls.key"
)

func main() {
	http.HandleFunc("/validatedeploy", ValidateDeploymentName)

	cert := "/etc/certs/webhook/tls.crt"
	key := "/etc/certs/webhook/tls.key"
	log.Fatalln(http.ListenAndServeTLS(httpsPort, cert, key, nil))
}
