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
	@make -C infra/ init

.PHONY: plan
plan:	## Create and display the execution plan for the infrastructure
	@make -C infra/ plan

.PHONY: up
up:	## Bring the infrastructure up und running and apply desired changes
	@make -C infra/ up

.PHONY: down
down:	## Bring the infrastructure down
	@make -C infra/ down

.PHONY: dev
dev:	## Set the environment to "Development"
	@make -C infra/ dev

.PHONY: prod
prod:	## Set the environment to "Production"
	@make -C infra/ prod

.PHONY: show-env
show-env:	## Displays the currently set environment
	@make -C infra/ show