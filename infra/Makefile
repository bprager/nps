# AWS profile (for PROD)
PROFILE ?= private
# TF := TF_LOG=DEBUG $(VAULT)/aws-vault --prompt=terminal exec $(PROFILE) -- $(TERRAFORM)/terraform
TF := aws-vault --prompt=terminal exec $(PROFILE) -- terraform
DEVIDENTIFIER := Development
PRODIDENTIFIER := Production
# Environment is DEV unless specified as PROD
ENV ?= DEV
ifeq ($(ENV),PROD)
	TF := aws-vault --prompt=terminal exec $(PROFILE) -- terraform
else
	TF := terraform
endif

.DEFAULT_GOAL := help

.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'

.PHONY: init
init:
	@$(TF) init
	@$(TF) workspace new $(DEVIDENTIFIER)
	@$(TF) workspace new $(PRODIDENTIFIER)

.PHONY: plan
plan:	## Create and display and execution plan for the infrastructure
	@$(TF) plan

.PHONY: up
up:	## Bring the infrastructure up und running and apply desired changes
	@$(TF) apply

.PHONY: down
down:	## Bring the infrastructure down
	@$(TF) destroy

.PHONY: dev
dev:	## Set "Development" terraform workspace
	@$(TF) workspace select $(DEVIDENTIFIER)

.PHONY: prod
prod:	## Set "Production" terraform workspace
	@$(TF) workspace select $(PRODIDENTIFIER)

.PHONY: show
show:	## Show current terraform workspace
	@$(TF) workspace show | awk 'BEGIN {FS = ":.*?## "}; {printf "The current environment is: \033[36m%-15s\033[0m %s\n", $$1, $$2}'
