package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	admissionv1 "k8s.io/api/admission/v1"
	"k8s.io/api/admission/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func ValidateDeploymentTag(w http.ResponseWriter, r *http.Request) {
	req, err := extractAdmission(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Read the deploy data
	d, err := validateDeploy(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var resp *admissionv1.AdmissionReview
	cpu, err := d.Spec.Template.Spec.Containers[0].Resources.Limits.Cpu().AsInt64()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if cpu == 2 {
	}
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
	} else if deploy.Spec.Template.Metadata.isNull() {
		fmt.Println(&deploy.Spec.Template)
		http.Error(w, "Deployment sem label", http.StatusBadRequest)
	}

	arReview.Response = &v1beta1.AdmissionResponse{
		UID:     arReview.Request.UID,
		Allowed: true,
	}

	// if deploy.Replicas == 3 {
	// 	arReview.Response.Allowed = false
	// 	arReview.Response.Result = &metav1.Status{
	// 		Message: "Nome deve ser nginx",
	// 	}
	cpu, err := strconv.Atoi(deploy.Spec.Template.SpecPod.Containers[0].Limits.CPU)
	if err != nil {
		return
	}
	if cpu < 3 {
		arReview.Response.Allowed = false
		arReview.Response.Result = &metav1.Status{
			Message: "CPU deve ser maior que 3",
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&arReview)
}
