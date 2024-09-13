# Go API client for celenium

Celenium API is a powerful tool to access all blockchain data that is processed and indexed by our proprietary indexer. With Celenium API you can retrieve all historical data, off-chain data, blobs and statistics through our REST API. Celenium API indexer are open source, which allows you to not depend on third-party services. You can clone, build and run them independently, giving you full control over all components. If you have any questions or feature requests, please feel free to contact us. We appreciate your feedback!

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0
- Package version: 1.0.0
- Generator version: 7.9.0-SNAPSHOT
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://discord.gg/3k83Przqk8](https://discord.gg/3k83Przqk8)

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import celenium "github.com/mismirnov/celenium-api-go"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `celenium.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), celenium.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `celenium.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), celenium.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `celenium.ContextOperationServerIndices` and `celenium.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), celenium.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), celenium.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api-mainnet.celenium.io/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AddressAPI* | [**AddressBlobs**](docs/AddressAPI.md#addressblobs) | **Get** /address/{hash}/blobs | Get blobs pushed by address
*AddressAPI* | [**AddressDelegations**](docs/AddressAPI.md#addressdelegations) | **Get** /address/{hash}/delegations | Get delegations made by address
*AddressAPI* | [**AddressGrantee**](docs/AddressAPI.md#addressgrantee) | **Get** /address/{hash}/granters | Get grants where address is grantee
*AddressAPI* | [**AddressGrants**](docs/AddressAPI.md#addressgrants) | **Get** /address/{hash}/grants | Get grants made by address
*AddressAPI* | [**AddressMessages**](docs/AddressAPI.md#addressmessages) | **Get** /address/{hash}/messages | Get address messages
*AddressAPI* | [**AddressRedelegations**](docs/AddressAPI.md#addressredelegations) | **Get** /address/{hash}/redelegations | Get redelegations made by address
*AddressAPI* | [**AddressStats**](docs/AddressAPI.md#addressstats) | **Get** /address/{hash}/stats/{name}/{timeframe} | Get address stats
*AddressAPI* | [**AddressTransactions**](docs/AddressAPI.md#addresstransactions) | **Get** /address/{hash}/txs | Get address transactions
*AddressAPI* | [**AddressUndelegations**](docs/AddressAPI.md#addressundelegations) | **Get** /address/{hash}/undelegations | Get undelegations made by address
*AddressAPI* | [**AddressVesting**](docs/AddressAPI.md#addressvesting) | **Get** /address/{hash}/vestings | Get vesting for address
*AddressAPI* | [**GetAddress**](docs/AddressAPI.md#getaddress) | **Get** /address/{hash} | Get address info
*AddressAPI* | [**GetAddressCount**](docs/AddressAPI.md#getaddresscount) | **Get** /address/count | Get count of addresses in network
*AddressAPI* | [**ListAddress**](docs/AddressAPI.md#listaddress) | **Get** /address | List address info
*BlockAPI* | [**BlockBlobsCount**](docs/BlockAPI.md#blockblobscount) | **Get** /block/{height}/blobs/count | Count of blobs which was pushed by transaction
*BlockAPI* | [**GetBlock**](docs/BlockAPI.md#getblock) | **Get** /block/{height} | Get block info
*BlockAPI* | [**GetBlockBlobs**](docs/BlockAPI.md#getblockblobs) | **Get** /block/{height}/blobs | List blobs which was pushed in the block
*BlockAPI* | [**GetBlockCount**](docs/BlockAPI.md#getblockcount) | **Get** /block/count | Get count of blocks in network
*BlockAPI* | [**GetBlockEvents**](docs/BlockAPI.md#getblockevents) | **Get** /block/{height}/events | Get events from begin and end of block
*BlockAPI* | [**GetBlockMessages**](docs/BlockAPI.md#getblockmessages) | **Get** /block/{height}/messages | Get messages contained in the block
*BlockAPI* | [**GetBlockStats**](docs/BlockAPI.md#getblockstats) | **Get** /block/{height}/stats | Get block stats by height
*BlockAPI* | [**ListBlock**](docs/BlockAPI.md#listblock) | **Get** /block | List blocks info
*GasAPI* | [**GasEstimateForPfb**](docs/GasAPI.md#gasestimateforpfb) | **Get** /gas/estimate_for_pfb | Get estimated gas for pay for blob
*GasAPI* | [**GasPrice**](docs/GasAPI.md#gasprice) | **Get** /gas/price | Get estimated gas price
*GeneralAPI* | [**GetConstants**](docs/GeneralAPI.md#getconstants) | **Get** /constants | Get network constants
*GeneralAPI* | [**GetEnums**](docs/GeneralAPI.md#getenums) | **Get** /enums | Get celenium enumerators
*GeneralAPI* | [**Head**](docs/GeneralAPI.md#head) | **Get** /head | Get current indexer head
*NamespaceAPI* | [**GetBlob**](docs/NamespaceAPI.md#getblob) | **Post** /blob | Get namespace blob by commitment on height
*NamespaceAPI* | [**GetBlobLogs**](docs/NamespaceAPI.md#getbloblogs) | **Get** /namespace/{id}/{version}/blobs | Get blob changes for namespace
*NamespaceAPI* | [**GetBlobMetadata**](docs/NamespaceAPI.md#getblobmetadata) | **Post** /blob/metadata | Get blob metadata by commitment on height
*NamespaceAPI* | [**GetBlobs**](docs/NamespaceAPI.md#getblobs) | **Get** /blob | List all blobs with filters
*NamespaceAPI* | [**GetNamespace**](docs/NamespaceAPI.md#getnamespace) | **Get** /namespace/{id} | Get namespace info
*NamespaceAPI* | [**GetNamespaceActive**](docs/NamespaceAPI.md#getnamespaceactive) | **Get** /namespace/active | Get last used namespace
*NamespaceAPI* | [**GetNamespaceBase64**](docs/NamespaceAPI.md#getnamespacebase64) | **Get** /namespace_by_hash/{hash} | Get namespace info by base64
*NamespaceAPI* | [**GetNamespaceBlobs**](docs/NamespaceAPI.md#getnamespaceblobs) | **Get** /namespace_by_hash/{hash}/{height} | Get namespace blobs on height
*NamespaceAPI* | [**GetNamespaceByVersionAndId**](docs/NamespaceAPI.md#getnamespacebyversionandid) | **Get** /namespace/{id}/{version} | Get namespace info by id and version
*NamespaceAPI* | [**GetNamespaceCount**](docs/NamespaceAPI.md#getnamespacecount) | **Get** /namespace/count | Get count of namespaces in network
*NamespaceAPI* | [**GetNamespaceMessages**](docs/NamespaceAPI.md#getnamespacemessages) | **Get** /namespace/{id}/{version}/messages | Get namespace messages by id and version
*NamespaceAPI* | [**GetNamespaceRollups**](docs/NamespaceAPI.md#getnamespacerollups) | **Get** /namespace/{id}/{version}/rollups | List rollups using the namespace
*NamespaceAPI* | [**ListNamespace**](docs/NamespaceAPI.md#listnamespace) | **Get** /namespace | List namespace info
*RollupAPI* | [**GetRollup**](docs/RollupAPI.md#getrollup) | **Get** /rollup/{id} | Get rollup info
*RollupAPI* | [**GetRollupAllSeries**](docs/RollupAPI.md#getrollupallseries) | **Get** /rollup/stats/series | Get series for all rollups
*RollupAPI* | [**GetRollupBlobs**](docs/RollupAPI.md#getrollupblobs) | **Get** /rollup/{id}/blobs | Get rollup blobs
*RollupAPI* | [**GetRollupBySlug**](docs/RollupAPI.md#getrollupbyslug) | **Get** /rollup/slug/{slug} | Get rollup by slug
*RollupAPI* | [**GetRollupDistribution**](docs/RollupAPI.md#getrollupdistribution) | **Get** /rollup/{id}/distribution/{name}/{timeframe} | Get rollup distribution
*RollupAPI* | [**GetRollupNamespaces**](docs/RollupAPI.md#getrollupnamespaces) | **Get** /rollup/{id}/namespaces | Get rollup namespaces info
*RollupAPI* | [**GetRollupStats**](docs/RollupAPI.md#getrollupstats) | **Get** /rollup/{id}/stats/{name}/{timeframe} | Get rollup stats
*RollupAPI* | [**GetRollupsCount**](docs/RollupAPI.md#getrollupscount) | **Get** /rollup/count | Get count of rollups in network
*RollupAPI* | [**ListRollup**](docs/RollupAPI.md#listrollup) | **Get** /rollup | List rollups info
*RollupAPI* | [**ListRollup24h**](docs/RollupAPI.md#listrollup24h) | **Get** /rollup/day | List rollups info with stats by previous 24 hours
*RollupAPI* | [**RollupExport**](docs/RollupAPI.md#rollupexport) | **Get** /rollup/{id}/export | Export rollup blobs
*SearchAPI* | [**Search**](docs/SearchAPI.md#search) | **Get** /search | Search by hash
*StatsAPI* | [**Stats24hChanges**](docs/StatsAPI.md#stats24hchanges) | **Get** /stats/changes_24h | Get changes for 24 hours
*StatsAPI* | [**StatsMessagesCount24h**](docs/StatsAPI.md#statsmessagescount24h) | **Get** /stats/messages_count_24h | Get messages distribution for the last 24 hours
*StatsAPI* | [**StatsNamespaceUsage**](docs/StatsAPI.md#statsnamespaceusage) | **Get** /stats/namespace/usage | Get namespaces with sorting by size.
*StatsAPI* | [**StatsNsSeries**](docs/StatsAPI.md#statsnsseries) | **Get** /stats/namespace/series/{id}/{name}/{timeframe} | Get histogram for namespace with precomputed stats
*StatsAPI* | [**StatsPriceCurrent**](docs/StatsAPI.md#statspricecurrent) | **Get** /stats/price/current | Get current TIA price
*StatsAPI* | [**StatsPriceSeries**](docs/StatsAPI.md#statspriceseries) | **Get** /stats/price/series/{timeframe} | Get histogram with TIA price
*StatsAPI* | [**StatsRollup24h**](docs/StatsAPI.md#statsrollup24h) | **Get** /stats/rollup_stats_24h | Get rollups stats for last 24 hours
*StatsAPI* | [**StatsSeries**](docs/StatsAPI.md#statsseries) | **Get** /stats/series/{name}/{timeframe} | Get histogram with precomputed stats
*StatsAPI* | [**StatsSeriesCumulative**](docs/StatsAPI.md#statsseriescumulative) | **Get** /stats/series/{name}/{timeframe}/cumulative | Get cumulative histogram with precomputed stats
*StatsAPI* | [**StatsSquareSize**](docs/StatsAPI.md#statssquaresize) | **Get** /stats/square_size | Get histogram for square size distribution
*StatsAPI* | [**StatsStakingSeries**](docs/StatsAPI.md#statsstakingseries) | **Get** /stats/staking/series/{id}/{name}/{timeframe} | Get histogram for staking with precomputed stats
*StatsAPI* | [**StatsSummary**](docs/StatsAPI.md#statssummary) | **Get** /stats/summary/{table}/{function} | Get value by table and function
*TransactionsAPI* | [**GetTransaction**](docs/TransactionsAPI.md#gettransaction) | **Get** /tx/{hash} | Get transaction by hash
*TransactionsAPI* | [**GetTransactionEvents**](docs/TransactionsAPI.md#gettransactionevents) | **Get** /tx/{hash}/events | Get transaction events
*TransactionsAPI* | [**GetTransactionMessages**](docs/TransactionsAPI.md#gettransactionmessages) | **Get** /tx/{hash}/messages | Get transaction messages
*TransactionsAPI* | [**GetTransactionsCount**](docs/TransactionsAPI.md#gettransactionscount) | **Get** /tx/count | Get count of transactions in network
*TransactionsAPI* | [**ListGenesisTransactions**](docs/TransactionsAPI.md#listgenesistransactions) | **Get** /tx/genesis | List genesis transactions info
*TransactionsAPI* | [**ListTransactionBlobs**](docs/TransactionsAPI.md#listtransactionblobs) | **Get** /tx/{hash}/blobs | List blobs which was pushed by transaction
*TransactionsAPI* | [**ListTransactions**](docs/TransactionsAPI.md#listtransactions) | **Get** /tx | List transactions info
*TransactionsAPI* | [**TransactionBlobsCount**](docs/TransactionsAPI.md#transactionblobscount) | **Get** /tx/{hash}/blobs/count | Count of blobs which was pushed by transaction
*ValidatorAPI* | [**GetValidator**](docs/ValidatorAPI.md#getvalidator) | **Get** /validators/{id} | Get validator info
*ValidatorAPI* | [**GetValidatorBlocks**](docs/ValidatorAPI.md#getvalidatorblocks) | **Get** /validators/{id}/blocks | Get blocks which was proposed by validator
*ValidatorAPI* | [**GetValidatorUptime**](docs/ValidatorAPI.md#getvalidatoruptime) | **Get** /validators/{id}/uptime | Get validator&#39;s uptime and history of signed block
*ValidatorAPI* | [**ListValidator**](docs/ValidatorAPI.md#listvalidator) | **Get** /validators | List validators
*ValidatorAPI* | [**ValidatorCount**](docs/ValidatorAPI.md#validatorcount) | **Get** /validators/count | Get validator&#39;s count by status
*ValidatorAPI* | [**ValidatorDelegators**](docs/ValidatorAPI.md#validatordelegators) | **Get** /validators/{id}/delegators | Get validator&#39;s delegators
*ValidatorAPI* | [**ValidatorJails**](docs/ValidatorAPI.md#validatorjails) | **Get** /validators/{id}/jails | Get validator&#39;s jails
*VestingAPI* | [**GetVestingPeriods**](docs/VestingAPI.md#getvestingperiods) | **Get** /vesting/{id}/periods | Periods vesting periods by id


## Documentation For Models

 - [GithubComCeleniumIoCelestiaIndexerInternalStorageTypesStatus](docs/GithubComCeleniumIoCelestiaIndexerInternalStorageTypesStatus.md)
 - [HandlerError](docs/HandlerError.md)
 - [ResponsesAddress](docs/ResponsesAddress.md)
 - [ResponsesBalance](docs/ResponsesBalance.md)
 - [ResponsesBlob](docs/ResponsesBlob.md)
 - [ResponsesBlobLog](docs/ResponsesBlobLog.md)
 - [ResponsesBlock](docs/ResponsesBlock.md)
 - [ResponsesBlockStats](docs/ResponsesBlockStats.md)
 - [ResponsesChange24hBlockStats](docs/ResponsesChange24hBlockStats.md)
 - [ResponsesConstants](docs/ResponsesConstants.md)
 - [ResponsesCountItem](docs/ResponsesCountItem.md)
 - [ResponsesDelegation](docs/ResponsesDelegation.md)
 - [ResponsesDenomMetadata](docs/ResponsesDenomMetadata.md)
 - [ResponsesDistributionItem](docs/ResponsesDistributionItem.md)
 - [ResponsesEnums](docs/ResponsesEnums.md)
 - [ResponsesEvent](docs/ResponsesEvent.md)
 - [ResponsesGasPrice](docs/ResponsesGasPrice.md)
 - [ResponsesGrant](docs/ResponsesGrant.md)
 - [ResponsesHistogramItem](docs/ResponsesHistogramItem.md)
 - [ResponsesJail](docs/ResponsesJail.md)
 - [ResponsesLightBlobLog](docs/ResponsesLightBlobLog.md)
 - [ResponsesMessage](docs/ResponsesMessage.md)
 - [ResponsesMessageForAddress](docs/ResponsesMessageForAddress.md)
 - [ResponsesNamespace](docs/ResponsesNamespace.md)
 - [ResponsesNamespaceKind](docs/ResponsesNamespaceKind.md)
 - [ResponsesNamespaceMessage](docs/ResponsesNamespaceMessage.md)
 - [ResponsesNamespaceUsage](docs/ResponsesNamespaceUsage.md)
 - [ResponsesODS](docs/ResponsesODS.md)
 - [ResponsesODSItem](docs/ResponsesODSItem.md)
 - [ResponsesPrice](docs/ResponsesPrice.md)
 - [ResponsesRedelegation](docs/ResponsesRedelegation.md)
 - [ResponsesRollup](docs/ResponsesRollup.md)
 - [ResponsesRollupAllSeriesItem](docs/ResponsesRollupAllSeriesItem.md)
 - [ResponsesRollupStats24h](docs/ResponsesRollupStats24h.md)
 - [ResponsesRollupWithDayStats](docs/ResponsesRollupWithDayStats.md)
 - [ResponsesRollupWithStats](docs/ResponsesRollupWithStats.md)
 - [ResponsesSearchItem](docs/ResponsesSearchItem.md)
 - [ResponsesSeriesItem](docs/ResponsesSeriesItem.md)
 - [ResponsesShortRollup](docs/ResponsesShortRollup.md)
 - [ResponsesShortValidator](docs/ResponsesShortValidator.md)
 - [ResponsesSignedBlocks](docs/ResponsesSignedBlocks.md)
 - [ResponsesState](docs/ResponsesState.md)
 - [ResponsesTPS](docs/ResponsesTPS.md)
 - [ResponsesTimeValueItem](docs/ResponsesTimeValueItem.md)
 - [ResponsesTx](docs/ResponsesTx.md)
 - [ResponsesTxForAddress](docs/ResponsesTxForAddress.md)
 - [ResponsesUndelegation](docs/ResponsesUndelegation.md)
 - [ResponsesValidator](docs/ResponsesValidator.md)
 - [ResponsesValidatorCount](docs/ResponsesValidatorCount.md)
 - [ResponsesValidatorUptime](docs/ResponsesValidatorUptime.md)
 - [ResponsesVesting](docs/ResponsesVesting.md)
 - [ResponsesVestingPeriod](docs/ResponsesVestingPeriod.md)
 - [TypesEventType](docs/TypesEventType.md)
 - [TypesMsgAddressType](docs/TypesMsgAddressType.md)
 - [TypesMsgType](docs/TypesMsgType.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### ApiKeyAuth

- **Type**: API key
- **API key parameter name**: apikey
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: ApiKeyAuth and passed in as the auth context for each request.

Example

```go
auth := context.WithValue(
		context.Background(),
		celenium.ContextAPIKeys,
		map[string]celenium.APIKey{
			"ApiKeyAuth": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

celenium@pklabs.me

