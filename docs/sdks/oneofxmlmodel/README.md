# OneofXMLModel
(*OneofXMLModel*)

### Available Operations

* [GetGenerate](#getgenerate) - Generate
* [GetGenerate1](#getgenerate1) - Generate1
* [PostValidate](#postvalidate) - Validate
* [PostValidate1](#postvalidate1) - Validate1

## GetGenerate

This endpoint returns a 'CatOrDog' model as xml.

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
    res, err := s.OneofXMLModel.GetGenerate(ctx)
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

**[*operations.GetGenerateResponse](../../pkg/models/operations/getgenerateresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetGenerate1

This endpoint returns a 'CatOrDog' model as xml.

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
    res, err := s.OneofXMLModel.GetGenerate1(ctx)
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

**[*operations.GetGenerate1Response](../../pkg/models/operations/getgenerate1response.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## PostValidate

This endpoint expects a 'CatOrDog' model as xml.

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
    res, err := s.OneofXMLModel.PostValidate(ctx, []byte("0xcf557fcA6a"))
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

**[*operations.PostValidateResponse](../../pkg/models/operations/postvalidateresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## PostValidate1

This endpoint expects a 'CatOrDog' model as xml.

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
    res, err := s.OneofXMLModel.PostValidate1(ctx, []byte("0xfc648e7E3F"))
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

**[*operations.PostValidate1Response](../../pkg/models/operations/postvalidate1response.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |
