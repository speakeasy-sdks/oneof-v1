# SimpleXMLModel

### Available Operations

* [Generate](#generate) - Generate
* [Generate1](#generate1) - Generate1
* [Validate](#validate) - Validate
* [Validate1](#validate1) - Validate1

## Generate

 This endpoint returns a cat model in xml

### Example Usage

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
    res, err := s.SimpleXMLModel.Generate(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.Body != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GenerateResponse](../../models/operations/generateresponse.md), error**


## Generate1

This endpoint returns a 'Dog' model as xml.

### Example Usage

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
    res, err := s.SimpleXMLModel.Generate1(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.Body != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.Generate1Response](../../models/operations/generate1response.md), error**


## Validate

This endpoint expects a 'Cat' model as xml.

### Example Usage

```go
package main

import(
	"context"
	"log"
	oneofv1 "github.com/speakeasy-sdks/oneof-v1"
	"github.com/speakeasy-sdks/oneof-v1/pkg/models/shared"
)

func main() {
    s := oneofv1.New()

    ctx := context.Background()
    res, err := s.SimpleXMLModel.Validate(ctx, []byte("error"))
    if err != nil {
        log.Fatal(err)
    }

    if res.Validate200TextPlainObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `request`                                             | [[]byte](../../models//.md)                           | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.ValidateResponse](../../models/operations/validateresponse.md), error**


## Validate1

This endpoint expects a 'Dog' model as xml.

### Example Usage

```go
package main

import(
	"context"
	"log"
	oneofv1 "github.com/speakeasy-sdks/oneof-v1"
	"github.com/speakeasy-sdks/oneof-v1/pkg/models/shared"
)

func main() {
    s := oneofv1.New()

    ctx := context.Background()
    res, err := s.SimpleXMLModel.Validate1(ctx, []byte("deserunt"))
    if err != nil {
        log.Fatal(err)
    }

    if res.Validate1200TextPlainObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `request`                                             | [[]byte](../../models//.md)                           | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.Validate1Response](../../models/operations/validate1response.md), error**

