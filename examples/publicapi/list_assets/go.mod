module immutable.com/imx-core-sdk-golang/examples/list_assets

go 1.18

replace immutable.com/imx-core-sdk-golang => ./../../../src/

replace immutable.com/imx-core-sdk-golang/api => ./../../../src/generated/api

require (
	github.com/joho/godotenv v1.4.0
	immutable.com/imx-core-sdk-golang v0.0.0-00010101000000-000000000000
	immutable.com/imx-core-sdk-golang/api v0.0.0-00010101000000-000000000000
)

require (
	github.com/golang/protobuf v1.4.3 // indirect
	github.com/google/go-cmp v0.5.2 // indirect
	golang.org/x/net v0.0.0-20220127200216-cd36cc0744dd // indirect
	golang.org/x/oauth2 v0.0.0-20210323180902-22b0adad7558 // indirect
	google.golang.org/appengine v1.6.6 // indirect
	google.golang.org/protobuf v1.25.0 // indirect
)
