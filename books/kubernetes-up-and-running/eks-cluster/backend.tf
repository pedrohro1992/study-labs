terraform {
  backend "s3" {
    bucket         = "kubernetes-up-and-running-novembro"
    key            = "global/s3/terraform.tfstate"
    region         = "us-east-1"
    encrypt        = true
  }
}  
