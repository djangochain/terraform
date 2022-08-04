terraform {
  required_providers {
    example = {
      version = "~> 1.0.0"
      source  = "terraform-example.com/exampleprovider/example"
    }
  }
}

provider "example" {
  username = "Sahdev"
  password = "Garg"
}

resource "example_server" "my-server-name" {
	uuid_count = "1"
}