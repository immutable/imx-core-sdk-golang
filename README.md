<div align="center">
  <p align="center">
    <a  href="https://docs.x.immutable.com/docs">
      <img src="https://cdn.dribbble.com/users/1299339/screenshots/7133657/media/837237d447d36581ebd59ec36d30daea.gif" width="280"/>
    </a>
  </p>
</div>

---
# Immutable Core SDK Golang

The Immutable Core SDK Golang provides convenient access to the Immutable API's and Ethereum contract methods for applications written on the Immutable X platform.

Currently, our SDK supports interactions with our application-specific rollup based on StarkWare's StarkEx. In the future, we'll be adding StarkNet support across our platform.

## Documentation

See the [Developer guides](https://docs.x.immutable.com) for information on building on Immutable X.

See the [API reference documentation](https://docs.x.immutable.com/reference) for more information on our API's.

## Installation

TODO: revisit after package deployment

```sh
go get immutable.com/imx-core-sdk-golang 
```


## Usage

### Configuration

A configuration object is required to be passed into Core SDK requests. This can be obtained by using the `GetConfig` function available within the Core SDK. You are required to select the Ethereum network. The Immutable X platform currently supports `ropsten` for testing and `mainnet` for production.

```golang
import "immutable.com/imx-core-sdk-golang/config"

const alchemyAPIKey = "alchemy api key"

func main() {
   cfg := config.GetConfig(config.Ropsten, alchemyAPIKey)
   ...
}
```

#### L2 Signer

Some methods require an L2 signer as a parameter. The Core SDK expects you will generate your own L2 signer.

```golang
import (
   "immutable.com/imx-core-sdk-golang/signers/stark"
   ...
)

func main() {
   // L1 credentials
   l1signer := YourImplementationOfL1SignerInterface() // See examples/workflows/utils/signer.go 

   // L2 credentials
   // Obtain the stark signer associated with this user.
   l2signer, err := stark.GenerateStarkSigner(l1signer) // this is the sdk helper function
   if err != nil {
      ...
   }
}
```

### Standard API Requests

The Core SDK includes classes that interact with the Immutable X APIs.

```golang
// Standard API request example usage

import (
   "immutable.com/imx-core-sdk-golang/api/client"
   "immutable.com/imx-core-sdk-golang/api/client/assets"
   "immutable.com/imx-core-sdk-golang/api/models"
   "immutable.com/imx-core-sdk-golang/config"
)

func GetYourAsset(tokenID, tokenAddress string) (*models.Asset, error) {
   apiUrl := config.GetAPIURL(config.Ropsten)

   httpClient := client.NewHTTPClientWithConfig(nil, config.NewTransportConfig(&apiUrl))

   getAssetParams := assets.NewGetAssetParams()
   getAssetParams.SetTokenID(tokenID)
   getAssetParams.SetTokenAddress(tokenAddress)
   getAssetResponse, err := httpClient.Assets.GetAsset(getAssetParams)
   if err != nil {
      return nil, err
   }

   return getAssetResponse.GetPayload(), nil
}
```

View the [OpenAPI spec](openapi.json) for a full list of API requests available in the Core SDK.

### Authorised project owner requests

Some methods require authorisation by the project owner, which consists of a Unix epoch timestamp signed with your ETH key and included in the request header.

On project and collection methods that require authorisation, this signed timestamp string can typically be passed as the `IMXSignature` and `IMXTimestamp` parameters.

```golang
// Example method to generate authorisation headers
func getProjectOwnerAuthorisationHeaders(l1signer signers.L1Signer) (timestamp, signature string) {
   timestamp = strconv.FormatInt(time.Now().Unix(), 10)
   sigBytes, _ := l1signer.SignMessage(timestamp)
   signature = hexutil.Encode(sigBytes)
   return timestamp, signature
}

func CreateProject(l1signer signers.L1Signer, name, companyName, contactEmail string) (*models.CreateProjectResponse, error) {
   apiUrl := config.GetAPIURL(config.Ropsten)
   httpClient := client.NewHTTPClientWithConfig(nil, config.NewTransportConfig(&apiUrl))

   timestamp, signature := getProjectOwnerAuthorisationHeaders(l1signer)

   createProjectParams := projects.NewCreateProjectParams()
   createProjectParams.SetCreateProjectRequest(&models.CreateProjectRequest{
      CompanyName:  &companyName,
      ContactEmail: &contactEmail,
      Name:         &name,
   })
   createProjectParams.SetIMXTimestamp(timestamp)
   createProjectParams.SetIMXSignature(signature)
   createProjectResponse, err := httpClient.Projects.CreateProject(createProjectParams)
   if err != nil {
      return nil, err
   }

   return createProjectResponse.GetPayload(), nil
}
```

The following methods require project owner authorisation:

**Projects**

- createProject
- getProject
- getProjects

**Collections**

- createCollection
- updateCollection

**Metadata**

- addMetadataSchemaToCollection
- updateMetadataSchemaByName


### Contract Requests

Immutable X is built as a ZK-rollup in partnership with StarkWare. We chose the ZK-rollups because it is the only solution capable of scale without compromise. This means whenever you mint or trade an NFT on Immutable X, you pay zero gas, and the validity of all transactions are directly enforced by Ethereum’s security using zero-knowledge proofs -- the first “layer 2” for NFTs on Ethereum.

The Core SDK provides interfaces for all smart contracts required to interact with the Immutable X platform.

[See all smart contract available in the Core SDK](#smart-contract-autogeneration)

```golang
// This example is only to demonstrate using the generated smart contract clients
// We recommend using the workflows to deposit NFT
func DepositNft(l1signer signers.L1Signer, starkKey, assetType, vaultID, tokenID *big.Int) {
   cfg := config.GetConfig(config.Ropsten, "alchemy api key")
   client, err := ethclient.Dial(cfg.EthereumClientEndpoint)
   if err != nil {
      log.Panic(err)
   }
   defer client.Close()

   ethAddress := ethcommon.HexToAddress(l1signer.GetAddress())
   auth := &bind.TransactOpts{
      From: ethAddress,
      Signer: func(address ethcommon.Address, tx *types.Transaction) (*types.Transaction, error) {
         if address != ethAddress {
            return nil, bind.ErrNotAuthorized
         }
         return l1signer.SignTx(tx)
      },
      Context: context.Background(),
   }

   core, err := contracts.NewCore(ethcommon.HexToAddress(cfg.StarkContractAddress), client)
   transaction, err := core.DepositNft(auth, starkKey, assetType, vaultID, tokenID)
   log.Println("transaction hash:", transaction.Hash())
}
```

### Workflows

A workflow is a combination of API and contract calls required for more complicated functionality.

```golang
   // User registration workflow example
   cfg := config.GetConfig(config.Ropsten, alchemyAPIKey)
   apiClient := factories.NewAPIClient(&cfg)

   // Setup L1 signer
   l1signer, err := utils.NewBaseL1Signer(signerPrivateKey, chainID)
   if err != nil {
      log.Panicf("error in creating BaseL1Signer: %v", err)
   }

   // Setup L2 signer
   l2signer, err := stark.GenerateStarkSigner(l1signer)
   if err != nil {
      log.Panicf("error in creating StarkSigner: %v", err)
   }

   ctx := context.Background()
   response, err := registration.RegisterOffchain(ctx, apiClient, l1signer, l2signer, "user@email.com")
   if err != nil {
      log.Panicf("error in RegisterOffchain: %v", err)
   }
```

The workflow can be found in the [workflows directory](src/workflows/).

### Available Workflows

| Workflow                   | Description                                                                  |
| -------------------------- | ---------------------------------------------------------------------------- |
| `RegisterOffchain`         | Register L2 wallet.                                                          |
| `IsRegisteredOnchain`      | Check wallet registered on L1.                                               |
| `MintTokensWorkflow`       | Mint tokens on L2.                                                           |
| `CreateTransfer`           | Transfer tokens to another wallet.                                           |
| `CreateBatchTransfer`      | Batch transfer NFT tokens.                                                   |
| `Burn`                     | Burn tokens.                                                                 |
| `GetBurn`                  | Verify burn/transfer details.                                                |
| `Deposit`                  | Deposit based on token type. (ETH, ERC20, ERC721)                            |
| `PrepareWithdrawal`        | Prepare token for withdrawal.                                                |
| `CompleteWithdrawal`       | Withdraw to L1 based on token type. (ETH, ERC20, ERC721)                     |
| `CreateOrder`              | Create an order to sell an asset.                                            |
| `CancelOrder`              | Cancel an order.                                                             |
| `CreateTrade`              | Create a trade to buy an asset.                                              |

## Autogenerated Code

Parts of the Core SDK are automagically generated.

### API Autogenerated Code

We use OpenAPI (formally known as Swagger) to auto-generate the API clients that connect to the public APIs.
The OpenAPI spec is retrieved from https://api.x.immutable.com/openapi and also saved in the repo.

### Smart contract autogeneration

The Immutable solidity contracts can be found under `contracts` folder. Contract bindings in golang is generated using [abigen](https://geth.ethereum.org/docs/dapp/native-bindings#abigen-go-binding-generator).

#### Core

The Core contract is Immutable's main interface with the Ethereum blockchain, based on [StarkEx](https://docs.starkware.co/starkex-v4).

[View contract](contracts/Core.sol)

#### Registration

The Registration contract is a proxy smart contract for the Core contract that combines transactions related to onchain registration, deposits and withdrawals. When a user who is not registered onchain attempts to perform a deposit or a withdrawal, the Registration combines requests to the Core contract in order to register the user first. - users who are not registered onchain are not able to perform a deposit or withdrawal.

For example, instead of making subsequent transaction requests to the Core contract, i.e. `registerUser` and `depositNft`, a single transaction request can be made to the proxy Registration contract - `registerAndWithdrawNft`.

[View contract](contracts/Registration.sol)

#### IERC20

Standard interface for interacting with ERC20 contracts, taken from [OpenZeppelin](https://docs.openzeppelin.com/contracts/4.x/api/token/erc20#IERC20).

#### IERC721

Standard interface for interacting with ERC721 contracts, taken from [OpenZeppelin](https://docs.openzeppelin.com/contracts/4.x/api/token/erc721#IERC721).


## Getting Help

Immutable X is open to all to build on, with no approvals required. If you want to talk to us to learn more, or apply for developer grants, click below:

[Contact us](https://www.immutable.com/contact)

### Project Support

To get help from other developers, discuss ideas, and stay up-to-date on what's happening, become a part of our community on Discord.

[Join us on Discord](https://discord.gg/TkVumkJ9D6)

You can also join the conversation, connect with other projects, and ask questions in our Immutable X Discourse forum.

[Visit the forum](https://forum.immutable.com/)

#### Still need help?

You can also apply for marketing support for your project. Or, if you need help with an issue related to what you're building with Immutable X, click below to submit an issue. Select _I have a question_ or _issue related to building on Immutable X_ as your issue type.

[Contact support](https://support.immutable.com/hc/en-us/requests/new)
