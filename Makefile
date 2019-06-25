# AWS profile (for PROD)
PROFILE ?= private
# TF := TF_LOG=DEBUG $(VAULT)/aws-vault --prompt=terminal exec $(PROFILE) -- $(TERRAFORM)/terraform
TF := aws-vault --prompt=terminal exec $(PROFILE) -- terraform

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
	$(TF) init

.PHONY: plan
plan:	## Create and display and execution plan for the infrastructure
	$(TF) plan

.PHONY: up
up:	## Bring the infrastructure up und running and apply desired changes
	$(TF) apply

.PHONY: down
down:	## Bring the infrastructure down
	$(TF) destroy
