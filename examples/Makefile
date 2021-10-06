TERRAFORM_STUFF ?= .terraform.lock.hcl *.tfstate *.tfstate.*
TERRAFORM_FOLDER ?= .terraform/
TERRAFORM_VARS ?= main.tfvars

default: clean init plan

.PHONY: clean
clean:
	@rm ${TERRAFORM_STUFF} 2> /dev/null | true
	@rm -Rf ${TERRAFORM_FOLDER}

.PHONY: init
init:
	@terraform init

.PHONY: plan
plan:
	@terraform plan -var-file=${TERRAFORM_VARS}

.PHONY: apply
apply:
	@terraform apply -var-file=${TERRAFORM_VARS}
