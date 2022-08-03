GENERATOR_TEMPLATES_DIR=generator-templates
GENERATED_CODE_DIR=generated/api
CURRENT_DIR = $(shell pwd)

# Todo: Auto generate this version string from the release git tags.
VERSION_STR="0.1.0"

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
	rm -rf src/$(GENERATED_CODE_DIR) && \
    rm -f ${GENERATOR_TEMPLATES_DIR}/templates/client.mustache && \
	mkdir -p src/$(GENERATED_CODE_DIR) && \
	SIGIL_DELIMS={{{{,}}}} gliderlabs-sigil -f ${GENERATOR_TEMPLATES_DIR}/client_template.mustache GOLANG_SDK_VERSION=${VERSION_STR} > ${GENERATOR_TEMPLATES_DIR}/templates/client.mustache; \
	docker run --rm -v $(shell pwd):/app openapitools/openapi-generator-cli generate \
		-i ./app/openapi.json \
		-c ./app/go-client-config.yaml \
		-t ./app/generator-templates/templates \
		-o /app/src/generated/api

	


	