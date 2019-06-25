# VAULT := /Users/bernd/Projects/BerndsRepo/Go/bin
# TERRAFORM := /opt/local/bin
PROFILE ?= private
# TF := TF_LOG=DEBUG $(VAULT)/aws-vault --prompt=terminal exec $(PROFILE) -- $(TERRAFORM)/terraform
# TF := $(VAULT)/aws-vault --prompt=terminal exec $(PROFILE) -- $(TERRAFORM)/terraform
TF := aws-vault --prompt=terminal exec $(PROFILE) -- terraform

# Environment is DEV unless specified else
ENV ?= DEV
ifeq ($(ENV),PROD)
	TF := aws-vault --prompt=terminal exec $(PROFILE) -- terraform
else
	TF := terraform
endif

.DEFAULT_GOAL := plan

.PHONY: init
init:
	$(TF) init

.PHONY: plan
plan:
	$(TF) plan

.PHONY: apply
apply:
	$(TF) apply

.PHONY: destroy
destroy:
	$(TF) destroy
