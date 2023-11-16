# SimpleXMLModel
(*SimpleXMLModel*)

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
	oneofv1 "github.com/speakeasy-sdks/oneof-v1"
	"context"
	"log"
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

**[*operations.GenerateResponse](../../pkg/models/operations/generateresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## Generate1

This endpoint returns a 'Dog' model as xml.

### Example Usage

```go
package main

import(
	oneofv1 "github.com/speakeasy-sdks/oneof-v1"
	"context"
	"log"
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

**[*operations.Generate1Response](../../pkg/models/operations/generate1response.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## Validate

This endpoint expects a 'Cat' model as xml.

### Example Usage

```go
package main

import(
	oneofv1 "github.com/speakeasy-sdks/oneof-v1"
	"context"
	"log"
)

func main() {
    s := oneofv1.New()

    ctx := context.Background()
    res, err := s.SimpleXMLModel.Validate(ctx, []byte("0xd6BB5B71e0"))
    if err != nil {
        log.Fatal(err)
    }

    if res.Res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `request`                                             | [[]byte](../../.md)                                   | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.ValidateResponse](../../pkg/models/operations/validateresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## Validate1

This endpoint expects a 'Dog' model as xml.

### Example Usage

```go
package main

import(
	oneofv1 "github.com/speakeasy-sdks/oneof-v1"
	"context"
	"log"
)

func main() {
    s := oneofv1.New()

    ctx := context.Background()
    res, err := s.SimpleXMLModel.Validate1(ctx, []byte("0xb1A694c3A3"))
    if err != nil {
        log.Fatal(err)
    }

    if res.Res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `request`                                             | [[]byte](../../.md)                                   | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.Validate1Response](../../pkg/models/operations/validate1response.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |
