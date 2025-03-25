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
  httpsPort := ":9443"


	log.Fatalln(http.ListenAndServeTLS(httpsPort, cert, key, nil))

}
