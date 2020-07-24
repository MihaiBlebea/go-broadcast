local:
	export HTTP_PORT=8099 && \
	cd ./application && \
	go run .

# Docker scripts

build-up: build up

build:
	cd ./deploy/local-go-broadcast && docker-compose build --no-cache

up:
	cd ./deploy/local-go-broadcast && docker-compose up -d

down:
	cd ./deploy/local-go-broadcast && docker-compose down

push:
	cd ./deploy/local-go-broadcast && docker-compose push

# Terraform scripts

tf-init:
	cd ./deploy/terraform && terraform init

tf-plan:
	cd ./deploy/terraform && terraform plan

tf-apply:
	cd ./deploy/terraform && terraform apply

terraform: tf-init tf-plan tf-apply

