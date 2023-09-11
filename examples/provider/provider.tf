terraform {
  required_providers {
    petstore = {
      source  = "Petstore/petstore"
      version = "0.1.0"
    }
  }
}

provider "petstore" {
  # Configuration options
}