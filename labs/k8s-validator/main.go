package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"k8s.io/api/admission/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	port = ":8080"
)

type Deploy struct {
	MetadataDeploy `json:"metadata"`
}

type MetadataDeploy struct {
	Name   string            `json:"name"`
	Labels map[string]string `json:"labels"`
}

func (m MetadataDeploy) isEmpty() bool {
	return m.Name == ""
}

func ValidateDeploymentName(w http.ResponseWriter, r *http.Request) {
	arReview := v1beta1.AdmissionReview{}
	if err := json.NewDecoder(r.Body).Decode(&arReview); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else if arReview.Request == nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		fmt.Println("Erro de Request")
		return
	}

	raw := arReview.Request.Object.Raw
	deploy := Deploy{}

	if err := json.Unmarshal(raw, &deploy); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else if deploy.MetadataDeploy.isEmpty() {
		http.Error(w, "invalid request body", http.StatusBadRequest)
	}

	arReview.Response = &v1beta1.AdmissionResponse{
		UID:     arReview.Request.UID,
		Allowed: true,
	}

	if len(deploy.MetadataDeploy.Labels) == 0 || deploy.MetadataDeploy.Labels["app"] != "nginx" {
		arReview.Response.Allowed = false
		arReview.Response.Result = &metav1.Status{
			Message: "Nome deve ser nginx",
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&arReview)
}

func main() {
	http.HandleFunc("/validatedeploy", ValidateDeploymentName)

	log.Fatal(http.ListenAndServe(port, nil))
}
