# OneofArrayXMLModel
(*OneofArrayXMLModel*)

### Available Operations

* [GetGenerate11](#getgenerate11) - Generate1
* [GetGenerate2](#getgenerate2) - Generate
* [PostValidate11](#postvalidate11) - Validate1
* [PostValidate2](#postvalidate2) - Validate

## GetGenerate11

This endpoint returns a 'CatsOrDogs' model as xml.

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
    res, err := s.OneofArrayXMLModel.GetGenerate11(ctx)
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

**[*operations.GetGenerate11Response](../../pkg/models/operations/getgenerate11response.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetGenerate2

This endpoint returns a 'CatsOrDogs' model as xml.

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
    res, err := s.OneofArrayXMLModel.GetGenerate2(ctx)
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

**[*operations.GetGenerate2Response](../../pkg/models/operations/getgenerate2response.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## PostValidate11

This endpoint expects a 'CatsOrDogs' model as xml.

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
    res, err := s.OneofArrayXMLModel.PostValidate11(ctx, []byte("0xd30B1A5d8E"))
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

**[*operations.PostValidate11Response](../../pkg/models/operations/postvalidate11response.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## PostValidate2

This endpoint expects a 'CatsOrDogs' model as xml.

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
    res, err := s.OneofArrayXMLModel.PostValidate2(ctx, []byte("0xC036A4f80F"))
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

**[*operations.PostValidate2Response](../../pkg/models/operations/postvalidate2response.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |
