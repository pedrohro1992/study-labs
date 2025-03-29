package main

import (
	"encoding/json"
	"errors"

	admissionv1 "k8s.io/api/admission/v1"
	appsv1 "k8s.io/api/apps/v1"
)

// checkContainerName checks witch container has the same name of the deployment
func checkContainerName(d appsv1.Deployment) int {
	deploymentName := d.Name
	var nameCounter int
	if len(d.Spec.Template.Spec.Containers) == 1 {
		nameCounter = 0
	} else {
		for i := 0; i < len(d.Spec.Template.Spec.Containers); i++ {
			if d.Spec.Template.Spec.Containers[i].Name == deploymentName {
				nameCounter = i
				break
			} else {
				nameCounter = 0
			}
		}
	}
	return nameCounter
}

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
