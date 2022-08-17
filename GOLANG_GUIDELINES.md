# Golang guidelines

**Note:** These guidelines are from [this doc](https://immutable.atlassian.net/wiki/spaces/ISP/pages/1866434963/Golang+Guidelines) and may have to be updated from time to time.

## Table of contents:
* [Filenames](#filenames)
* [Folder names](#folder-names)
* [Package names](#package-names)
* [Comments](#comments)
* [Function length](#function-length)
* [DRY or WET](#dry-or-wet)
* [Pass by value or pointer](#pass-by-value-or-pointer)
* [Golang context](#golang-context)
* [Errors](#errors)
* [Package level state](#package-level-state)
* [Ordering](#ordering)
* [Documentation](#documentation)
* [Getters](#getters)
* [Variable names](#variable-names)
* [Don't stutter](#dont-stutter)
* [Slice capacity](#slice-capacity)
* [Return nil slice intentionally](#return-nil-slice-intentionally)
* [Pointers to interfaces](#pointers-to-interfaces)
* [Composition through embedding](#composition-through-embedding)
* [Use interfaces to inject dependencies](#use-interfaces-to-inject-dependencies)
* [Don't panic in production runtimes](#dont-panic-in-production-runtimes)
* [Use `common/errors`](#use-commonerrors)
* [Do not use channels and goroutines for synchronization](#do-not-use-channels-and-goroutines-for-synchronization)
* [Test names](#test-names)

---

## Filenames

Use lowercase and underscore.

:white_check_mark: Good:
```
full_withdrawal.go
```

## Folder names

[a-z0-9-] 

## Package names

* Short and clear
* Singular
* Lowercase
* No under_scores or mixedCaps
* Package names may be abbreviated if the abbreviation will familiar to others. If abbreviating a package name makes it ambiguous or unclear, don’t do it
* Avoid meaningless package names such as `utils`, `common` or `misc`
* Avoid having a single package for all your APIs
* Avoid using popular standard package names such as `io` or `http`
* When renaming an imported package, the local name should follow the same guidelines as package names
* For more information, please see https://go.dev/blog/package-names

:white_check_mark: Good:
```
package reconcil
package launchdarkly // single word, no underscores or mixed Caps.
package trade
```

:x: Bad:
```
package onchaineventreconcil
package onchain_event_reconcil
package onchainEventReconcil
package trades
```

## Comments

* To document a type, variable, constant, function, or even a package, write a regular comment directly preceding its declaration, with no intervening blank line.
* All packages, public functions, variables must have comments.

:white_check_mark: Good:
```
// Package sort provides primitives for sorting slices and user-defined
// collections.
package sort
```

* Comments should explain "why", not "how".
* For more details, see https://go.dev/blog/godoc

:white_check_mark: Good:
```
// check if all users have correct names
for _, user := range users {
    ...
}
```

:x: Bad:
```
// loop through users
for _, user := range users {
	// check if name contains underscore and slashes
    ...
}
```

## Function length
Functions should have a max length of 25 lines of code.

## DRY or WET
DRY = "Don't repeat yourself"

WET = "Write everything twice"

Only the following things must be DRY:
* Code that does not contain business logic

:white_check_mark: Good:
```
package newrelic

// Wrap around lambda function handler to provide an extra argument of newrelic transaction 
func Start(handler ...) {
...
}
```

* Code containing atomic business logic

:x: Bad (expose checking assets contained in a vault to be fungible or non-fungible. Leave room for missing this check in new places that require crediting a vault):
```
if vault is fungible {
  CreditFungibleVault()
} else {
  CreditNonFungibleVault()
}
...

...
func CreditFungibleVault(...)
func CreditNonFungibleVault(...)
```

:white_check_mark: Good (treat checking fungibility of an asset in a vault as part of crediting a vault. Eliminate error margin for missing this check):
```
...
// atomic business logic
CreditVault(vaultId, 1000)
...

// CreditVault applies correct check based on the asset contained in the vault.
func CreditVault(vaultId, amount uint64) error {
  ...
  if asset is fungible 
    return creditFungibleVault(...)
  if nonFungibleVault balance is not 1
    return err 
  return creditNonFungibleVault(...)
}
```

## Pass by value or pointer

* By default, pass by value and avoid mutation.

:white_check_mark: Good:
```
var newSlice:...
for _, elem := range slice {
  newSlice = append(newSlice, transform(elem)) // reference by value
}
```
```
va := Atype { ... }
type Btype struct { Atype }

func transformV(v Atype) Btype {
  b := Btype {v}
  b.xxx = ...
  return b
}

v1 := transformV(v) // use return value and different type
```

:x: Bad (reference by pointer and mutation):
```
for i, _ := range slice {
  slice[i].someField = ... // reference by pointer and in place mutation
}

func ..(v *Atype) {
  v.xxx = ....
}
```

* Delay performance optimisation until it’s called out by linter (https://go-critic.com/overview.html#rangevalcopy) or proven in production that a particular code path is slow/resource hungry because of “pass by value”.

## Golang context

* Usage of Golang context to carry (`context.WithValue` and `context.Value`) implicit values is discouraged.
* Usage of `fromContext` for 3rd party module is allowed.

:white_check_mark: Good:
```
// pass v in 
func (ctx context.Context, v string) {
}
```
```
// 3rd party lib
func (ctx context.Context) {
  logEntry := logrus.FromContext(ctx)
  txn := newrelic.FromContext(ctx)
}
```

:x: Bad:
```
func (ctx context.Context) {
  ctx := context.WithValue(r.Context(), key, val)
  v := ctx.Value(....)
  
  // concert the type and use v here
  v.(string)
}
```

## Errors
* Never hide them

:white_check_mark: Good:
```
if err := someFunction(...); err != nil {
... // handle the error here.
}

returnValue, err := someFunction(...)
if err != nil {
... // handle the error here.
}
```

:x: Bad:
```
returnValue, _ := someFunction(...)
```

* 500 http status indicates reduced service availability and will trigger alerts. Do NOT produce user facing errors as internal errors

:white_check_mark: Good:
```
// This error is user facing. It should be expressed as a bad request (status 400) on the http layer.
type UserDoesNotOwnAsset struct {
}
...

if !doesUserOwnThisAsset(user, asset) {
  return imxerrors.WrapUser(UserDoesNotOwnAsset{})
}
```

:x: Bad (validation error is treated as critical error):
```
// This Error is user facing. It should be expressed as a bad request (status 400) on the http layer.
type UserDoesNotOwnAsset struct {
}
...

if !doesUserOwnThisAsset(user, asset) {
  return imxerrors.WrapCritical(UserDoesNotOwnAsset{})
}
```

## Package level state

* Packages should be stateless by default.
* This is unless there is no alternative (ie. lambda function handler top level var will be cached for much better performance)

:white_check_mark: Good (lambda handler):
```
package main

// store is cached in between lambda function invocation to make a noticable performance impact.
var store *pg.Store

func Handler(....) {
  ...
}

func main() {
	newrelic.Start(Handler)
}
```

* Avoid exposing package level var to prevent it from being mutated.

:white_check_mark: Good:
```
package ...

var someConfig // private in this package

func New() *Config {
	// Copy config
	conf := someConfig
	return &conf // expose a copy.
}
```

:x: Bad:
```
package abc

// This is exposed as abc.SomeConfig on this package.
var SomeConfig = ...
```

## Ordering

Source files should exhibit this order:
* Constants
* Variables
* Types
* Funcs (in call order)

:x: Bad (funcs not in call order):
```
package abc

type Car struct {}

func startEngine() {}

func (c Car) Start() {
  startEngine()
}
```

:white_check_mark: Good (funcs in correct order):
```
package ...

const (
  hello = "world"
)

type Car struct {}

func (c Car) Start() {
  startEngine()
}

func startEngine() {}
```

## Documentation

* Err on the side of over-documenting through comments:
    * Above func, const, var, and type declarations (Godoc format)
    * Comments within funcs that split up logical blocks of code
* Gofmt will enforce [Godoc](https://pkg.go.dev/golang.org/x/tools/cmd/godoc) commenting rules

:x: Bad (no comments):
```
func (c *Controller) carrySnapshotForward(ctx context.Context, srcBlockNum, destBlockNum, limit uint64) error {
	count, err := c.GetSnapshotBalanceCount(ctx, srcBlockNum)
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.NewCarryForwardNoRows(srcBlockNum, destBlockNum)
	}
	if err := c.SetCarryForwardCheckpoint(ctx, srcBlockNum, destBlockNum, count); err != nil {
		return err
	}
	checkpoint, err := c.GetCarryForwardCheckpoint(ctx, srcBlockNum, destBlockNum)
	if err != nil {
		return err
	}
	if checkpoint.IsComplete() {
		return nil
	}
	for offset := checkpoint.NextRow; offset < count; offset += limit {
		if err := c.CarrySnapshotForward(ctx, srcBlockNum, destBlockNum, limit, offset); err != nil {
			return err
		}
		if err := c.UpdateCarryForwardCheckpoint(ctx, srcBlockNum, destBlockNum, offset+limit); err != nil {
			return err
		}
	}
	return nil
}
```

:white_check_mark: Good (Godoc and inline comments):
```
// CarrySnapshotForward will upsert all existing snapshots to the new specified block
func (c *Controller) CarrySnapshotForward(ctx context.Context, srcBlockNum, destBlockNum, limit uint64) error {
	// Get snapshot count
	count, err := c.GetSnapshotBalanceCount(ctx, srcBlockNum)
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.NewCarryForwardNoRows(srcBlockNum, destBlockNum)
	}

	// Initial checkpoint
	if err := c.SetCarryForwardCheckpoint(ctx, srcBlockNum, destBlockNum, count); err != nil {
		return err
	}

	// Get checkpoint in case previously processed
	checkpoint, err := c.GetCarryForwardCheckpoint(ctx, srcBlockNum, destBlockNum)
	if err != nil {
		return err
	}

	// Already finished
	if checkpoint.IsComplete() {
		return nil
	}

	// Process from checkpoint
	for offset := checkpoint.NextRow; offset < count; offset += limit {
		// Carry over
		if err := c.CarrySnapshotForward(ctx, srcBlockNum, destBlockNum, limit, offset); err != nil {
			return err
		}

		// Update checkpoint
		if err := c.UpdateCarryForwardCheckpoint(ctx, srcBlockNum, destBlockNum, offset+limit); err != nil {
			return err
		}
	}

	return nil
}
```

## Getters
* Getters and setters are encouraged for sophisticated types that perform some logic through member types
    * Member variables are private
* Types used for serialisation must have public member variables:
    * This is because other pkgs must read / write their member variables directly to marshal/unmarshal etc
    * Treat them as dumb types that don’t need member funcs at all, including getters
* Getters do not have the prefix “Get” in the func names

:white_check_mark: Good (private member variables for sophisticated types):
```
type MySmartType struct{
  data string
}

func New(data string) *MySmartType {
  return &MySmartType{
    data: data
  }
}

func (mst MySmartType) SomeLogic() {}

func (mst MySmartType) Data() string {
  return mst.data
}
```

:x: Bad (public member variables for sophisticated types):
```
type MySmartType struct{
  Data string
}

func (mst MySmartType) SomeLogic() {}
```

## Variable names
* All variables are at least 3 characters in length (except for loop index)
* Shorthand names must be unequivocal in the context
* Acronyms must have single capitalization
    * Linter will complain about common acronyms, but not unique ones like IMX

:white_check_mark: Good (acronym single capitalization):
```
myIMX := ""
imxForYou := ""
IMXForYou := ""
```

:x: Bad (acronym multiple capitalization):
```
myImx := ""
ImxForYou := ""
```

* Constants:
    * `camelCase` for private constants
    * `PascalCase` for public constants
    * Common prefixes for the same context

:white_check_mark: Good (constants with a common prefix):
```
package store

const (
    TransactionStatusActive = TransactionStatus("active")
	TransactionStatusFilled = TransactionStatus("filled")
	TransactionStatusCancelled = TransactionStatus("cancelled")
)
```

:x: Bad (no prefix for constants with the same context):
```
package store

const (
    Active = TransactionStatus("active")
	Filled = TransactionStatus("filled")
	Cancelled = TransactionStatus("cancelled")
)
```

## Don’t stutter
When implementing package-level APIs, make sure your public functions and variables don’t stutter when prefixed with the package name.

:white_check_mark: Good (public function does not repeat package name (does not stutter)):
```
log.Info()
```

:x: Bad (public function repeats package name (stutters)):
```
log.LogInfo() // Stutter
```

## Slice capacity
Specify slice capacity where possible in order to allocate memory for the container up front. This minimizes subsequent allocations (by copying and resizing of the underlying array) as elements are added.

:white_check_mark: Good (slice capacity pre-allocated):
```
tokens := make([]tokens{}, len(orders))
for i := range orders {
  tokens = append(tokens, NewToken(orders[i].Token))
}
```

:x: Bad (slice capacity not pre-allocated):
```
tokens := []Token{}
for i := range orders {
  tokens = append(tokens, NewToken(orders[i].Token))
}
```

## Return nil slice intentionally
* When specifying a slice to return in function declaration, the slice will be nil by default
* It is fine to return nil slice, just make sure the calling code is aware of this

:white_check_mark: Good (nil slice never returned):
```
func GetTokens(count int) (tokens []Tokens) {
  if count == 0 {
    return []Tokens{}
  }
  // ...
}
```

:x: Bad (nil slice returned, unintentionally):
```
func GetTokens(count int) (tokens []Tokens) {
  if count == 0 {
    return
  }
  // ...
}
```

## Pointers to interfaces
* Interfaces behave like pointers (see [reference types](https://github.com/go101/go101/wiki/About-the-terminology-%22reference-type%22-in-Go))
    * Pointer to interface is akin to pointer to pointer
* Do not use pointers to interfaces
    * You will get this kind of error:
    ```
    *X does not implement *Y (type *Z is pointer to interface, not interface)
    ```
* Do not use pointers to pointers
    * You will just never have a use case for it. Here is just an example:

:x: Bad (assigning a slice to a pointer):
```
func AllocToken(ptr *[]Token, count int) {
  *ptr = make([]Token{}, count)
}
```

:white_check_mark: Good (return slice):
```
func AllocToken(count int) []Token {
  return make([]Token{}, count)
}
```

## Composition through embedding
* Golang does not support classes and inheritance. Instead, it supports extension through composition
    * Nonetheless, when structs are embedded, their member vars and funcs are accessible through the parent type, similar to inheritance
* Structs can be embedded into structs
```
// ReadWriter stores pointers to a Reader and a Writer.
// It implements io.ReadWriter.
type ReadWriter struct {
    *Reader  // *bufio.Reader
    *Writer  // *bufio.Writer
}
```
* Interfaces can be embedded into structs
```
// ReadWriter is the interface that combines the Reader and Writer interfaces.
type ReadWriter interface {
    Reader
    Writer
}
```

## Use interfaces to inject dependencies
* Injecting dependancies allows us to replace them with other implementations and mocks for testing
* Data / metadata is not a dependancy and should be passed in through integral types or value objects (e.g. do not pass config around)

:white_check_mark: Good (client is injected, config types are not introduced at all):
```
type Executor struct {
  cloudClient cloud.Client
}

func New(client cloud.Client, retries int) *Executor {
  return &Executor{
    cloudClient: client,
    maxRetries: retries,
  }
}
```

:x: Bad (config and client is nested into business logic type):
```
type Executor struct {
  awsClient aws.Client
  maxRetries int
}

func New() *Executor {
  return &Executor{
    awsClient: aws.NewClient(config.New().GetAWSConf()),
    maxRetries: config.New().Retries,
  }
}
```

## Don’t panic in production runtimes
* Panics can cause cascading failures and interrupt unrelated workloads in the case of concurrent programs
* Use panics with unit tests to prove that an error case is impossible and there is no need to return errors in the respective func
* You can recover() from panic() for test scenarios, but these should be assessed case by case - feel free to raise in the Golang Slack channel:  

## Use `common/errors`
* The errors pkg should provide for all the error use cases you have
    * If you have a situation where you need an error that exhibits behaviours not supported by the types found in common/errors, make your own error type that inherits the behaviours you need. Make sure to implement unit tests for them.
* Each of the error types map to an HTTP response
    * Use WrapXX() to override error behaviours
```
accounts, err := c.GetAccounts(ctx, starkKeys)
if err != nil {
	if imxerrors.IsNotFound(err) {
		return imxerrors.WrapCritical(err)
	}
	return err
}
```

## Do not use channels and goroutines for synchronization
* Use goroutines and channels for concurrent workloads
* Use  https://pkg.go.dev/sync for synchronizing data access
    * Much less code
    * More performant

## Test names
* Tests should start with Test followed by the function name, ie. `TestMyFunction`
* Tests should also include the state and expected outcome, delimited with underscores: TestMyFunction_WithThisState_HasThisResult

:white_check_mark: Good (clearly specifies state and expected outcome):
```
func TestQualification_UsersAlreadyQualified_DoesNotDuplicateQualification(t *testing.T) {
  ...
}
```

:x: Bad (does not clearly specify state and expected outcome):
```
func TestQualificationUsersQualified(t *testing.T) {
  ...
}
```