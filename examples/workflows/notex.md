	signerPrivateKey string = "c383bf49c2a18236f2da3247d39303e866c87eff868a8b148d0552e554b9e18d" // #1
    signerPrivateKey string = "4b56f5ce0ebaf6665a2712b4b0cbcafe42e7a0661091d67776547870a3e3d1a5" // #2
	alchemyApiKey    string = "DvukuyBzEK-JyP6zp1NVeNVYLJCrzjp_"


	tokenAddress = "0x4b089f5006778fef128506427235937d60b74dfb"
	tokenId      = "2"

contractAddress = "0x4b089f5006778fef128506427235937d60b74dfb"

# Created Contract

## 1

Deploying Contract with the account: 0x99c41E6AC42c909A228D558d0081C43234d6Ad21
Account Balance: 268184574637291576
Deployed Contract Address: 0xefb053d2967CCd58227F982d80d2c795D28b7759

## 2 

No need to generate any newer typings.
Deploying Contract with the account: 0x4db6Bf7AbafaAE73d9010610166660032D6c796A
Account Balance: 10486425281991929286
Deployed Contract Address: 0x4B089f5006778FeF128506427235937D60B74DFB
Verifying contract in 5 minutes...
Nothing to compile
No need to generate any newer typings.
Successfully submitted source code for contract
contracts/Asset.sol:Asset at 0x4B089f5006778FeF128506427235937D60B74DFB
for verification on Etherscan. Waiting for verification result...

# Created Project

## 1

INFO 2022-07-07T17:42:03+10:00 Created project with ID: 65460 | component=[IMX-CREATE-PROJECT] extra=undefined

## 2
INFO 2022-07-08T17:09:30+10:00 Creating project... | component=[IMX-CREATE-PROJECT] extra=undefined
REQUEST:
 POST https://api.ropsten.x.immutable.comv1projects
RESPONSE:
 201 Created
INFO 2022-07-08T17:09:31+10:00 Created project with ID: 65790 | component=[IMX-CREATE-PROJECT] extra=undefined

# Created Collection

## 1

INFO 2022-07-07T17:43:31+10:00 Created collection | component=[IMX-CREATE-COLLECTION] extra=undefined
{
  "address": "0xefb053d2967ccd58227f982d80d2c795d28b7759",
  "name": "ENTER_COLLECTION_NAME",
  "description": "",
  "icon_url": "",
  "collection_image_url": "",
  "project_id": 65460,
  "project_owner_address": "",
  "metadata_api_url": ""
}

## 2

INFO 2022-07-08T19:25:50+10:00 Updated collection | component=[IMX-UPDATE-COLLECTION] extra=undefined
{
  "address": "0x4b089f5006778fef128506427235937d60b74dfb",
  "name": "Praveen Work Adventures Collection",
  "description": "",
  "icon_url": "",
  "collection_image_url": "",
  "project_id": 65790,
  "project_owner_address": "0x4db6bf7abafaae73d9010610166660032d6c796a",
  "metadata_api_url": ""
}

# Updated Collection

{
  "address": "0xefb053d2967ccd58227f982d80d2c795d28b7759",
  "name": "Praveen Work Adventures Collection",
  "description": "Work Adventures Collection from Praveen pertaining to IMX work",
  "icon_url": "https://gateway.pinata.cloud/ipfs/QmezNgGfKKTFA9Dc61EWns8kqfqP8XoGcsd6YsdecXXY8P",
  "collection_image_url": "https://gateway.pinata.cloud/ipfs/QmezNgGfKKTFA9Dc61EWns8kqfqP8XoGcsd6YsdecXXY8P",
  "project_id": 65460,
  "project_owner_address": "0x99c41e6ac42c909a228d558d0081c43234d6ad21",
  "metadata_api_url": "https://gateway.pinata.cloud/ipfs/QmevHJTBWG3eWbcKoaNnqQNQQRW38VVBebbK8b17rktYow"
}

# Added Metadata Schema

INFO 2022-07-07T17:46:51+10:00 Adding metadata schema to collection | component=[IMX-ADD-COLLECTION-METADATA-SCHEMA]
extra="0xefb053d2967CCd58227F982d80d2c795D28b7759"
REQUEST:
 POST https://api.ropsten.x.immutable.comv1collections/0xefb053d2967CCd58227F982d80d2c795D28b7759/metadata-schema
RESPONSE:
 201 Created
INFO 2022-07-07T17:46:51+10:00 Added metadata schema to collection | component=[IMX-ADD-COLLECTION-METADATA-SCHEMA]
extra="0xefb053d2967CCd58227F982d80d2c795D28b7759"
{
  "result": "OK"
}

# Minting NFTs

