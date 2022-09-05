#!/bin/sh

# Go Code Generation for Smart Contracts
# Downloaded token relevant solidity files from https://github.com/OpenZeppelin/openzeppelin-contracts
# Referenced https://goethereumbook.org/smart-contract-compile/

# Outpiut directory for the generated contracts go files.
OUT_DIR="../imx/internal/contracts"

ABI_TEMP_DIR="generated/abi"
BIN_TEMP_DIR="generated/abi"

# Generating ABI from solidity files
solc --abi Registration.sol IERC20.sol IERC721.sol -o $ABI_TEMP_DIR --overwrite
solc --bin Registration.sol IERC20.sol IERC721.sol -o $BIN_TEMP_DIR --overwrite

# Convert ABI to go
mkdir -p $OUT_DIR
abigen --bin $BIN_TEMP_DIR/Registration.bin --abi $ABI_TEMP_DIR/Registration.abi --pkg=contracts --type=registration --out $OUT_DIR/registration.go
abigen --bin $BIN_TEMP_DIR/Core.bin         --abi $ABI_TEMP_DIR/Core.abi         --pkg=contracts --type=core         --out $OUT_DIR/core.go
abigen --bin $BIN_TEMP_DIR/IERC20.bin       --abi $ABI_TEMP_DIR/IERC20.abi       --pkg=contracts --type=IERC20       --out $OUT_DIR/IERC20.go
abigen --bin $BIN_TEMP_DIR/IERC721.bin      --abi $ABI_TEMP_DIR/IERC721.abi      --pkg=contracts --type=IERC721      --out $OUT_DIR/IERC721.go