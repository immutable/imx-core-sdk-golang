module immutable.com/imx-core-sdk-golang

go 1.18

replace immutable.com/imx-core-sdk-golang/api => ./generated/api

require github.com/ethereum/go-ethereum v1.10.19

require (
	github.com/btcsuite/btcd/btcec/v2 v2.2.0 // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.0.1 // indirect
	golang.org/x/crypto v0.0.0-20210921155107-089bfa567519 // indirect
	golang.org/x/sys v0.0.0-20211019181941-9d821ace8654 // indirect
	immutable.com/imx-core-sdk-golang/api v0.0.0-00010101000000-000000000000 // indirect
)
