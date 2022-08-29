module github.com/immutable/imx-core-sdk-golang/examples/create_project

go 1.18

replace (
	github.com/immutable/imx-core-sdk-golang => ./../../../
	github.com/immutable/imx-core-sdk-golang/examples/workflows => ./../../../examples/workflows
)

require (
	github.com/ethereum/go-ethereum v1.10.23
	github.com/immutable/imx-core-sdk-golang v0.0.0-00010101000000-000000000000
	github.com/immutable/imx-core-sdk-golang/examples/workflows v0.0.0-00010101000000-000000000000
)

require (
	github.com/btcsuite/btcd/btcec/v2 v2.2.0 // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.0.1 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/shopspring/decimal v1.3.1 // indirect
	golang.org/x/crypto v0.0.0-20210921155107-089bfa567519 // indirect
	golang.org/x/net v0.0.0-20220624214902-1bab6f366d9e // indirect
	golang.org/x/oauth2 v0.0.0-20220808172628-8227340efae7 // indirect
	golang.org/x/sys v0.0.0-20220520151302-bc2c85ada10a // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
)
