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
    mkdir -p src/$(GENERATED_CODE_DIR) && \
	SIGIL_DELIMS={{{,}}} gliderlabs-sigil -f ${GENERATOR_TEMPLATES_DIR}/parameter_template.gotmpl GOLANG_SDK_VERSION=${VERSION_STR} > ${GENERATOR_TEMPLATES_DIR}/template/client/parameter.gotmpl; \
	cd src/$(GENERATED_CODE_DIR); go mod init immutable.com/imx-core-sdk-golang/api; \
    swagger generate client -f $(CURRENT_DIR)/openapi.json -T $(CURRENT_DIR)/generator-templates/template && \
	go mod tidy

	