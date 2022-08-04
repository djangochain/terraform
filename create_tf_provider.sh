#/bin/bash

go mod tidy
go fmt
go build

cp terraform-provider-example ~/.terraform.d/plugins/terraform-example.com/exampleprovider/example/1.0.0/darwin_amd64