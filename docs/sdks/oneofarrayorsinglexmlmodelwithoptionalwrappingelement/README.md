# OneofArrayOrSingleXMLModelWithOptionalWrappingElement

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
	"context"
	"log"
	"github.com/speakeasy-sdks/oneof-v1"
)

func main() {
    s := oneof.New()

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

**[*operations.Generate2Response](../../models/operations/generate2response.md), error**


## GetGenerate12

This endpoint returns a 'CatsOrADogOrWolves' model as xml.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/oneof-v1"
)

func main() {
    s := oneof.New()

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

**[*operations.GetGenerate12Response](../../models/operations/getgenerate12response.md), error**


## GetGenerate3

This endpoint returns a 'CatsOrADogOrWolves' model as xml.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/oneof-v1"
)

func main() {
    s := oneof.New()

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

**[*operations.GetGenerate3Response](../../models/operations/getgenerate3response.md), error**


## PostValidate12

This endpoint expects a 'CatsOrADogOrWolves' model as xml.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/oneof-v1"
	"github.com/speakeasy-sdks/oneof-v1/pkg/models/shared"
)

func main() {
    s := oneof.New()

    ctx := context.Background()
    res, err := s.OneofArrayOrSingleXMLModelWithOptionalWrappingElement.PostValidate12(ctx, []byte("unde"))
    if err != nil {
        log.Fatal(err)
    }

    if res.PostValidate12200TextPlainObject != nil {
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

**[*operations.PostValidate12Response](../../models/operations/postvalidate12response.md), error**


## PostValidate3

This endpoint expects a 'CatsOrADogOrWolves' model as xml.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/oneof-v1"
	"github.com/speakeasy-sdks/oneof-v1/pkg/models/shared"
)

func main() {
    s := oneof.New()

    ctx := context.Background()
    res, err := s.OneofArrayOrSingleXMLModelWithOptionalWrappingElement.PostValidate3(ctx, []byte("nulla"))
    if err != nil {
        log.Fatal(err)
    }

    if res.PostValidate3200TextPlainObject != nil {
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

**[*operations.PostValidate3Response](../../models/operations/postvalidate3response.md), error**


## Validate2

This endpoint expects a 'CatsOrADogOrWolves' model as xml.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/oneof-v1"
	"github.com/speakeasy-sdks/oneof-v1/pkg/models/shared"
)

func main() {
    s := oneof.New()

    ctx := context.Background()
    res, err := s.OneofArrayOrSingleXMLModelWithOptionalWrappingElement.Validate2(ctx, []byte("corrupti"))
    if err != nil {
        log.Fatal(err)
    }

    if res.Validate2200TextPlainObject != nil {
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

**[*operations.Validate2Response](../../models/operations/validate2response.md), error**

