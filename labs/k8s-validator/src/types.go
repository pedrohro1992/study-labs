package main

type Deploy struct {
	Spec `json:"spec"`
}

type Spec struct {
	Template `json:"template"`
	Replicas int `json:"replicas"`
}

type Template struct {
	Metadata `json:"metadata"`
	SpecPod  `json:"spec"`
}

type SpecPod struct {
	Containers `json:"containers"`
}

type Containers []struct {
	Name      string `json:"name"`
	Resources `json:"resources"`
}

type Resources struct {
	Limits   `json:"limits"`
	Requests `json:"requests"`
}

type Limits struct {
	CPU    string `json:"cpu"`
	Memory string `json:"memory"`
}

type Requests struct {
	CPU    string `json:"cpu"`
	Memory string `json:"memory"`
}

type Metadata struct {
	Labels map[string]string `json:"labels"`
}

// isNull checks if the deployment do not have lavels
func (m Metadata) isNull() bool {
	return len(m.Labels) == 0
}
