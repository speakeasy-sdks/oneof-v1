# OneofArrayOrSingleXMLModelWithOptionalWrappingElement
(*OneofArrayOrSingleXMLModelWithOptionalWrappingElement*)

### Available Operations

* [Generate2](#generate2) - Generate2
* [GetGenerate12](#getgenerate12) - Generate1
* [GetGenerate3](#getgenerate3) - Generate
* [PostValidate12](#postvalidate12) - Validate1
* [PostValidate3](#postvalidate3) - Validate
* [Validate2](#validate2) - Validate2

## Generate2

This endpoint returns a 'CatsOrADogOrWolves' model as xml.

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
    res, err := s.OneofArrayOrSingleXMLModelWithOptionalWrappingElement.Generate2(ctx)
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

**[*operations.Generate2Response](../../pkg/models/operations/generate2response.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetGenerate12

This endpoint returns a 'CatsOrADogOrWolves' model as xml.

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
    res, err := s.OneofArrayOrSingleXMLModelWithOptionalWrappingElement.GetGenerate12(ctx)
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

**[*operations.GetGenerate12Response](../../pkg/models/operations/getgenerate12response.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetGenerate3

This endpoint returns a 'CatsOrADogOrWolves' model as xml.

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
    res, err := s.OneofArrayOrSingleXMLModelWithOptionalWrappingElement.GetGenerate3(ctx)
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

**[*operations.GetGenerate3Response](../../pkg/models/operations/getgenerate3response.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## PostValidate12

This endpoint expects a 'CatsOrADogOrWolves' model as xml.

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
    res, err := s.OneofArrayOrSingleXMLModelWithOptionalWrappingElement.PostValidate12(ctx, []byte("0xf9361EfF98"))
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

**[*operations.PostValidate12Response](../../pkg/models/operations/postvalidate12response.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## PostValidate3

This endpoint expects a 'CatsOrADogOrWolves' model as xml.

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
    res, err := s.OneofArrayOrSingleXMLModelWithOptionalWrappingElement.PostValidate3(ctx, []byte("0xcE70FDb90d"))
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

**[*operations.PostValidate3Response](../../pkg/models/operations/postvalidate3response.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## Validate2

This endpoint expects a 'CatsOrADogOrWolves' model as xml.

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
    res, err := s.OneofArrayOrSingleXMLModelWithOptionalWrappingElement.Validate2(ctx, []byte("0x9f52AfDFcE"))
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

**[*operations.Validate2Response](../../pkg/models/operations/validate2response.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |
