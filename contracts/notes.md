# Go Code Generation for Smart Contracts

Downloaded token relevant solidity files from https://github.com/OpenZeppelin/openzeppelin-contracts

Referenced https://goethereumbook.org/smart-contract-compile/

1. Generated ABI from solidity files. Run follwing command from `imx-core-sdk-golang/contracts` folder.
    ```
        solc --abi Registration.sol IERC20.sol IERC721.sol -o generated/abi --overwrite
        solc --bin Registration.sol IERC20.sol IERC721.sol -o generated/bin --overwrite
    ```
1. Convert ABI to go
    ```
        mkdir -p ../src/contracts
        abigen --bin generated/bin/Registration.bin --abi abi/Registration.abi --pkg=registration --out ../src/contracts/registration.go
        abigen --bin generated/bin/Core.bin --abi abi/Core.abi --pkg=core --out ../src/contracts/core.go
        abigen --bin generated/bin/IERC20.bin --abi abi/IERC20.abi --pkg=IERC20 --out ../src/contracts/IERC20.go
        abigen --bin generated/bin/IERC165.bin --abi abi/IERC165.abi --pkg=IERC165 --out ../src/contracts/IERC165.go
        abigen --bin generated/bin/IERC721.bin --abi abi/IERC721.abi --pkg=IERC721 --out ../src/contracts/IERC721.go
    ```
