package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	admissionv1 "k8s.io/api/admission/v1"
)

// extractAdmission extrac the admission request
func extractAdmission(r *http.Request) (*admissionv1.AdmissionReview, error) {
	if r.Header.Get("Content-Type") != "application/json" {
		return nil, fmt.Errorf("incorrect content type %q - Should be %s", r.Header.Get("Content-Type"), "application/json")
	}

	buff := new(bytes.Buffer)
	buff.ReadFrom(r.Body)
	body := buff.Bytes()

	if len(body) == 0 {
		return nil, fmt.Errorf("body is empty")
	}

	var adm admissionv1.AdmissionReview
	if err := json.Unmarshal(body, &adm); err != nil {
		return nil, fmt.Errorf("couldn't read the admission review request: %s", err)
	}

	fmt.Println(adm)

	if adm.Request == nil {
		return nil, fmt.Errorf("admission request is nil")
	}

	return &adm, nil
}
