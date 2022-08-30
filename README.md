<div align="center">
  <p align="center">
    <a  href="https://docs.x.immutable.com/docs">
      <img src="https://cdn.dribbble.com/users/1299339/screenshots/7133657/media/837237d447d36581ebd59ec36d30daea.gif" width="280"/>
    </a>
  </p>
</div>

---
**The Golang SDK interface is under active development and hasn't hit v1.0 yet.**

Its public interface shouldn't be considered final, and we may need to release breaking changes as we push towards v1.0.

---
# Immutable Core SDK Golang

The Immutable Core SDK Golang provides convenient access to the Immutable X API and Ethereum contract methods for applications written on the Immutable X platform.

Currently, our SDK supports interactions with our application-specific rollup based on StarkWare's [StarkEx](https://starkware.co/starkex/). In the future, we'll be adding [StarkNet](https://starknet.io/) support across our platform.

## Documentation

See the [Developer Home](https://docs.x.immutable.com/) for general information on building on Immutable X.

* Getting started, see [Developer Docs](https://docs.x.immutable.com/docs/welcome/)
* See the [API reference documentation](https://docs.x.immutable.com/reference) for more information on our API's.
* Reference docs related to this (Golang imx core) SDK can be found [here](https://docs.x.immutable.com/sdk-docs/core-sdk-golang/overview) 
## Installation

The supported go versions are 1.18 or above.

```sh
go get github.com/immutable/imx-core-sdk-golang 
```

## How to use

The following code snippets talk about how to get setup and use the various sections of the Golang SDK for most common use cases. Any sections that are not covered here please reach out to team. See [Getting Help](#getting-help)

### Configuration

A configuration object is provided with preset configurations for both Test and Production networks to simplify the environment setup required for which the Core SDK requests are made. This can be obtained by using the `GetConfig` function available within the `config` package of Core SDK. 

Select one of the following Ethereum networks Immutable X platform currently supports.

| Environment | Description   |  
|-------------|---------------|
| Sandbox     | Test network  |
| Mainnet     | Production    | 

```go
import "github.com/immutable/imx-core-sdk-golang/config"

const alchemyAPIKey = "alchemy api key"

func main() {
   cfg := config.GetConfig(config.Sandbox, alchemyAPIKey)
   ...
}
```

#### Ethereum Client

For information about how Ethereum client is setup, see `examples/workflows/main.go`

### How to generate the required signers

#### L1 Signer

Almost all the POST requests will need signed message. To sign a message as a minimum an L1 signer is required. An Ethereum wallet can be used to implement an L1 signer ([Getting started > Wallet](https://docs.x.immutable.com/docs/getting-started-guide/#wallet)).

When you implement an L1signer, it must satisfy [L1Signer interface](signers/signers.go). See [BaseL1Signer](examples/workflows/utils/signer.go) for a sample implementation of L1 Signer.

Also refer `examples/publicapi/list_assets/main.go` for environment setup examples.
#### L2 Signer

Some of the endpoints like Withdrawal, Orders, Trades, Transfers require an L2 signer. See `signers/stark` for information about generating your own L2 signer and also the following code snippet.

```go
import (
   "github.com/immutable/imx-core-sdk-golang/signers/stark"
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

Also see the examples for a sample l2 signer usage `examples/workflows/main.go`

### Standard API Requests

The Core SDK includes classes that interact with the Immutable X APIs.

```go
// Standard API request example usage

import (
    "context"
    "fmt"
    "github.com/immutable/imx-core-sdk-golang/generated/api"
    "github.com/immutable/imx-core-sdk-golang/config"
)

func GetYourAsset(tokenID, tokenAddress string) (*api.Asset, error) {
    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    ctx := context.WithValue(context.Background(), api.ContextServerIndex, config.Sandbox)

    getAssetResponse, httpResp, err := apiClient.AssetsApi.GetAsset(ctx, tokenAddress, tokenID).Execute()
    if err != nil {
        return nil, fmt.Errorf("error when calling `GetAsset`: %v, HTTP response body: %v", err, httpResp.Body)
    }

    return getAssetResponse, nil
}
```

View the [OpenAPI spec](openapi.json) for a full list of API requests available in the Core SDK.

### Authorised project owner requests

Some methods require authorisation by the project owner, which consists of a Unix epoch timestamp signed with your ETH key and included in the request header.

On project and collection methods that require authorisation, this signed timestamp string can typically be passed as the `IMXSignature` and `IMXTimestamp` parameters.
See [here](#how-to-generate-the-required-signers) for how to generate the signers required.

```go
// Example method to generate authorisation headers
func GetProjectOwnerAuthorisationHeaders(l1signer immutable.L1Signer) (timestamp, signature string, err error) {
    timestamp = strconv.FormatInt(time.Now().Unix(), 10)
    signedTimestamp, err := l1signer.SignMessage(timestamp)
    if err != nil {
        return "", "", err
    }
    signature = hexutil.Encode(signedTimestamp)
    return timestamp, signature, nil
}

func CreateProject(l1signer immutable.L1Signer, name, companyName, contactEmail string) (*api.CreateProjectResponse, error) {
    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    ctx := context.WithValue(context.Background(), api.ContextServerIndex, config.Sandbox)

    timestamp, signature, err := GetProjectOwnerAuthorisationHeaders(l1signer)
    if err != nil {
        return nil, err
    }

    createProjectRequest := api.NewCreateProjectRequest(companyName, contactEmail, name)
    createProjectResponse, httpResp, err := apiClient.ProjectsApi.CreateProject(ctx).
        CreateProjectRequest(*createProjectRequest).
        IMXTimestamp(timestamp).
        IMXSignature(signature).
        Execute()
    if err != nil {
        return nil, fmt.Errorf("error when calling `CreateProject`: %v, HTTP response body: %v", err, httpResp.Body)
    }

    return createProjectResponse, nil
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


### Contract requests

Immutable X is built as a ZK-rollup in partnership with StarkWare. We chose ZK-rollups because it is the only L2 scaling solution that has the same security guarantees as layer 1 Ethereum. The upshot of this is that you can mint or trade NFTs on Immutable X with zero gas costs whilst not compromising on security -- the first true “layer 2” for NFTs on Ethereum.

The Core SDK provides interfaces for all smart contracts required to interact with the Immutable X platform.

[See all smart contracts available in the Core SDK](#smart-contract-autogeneration)

```go
// This example is only to demonstrate using the generated smart contract clients
// We recommend using the workflows to deposit NFT
func DepositNft(l1signer immutable.L1Signer, starkKey, assetType, vaultID, tokenID *big.Int) {
   cfg := config.GetConfig(config.Sandbox, "alchemy api key")
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

```go
package onboarding

import (
    "context"
    "fmt"
    "math/big"

    "github.com/immutable/imx-core-sdk-golang/config"
    "github.com/immutable/imx-core-sdk-golang/examples/workflows/utils"
    "github.com/immutable/imx-core-sdk-golang/generated/api"
    "github.com/immutable/imx-core-sdk-golang/signers/stark"
    "github.com/immutable/imx-core-sdk-golang/workflows/registration"
)

func Register(signerPrivateKey string, chainID *big.Int) (*api.RegisterUserResponse, error) {
    // User registration workflow example
    configuration := api.NewConfiguration()
    apiClient := api.NewAPIClient(configuration)
    ctx := context.WithValue(context.Background(), api.ContextServerIndex, config.Sandbox)

    // Setup L1 signer
    l1signer, err := utils.NewBaseL1Signer(signerPrivateKey, chainID)
    if err != nil {
        return nil, fmt.Errorf("error in creating BaseL1Signer: %v", err)
    }

    // Setup L2 signer
    l2signer, err := stark.GenerateStarkSigner(l1signer)
    if err != nil {
        return nil, fmt.Errorf("error in creating StarkSigner: %v", err)
    }

    response, err := registration.RegisterOffchain(ctx, apiClient.UsersApi, l1signer, l2signer, "user@email.com")
    if err != nil {
        return nil, fmt.Errorf("error in RegisterOffchain: %v", err)
    }
    return response, nil
}
```

The workflows can be found in the [workflows directory](workflows/).
Sample usage of workflows can be found in [examples](examples/workflows).

### Available workflows

| Workflow                  | Description                                         |
|---------------------------|-----------------------------------------------------|
| `RegisterOffchain`        | Register L2 wallet.                                 |
| `IsRegisteredOnchain`     | Check wallet registered on L1.                      |
| `MintTokensWorkflow`      | Mint tokens on L2.                                  |
| `CreateTransfer`          | Transfer tokens to another wallet.                  |
| `CreateBatchTransfer`     | Batch transfer NFT tokens.                          |
| `Burn`                    | Burn tokens.                                        |
| `GetBurn`                 | Verify burn/transfer details.                       |
| `Deposit`                 | Deposit based on token type. (ETH, ERC20, ERC721)   |
| `PrepareWithdrawal`       | Prepare token (ETH, ERC20, ERC721) for withdrawal.  |
| `CompleteEthWithdrawal`   | Withdraw ETH to L1.                                 |
| `CompleteWithdrawal`      | Withdraw to L1 based on token type. (ERC20, ERC721) |
| `CreateOrder`             | Create an order to sell an asset.                   |
| `CancelOrder`             | Cancel an order.                                    |
| `CreateTrade`             | Create a trade to buy an asset.                     |

## Autogenerated code

Parts of the Core SDK are automagically generated.

### API autogenerated code

We use OpenAPI (formally known as Swagger) to auto-generate the API clients that connect to the public APIs.
The OpenAPI spec is retrieved from https://api.x.immutable.com/openapi and also saved in the repo.

### Smart contract autogeneration

The Immutable Solidity contracts can be found in the `contracts` folder. Contract bindings in Golang are generated using [abigen](https://geth.ethereum.org/docs/dapp/native-bindings#abigen-go-binding-generator).

#### Core

The Core contract is Immutable's main interface with the Ethereum blockchain, based on [StarkEx](https://docs.starkware.co/starkex-v4).

[View contract](contracts/Core.sol)

#### Registration

The Registration contract is a proxy smart contract for the Core contract that combines transactions related to onchain registration, deposits and withdrawals. When a user who is not registered onchain attempts to perform a deposit or a withdrawal, the Registration combines requests to the Core contract in order to register the user first. Users who are not registered onchain are not able to perform a deposit or withdrawal.

For example, instead of making subsequent transaction requests to the Core contract, i.e. `registerUser` and `depositNft`, a single transaction request can be made to the proxy Registration contract - `registerAndWithdrawNft`.

[View contract](contracts/Registration.sol)

#### IERC20

Standard interface for interacting with ERC20 contracts, taken from [OpenZeppelin](https://docs.openzeppelin.com/contracts/4.x/api/token/erc20#IERC20).

#### IERC721

Standard interface for interacting with ERC721 contracts, taken from [OpenZeppelin](https://docs.openzeppelin.com/contracts/4.x/api/token/erc721#IERC721).


## Getting help

Immutable X is open to all to build on, with no approvals required. If you want to talk to us to learn more, or apply for developer grants, click below:

[Contact us](https://www.immutable.com/contact)

### Project support

To get help from other developers, discuss ideas, and stay up-to-date on what's happening, become a part of our community on Discord.

[Join us on Discord](https://discord.gg/TkVumkJ9D6)

You can also join the conversation, connect with other projects, and ask questions in our Immutable X Discourse forum.

[Visit the forum](https://forum.immutable.com/)

#### Still need help?

You can also apply for marketing support for your project. Or, if you need help with an issue related to what you're building with Immutable X, click below to submit an issue. Select _I have a question_ or _issue related to building on Immutable X_ as your issue type.

[Contact support](https://support.immutable.com/hc/en-us/requests/new)


# Glossary

A lot of terminology here are specific to the Ethereum, see glossary page: https://goethereumbook.org/en/GLOSSARY.html