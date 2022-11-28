# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased] - YYYY-MM-DD

### Changed

- In examples moved stark signer creation to separate function as not every workflow will need it. Also done to address [ISSUE-94](https://github.com/immutable/imx-core-sdk-golang/issues/94)

- IMX `Client` class now exposes public API methods for users who need access outside of basic workflows

### Fixed

- Deposit workflow for register and deposit NFT

## [v0.2.2] - 2022-11-11

### Fixed

- Bug fix to initialize the metadataRefreshesAPI object.

## [v0.2.1] - 2022-10-10

### Changed

- Added Goerli configuration.

## [v0.2.0] - 2022-10-10

### Changed

- Expose fewer public methods to make it easier for us to maintain the SDK.
- Introduced a [single entry point](./imx/interface.go#L47) for the SDK.
- Clear response and error types with IMXError struct getting all the req error details.
- Simplified complex types required for creating trades, orders and transfers, no more SignableToken in the interface
- Make the deposit method on SDK consistent with the API and industry norms.
- SDK now handles L1 signature-protected APIs; no need to generate imx-signature separately for authentication headers
- Updated signer key generation to be random so that there is no bias due to usage of l1 key.
- Removed un-necessary third-party dependencies.
- Examples provided with independent env files for each piece of work.
- Updated readme to use new interface.
- Added metadata refresh methods.

## [v0.1.0] - 2022-08-29

Initial release of Core SDK for Golang.

Includes a client for the public API and the following workflows:

- Register
- Mint
- Order
- Transfer
- Burn
- Deposit
- Withdrawals

Includes the examples for how to use the public API and each of the workflows.