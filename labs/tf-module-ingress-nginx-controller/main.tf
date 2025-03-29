resource "helm_release" "external_ingress" {
  name             = var.ingress_controller_release_name
  namespace        = var.ingress_controller_release_namespace
  create_namespace = true
  chart            = "ingress-nginx"
  repository       = "https://kubernetes.github.io/ingress-nginx"
  version          = var.ingress_controller_version

  values = [
    templatefile("${path.module}/templates/values.yaml.tpl", {})
  ]
}
