GENERATED_CODE_DIR=generated

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
	cd $(GENERATED_CODE_DIR); go mod init $(GENERATED_CODE_DIR); \
    swagger generate client -f ../openapi.json && \
	go mod tidy

	