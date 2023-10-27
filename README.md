# github.com/speakeasy-sdks/oneof-v1

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/oneof-v1
```
<!-- End SDK Installation -->

## SDK Example Usage
<!-- Start SDK Example Usage -->
```go
package main

import (
	"context"
	oneofv1 "github.com/speakeasy-sdks/oneof-v1"
	"log"
)

func main() {
	s := oneofv1.New()

	ctx := context.Background()
	res, err := s.ArrayOfModelWithOneofModelsInside.GetGenerate13(ctx)
	if err != nil {
		log.Fatal(err)
	}

	if res.Body != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## Available Resources and Operations


### [ArrayOfModelWithOneofModelsInside](docs/sdks/arrayofmodelwithoneofmodelsinside/README.md)

* [GetGenerate13](docs/sdks/arrayofmodelwithoneofmodelsinside/README.md#getgenerate13) - Generate1
* [GetGenerate4](docs/sdks/arrayofmodelwithoneofmodelsinside/README.md#getgenerate4) - Generate
* [PostValidate13](docs/sdks/arrayofmodelwithoneofmodelsinside/README.md#postvalidate13) - Validate1
* [PostValidate4](docs/sdks/arrayofmodelwithoneofmodelsinside/README.md#postvalidate4) - Validate

### [OneofArrayXMLModel](docs/sdks/oneofarrayxmlmodel/README.md)

* [GetGenerate11](docs/sdks/oneofarrayxmlmodel/README.md#getgenerate11) - Generate1
* [GetGenerate2](docs/sdks/oneofarrayxmlmodel/README.md#getgenerate2) - Generate
* [PostValidate11](docs/sdks/oneofarrayxmlmodel/README.md#postvalidate11) - Validate1
* [PostValidate2](docs/sdks/oneofarrayxmlmodel/README.md#postvalidate2) - Validate

### [OneofArrayOrSingleXMLModelWithOptionalWrappingElement](docs/sdks/oneofarrayorsinglexmlmodelwithoptionalwrappingelement/README.md)

* [Generate2](docs/sdks/oneofarrayorsinglexmlmodelwithoptionalwrappingelement/README.md#generate2) - Generate2
* [GetGenerate12](docs/sdks/oneofarrayorsinglexmlmodelwithoptionalwrappingelement/README.md#getgenerate12) - Generate1
* [GetGenerate3](docs/sdks/oneofarrayorsinglexmlmodelwithoptionalwrappingelement/README.md#getgenerate3) - Generate
* [PostValidate12](docs/sdks/oneofarrayorsinglexmlmodelwithoptionalwrappingelement/README.md#postvalidate12) - Validate1
* [PostValidate3](docs/sdks/oneofarrayorsinglexmlmodelwithoptionalwrappingelement/README.md#postvalidate3) - Validate
* [Validate2](docs/sdks/oneofarrayorsinglexmlmodelwithoptionalwrappingelement/README.md#validate2) - Validate2

### [OneofXMLModel](docs/sdks/oneofxmlmodel/README.md)

* [GetGenerate](docs/sdks/oneofxmlmodel/README.md#getgenerate) - Generate
* [GetGenerate1](docs/sdks/oneofxmlmodel/README.md#getgenerate1) - Generate1
* [PostValidate](docs/sdks/oneofxmlmodel/README.md#postvalidate) - Validate
* [PostValidate1](docs/sdks/oneofxmlmodel/README.md#postvalidate1) - Validate1

### [SimpleXMLModel](docs/sdks/simplexmlmodel/README.md)

* [Generate](docs/sdks/simplexmlmodel/README.md#generate) - Generate
* [Generate1](docs/sdks/simplexmlmodel/README.md#generate1) - Generate1
* [Validate](docs/sdks/simplexmlmodel/README.md#validate) - Validate
* [Validate1](docs/sdks/simplexmlmodel/README.md#validate1) - Validate1
<!-- End SDK Available Operations -->



<!-- Start Dev Containers -->

<!-- End Dev Containers -->



<!-- Start Pagination -->
# Pagination

Some of the endpoints in this SDK support pagination. To use pagination, you make your SDK calls as usual, but the
returned response object will have a `Next` method that can be called to pull down the next group of results. If the
return value of `Next` is `nil`, then there are no more pages to be fetched.

Here's an example of one such pagination call:
<!-- End Pagination -->



<!-- Start Go Types -->

<!-- End Go Types -->



<!-- Start Error Handling -->
# Error Handling

Handling errors in your SDK should largely match your expectations.  All operations return a response object or an error, they will never return both.  When specified by the OpenAPI spec document, the SDK will return the appropriate subclass.


<!-- End Error Handling -->



<!-- Start Server Selection -->
# Server Selection

## Select Server by Index

You can override the default server globally using the `WithServerIndex` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the indexes associated with the available servers:

| # | Server | Variables |
| - | ------ | --------- |
| 0 | `https://example.com/api` | None |
| 1 | `http://localhost:3000/body-xml` | None |

For example:


```go
package main

import (
	"context"
	oneofv1 "github.com/speakeasy-sdks/oneof-v1"
	"log"
)

func main() {
	s := oneofv1.New(
		oneofv1.WithServerIndex(1),
	)

	ctx := context.Background()
	res, err := s.ArrayOfModelWithOneofModelsInside.GetGenerate13(ctx)
	if err != nil {
		log.Fatal(err)
	}

	if res.Body != nil {
		// handle response
	}
}

```


## Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL` option when initializing the SDK client instance. For example:


```go
package main

import (
	"context"
	oneofv1 "github.com/speakeasy-sdks/oneof-v1"
	"log"
)

func main() {
	s := oneofv1.New(
		oneofv1.WithServerURL("https://example.com/api"),
	)

	ctx := context.Background()
	res, err := s.ArrayOfModelWithOneofModelsInside.GetGenerate13(ctx)
	if err != nil {
		log.Fatal(err)
	}

	if res.Body != nil {
		// handle response
	}
}

```
<!-- End Server Selection -->



<!-- Start Custom HTTP Client -->
# Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client -->

<!-- Placeholder for Future Speakeasy SDK Sections -->



### Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release!

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