> @imtbl/imx-sdk-example@1.0.0 bulk-mint
> ts-node -r tsconfig-paths/register -r dotenv/config ./src/bulk-mint.ts "-n" "3" "-w" "0x99c41E6AC42c909A228D558d0081C43234d6Ad21"

tokenId
1
INFO 2022-07-07T23:23:24+10:00 MINTER REGISTRATION | component=imx-bulk-mint-script extra=undefined
INFO 2022-07-07T23:23:26+10:00 Minter registered, continuing... | component=imx-bulk-mint-script extra=undefined
INFO 2022-07-07T23:23:26+10:00 OFF-CHAIN MINT 3 NFTS | component=imx-bulk-mint-script extra=undefined
{
  results: [
    {
      token_id: '1',
      contract_address: '0xefb053d2967ccd58227f982d80d2c795d28b7759',
      tx_id: 4809263
    },
    {
      token_id: '2',
      contract_address: '0xefb053d2967ccd58227f982d80d2c795d28b7759',
      tx_id: 4809264
    },
    {
      token_id: '3',
      contract_address: '0xefb053d2967ccd58227f982d80d2c795d28b7759',
      tx_id: 4809265
    }
  ]
}


{"code":"bad_request","message":"user 0x99c41e6ac42c909a228d558d0081c43234d6ad21 does not own asset 0x04004f253cc6680bbea22e23007fa2244f3672c0efb8d336e783fe14691f6258"}

{"contract_address":"0x4b089f5006778fef128506427235937d60b74dfb","royalties":[{"recipient":"0x4db6bf7abafaae73d9010610166660032d6c796a","percentage":1}],"users":[{"ether_key":"0x4db6bf7abafaae73d9010610166660032d6c796a","tokens":[{"id":"4","blueprint":"123","royalties":[{"recipient":"0x4db6bf7abafaae73d9010610166660032d6c796a","percentage":1}]}]}],"auth_signature":""}

{"contract_address":"0x4b089f5006778fef128506427235937d60b74dfb","royalties":[{"recipient":"0x4db6bf7abafaae73d9010610166660032d6c796a","percentage":1}],"users":[{"ether_key":"0x4db6bf7abafaae73d9010610166660032d6c796a","tokens":[{"blueprint":"123","id":"4","royalties":[{"percentage":1,"recipient":"0x4db6bf7abafaae73d9010610166660032d6c796a"}]}]}],"auth_signature":""}

{"contract_address":"0x4b089f5006778fef128506427235937d60b74dfb","royalties":[{"recipient":"0x4db6bf7abafaae73d9010610166660032d6c796a","percentage":1}],"users":[{"ether_key":"0x4db6bf7abafaae73d9010610166660032d6c796a","tokens":[{"id":"4","blueprint":"1233","royalties":[{"recipient":"0x4db6bf7abafaae73d9010610166660032d6c796a","percentage":1}]}]}],"auth_signature":""}

Got a connection, launched process /private/var/folders/wn/9t_8m05102763q6pv080yk640000gn/T/GoLand/___go_build_main_go__1_ (pid = 86393).
2022/07/13 12:23:11 Mint Request being hashed: K�8��)�ѕJ�WDf|��N;#���"��t
2022/07/13 12:23:11 Auth Signature: 0x3b84ade8d20b9895039832791975a293e98f2b7be855d906b807967c50808ab6782ce43096cba10ec8379271a44147c325bb957f5ba67aff004d22ec7446504500
2022/07/13 13:41:36 Mint Tokens response:
&{[0xc00000e768]}
Exiting.

Got a connection, launched process /private/var/folders/wn/9t_8m05102763q6pv080yk640000gn/T/GoLand/___go_build_main_go__1_ (pid = 90659).
2022/07/14 01:05:49 Auth Signature: 0x3b84ade8d20b9895039832791975a293e98f2b7be855d906b807967c50808ab6782ce43096cba10ec8379271a44147c325bb957f5ba67aff004d22ec7446504500


Got a connection, launched process /private/var/folders/wn/9t_8m05102763q6pv080yk640000gn/T/GoLand/___go_build_main_go__1_ (pid = 88395).
2022/07/13 18:08:45 Auth Signature: 0x3f36839dfa73bae4f157fa23192eb128852ce687db8ba9f9b8b388a44a4f95e60d7a698724daa5c1881f62cb7fd6df0978ea08e4855fe7ebc875322b35958de401

Got a connection, launched process /private/var/folders/wn/9t_8m05102763q6pv080yk640000gn/T/GoLand/___go_build_main_go__1_ (pid = 90515).
2022/07/14 01:03:12 Auth Signature: 0x3f36839dfa73bae4f157fa23192eb128852ce687db8ba9f9b8b388a44a4f95e60d7a698724daa5c1881f62cb7fd6df0978ea08e4855fe7ebc875322b35958de401


