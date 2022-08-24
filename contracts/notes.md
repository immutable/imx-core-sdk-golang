# Go Code Generation for Smart Contracts

Downloaded token relevant solidity files from https://github.com/OpenZeppelin/openzeppelin-contracts

Referenced https://goethereumbook.org/smart-contract-compile/

1. Generated ABI from solidity files. Run following command from `imx-core-sdk-golang/contracts` folder.
    ```
        solc --abi Registration.sol IERC20.sol IERC721.sol -o generated/abi --overwrite
        solc --bin Registration.sol IERC20.sol IERC721.sol -o generated/bin --overwrite
    ```
1. Convert ABI to go
    ```
        mkdir -p ../contracts
        abigen --bin generated/bin/Registration.bin --abi generated/abi/Registration.abi --pkg=registration --out ../generated/contracts/registration.go
        abigen --bin generated/bin/Core.bin --abi generated/abi/Core.abi --pkg=contracts --type=core --out ../generated/contracts/core.go
        abigen --bin generated/bin/IERC20.bin --abi generated/abi/IERC20.abi --pkg=IERC20 --out ../generated/contracts/IERC20.go
        abigen --bin generated/bin/IERC165.bin --abi generated/abi/IERC165.abi --pkg=IERC165 --out ../generated/contracts/IERC165.go
        abigen --bin generated/bin/IERC721.bin --abi generated/abi/IERC721.abi --pkg=IERC721 --out ../generated/contracts/IERC721.go
    ```
