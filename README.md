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

import(
	"context"
	"log"
	oneofv1 "github.com/speakeasy-sdks/oneof-v1"
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

<!-- Placeholder for Future Speakeasy SDK Sections -->



### Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release!

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
