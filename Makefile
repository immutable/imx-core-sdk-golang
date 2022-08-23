GENERATOR_TEMPLATES_DIR=generator-templates
GENERATED_CODE_DIR=generated/api
CURRENT_DIR = $(shell pwd)

.PHONY: generate-openapi-prod
generate-openapi-prod: get-openapi-prod generate-api

.PHONY: generate-openapi-ropsten
generate-openapi-ropsten: get-openapi-ropsten generate-api

.PHONY: get-openapi-prod
get-openapi-prod:
	rm -f openapi.json && touch openapi.json && \
	curl -H "Accept: application/json+v3" \
    https://api.x.immutable.com/openapi \
    -o openapi.json

.PHONY: get-openapi-ropsten
get-openapi-ropsten:
	rm -f openapi.json && touch openapi.json && \
	curl -H "Accept: application/json+v3" \
    https://api.ropsten.x.immutable.com/openapi \
    -o openapi.json

.PHONY: generate-api
generate-api:
	rm -rf imx/$(GENERATED_CODE_DIR) && \
    mkdir -p imx/$(GENERATED_CODE_DIR) && \
	docker run --rm -v $(shell pwd):/app openapitools/openapi-generator-cli generate \
		-i ./app/openapi.json \
		-c ./app/go-client-config.yaml \
		-t ./app/generator-templates/templates \
		-o /app/imx/generated/api
	rm -rf imx/$(GENERATED_CODE_DIR)/go.mod imx/$(GENERATED_CODE_DIR)/go.sum imx/$(GENERATED_CODE_DIR)/git_push.sh

	