GENERATOR_TEMPLATES_DIR=generator-templates
GENERATED_CODE_DIR=api
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
	rm -rf $(GENERATED_CODE_DIR) && \
    mkdir -p $(GENERATED_CODE_DIR) && \
	docker run --rm -v $(shell pwd):/app openapitools/openapi-generator-cli generate \
		-i ./app/openapi.json \
		-c ./app/go-client-config.yaml \
		-t ./app/generator-templates/templates \
		-o /app/api
	rm -rf $(GENERATED_CODE_DIR)/go.mod $(GENERATED_CODE_DIR)/go.sum $(GENERATED_CODE_DIR)/git_push.sh

	