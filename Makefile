NAME = terraform-provider-convox
DESTDIR ?= ~/.terraform.d/plugins
build:
	go build -o terraform/terraform-provider-convox main.go

install:
	go build -o $(DESTDIR)/$(NAME) main.go
