
.terraform:
	terraform get -update=true

init:
	terraform init
plan: .terraform
	terraform plan -out plan

apply:
	terraform apply "plan"

destroy:
	terraform destroy
.PHONY: plan apply init destroy .terraform
