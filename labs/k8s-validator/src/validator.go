package main

import (
	"encoding/json"
	"errors"

	admissionv1 "k8s.io/api/admission/v1"
	appsv1 "k8s.io/api/apps/v1"
)

// validateDeploy validates the object is a deploy and can be unmarshalled

func validateDeploy(req *admissionv1.AdmissionReview) (*appsv1.Deployment, error) {
	if req.Request.Kind.Kind != "Deployment" {
		return nil, errors.New("apenas deploy sao suportados")
	}

	var d appsv1.Deployment
	if err := json.Unmarshal(req.Request.Object.Raw, &d); err != nil {
		return nil, errors.New("nao deu pra ler o deploy")
	}

	return &d, nil
}
