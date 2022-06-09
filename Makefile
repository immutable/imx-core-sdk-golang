.PHONY: generate-openapi-prod
generate-openapi-prod: get-openapi-prod generate-api

.PHONY: generate-openapi-ropsten
generate-openapi-ropsten: get-openapi-ropsten generate-api

.PHONY: get-openapi-prod
get-openapi-prod:
	rm openapi.json ; \
	curl -H "Accept: application/json+v3" \
		https://api.x.immutable.com/openapi \
		-o openapi.json

.PHONY: get-openapi-ropsten
get-openapi-ropsten:
	rm openapi.json ; \
	curl -H "Accept: application/json+v3" \
		https://api.ropsten.x.immutable.com/openapi \
		-o openapi.json

.PHONY: generate-api
generate-api:
	rm -rf src/api ; \
    mkdir -p src/api && \
	docker run --rm -v $(shell pwd):/app openapitools/openapi-generator-cli generate \
		-i ./app/openapi.json \
		-g go \
		--additional-properties=packageName=api,isGoSubmodule=true \
		-o /app/src/api && \
		awk '/^module/ { print "module immutable.com/imx-core-sdk/api"; next; }; {print;}' src/api/go.mod > go.mode && \
		mv go.mode src/api/go.mod
