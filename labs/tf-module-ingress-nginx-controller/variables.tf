variable "ingress_controller_release_name" {
  type        = string
  description = "Set the release name for the ingress nginx installation"
  default     = "nginx-ingress-controller"
}

variable "ingress_controller_release_namespace" {
  type        = string
  description = "Set the release namespace for the ingress nginx installation"
  default     = "nginx-ingress-controller"
}

variable "ingress_controller_version" {
  type        = string
  description = "Set the version of helm chart"
  default     = "4.5.2"
}
