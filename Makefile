DOCKER_PATH := ./deploy/local-blog
TERRAFORM_PATH := ./deploy/terraform
HELM_PATH := ./deploy/broadcast

local:
	export HTTP_PORT=8099 && \
	cd ./blog && \
	go run .

bundle:
	cd ./blog && \
	make bundle

open:
	open http://localhost:8088

# Docker scripts

build-up: bundle build up open

build:
	cd $(DOCKER_PATH) && docker-compose build --no-cache

up:
	cd $(DOCKER_PATH) && docker-compose up -d

down:
	cd $(DOCKER_PATH) && docker-compose down

push:
	cd $(DOCKER_PATH) && docker-compose push

# Terraform scripts

tf-init:
	cd $(TERRAFORM_PATH) && terraform init

tf-plan:
	cd $(TERRAFORM_PATH) && terraform plan

tf-apply:
	cd $(TERRAFORM_PATH) && terraform apply -auto-approve

tf-destroy:
	cd $(TERRAFORM_PATH) && terraform destroy

tf-get:
	cd $(TERRAFORM_PATH) && terraform get

terraform: tf-init tf-plan tf-apply

# Helm scripts

helm-debug:
	cd $(HELM_PATH) && helm install --generate-name --debug --dry-run .

# Kubernetes scripts

manual-cronjob: 
	kubectl create job --from=cronjob/broadcast broadcast-test-02
