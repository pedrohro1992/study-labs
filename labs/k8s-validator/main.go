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
	Template       `json:"template"`
}

type MetadataDeploy struct {
	Name   string            `json:"name"`
	Labels map[string]string `json:"labels"`
}

type Template struct {
	Spec Spec `json:"spec"`
}

type Spec struct {
	Containers
}

type Containers struct {
	NameContainer string `json:"name"`
}

func (c Containers) isNull() bool {
	return c.NameContainer == ""
}

func ValidateDeploymentName(w http.ResponseWriter, r *http.Request) {
	arReview := v1beta1.AdmissionReview{}
	if err := json.NewDecoder(r.Body).Decode(&arReview); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else if arReview.Request == nil {
		http.Error(w, "invalid request body 1234", http.StatusBadRequest)
		fmt.Println("Erro de Request")
		return
	}

	raw := arReview.Request.Object.Raw
	deploy := Deploy{}

	if err := json.Unmarshal(raw, &deploy); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else if deploy.Spec.isNull() {
		http.Error(w, "invalid request body 666tapa", http.StatusBadRequest)
	}

	arReview.Response = &v1beta1.AdmissionResponse{
		UID:     arReview.Request.UID,
		Allowed: true,
	}

	if len(deploy.MetadataDeploy.Labels) == 0 || deploy.Spec.NameContainer != "nginx" {
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
