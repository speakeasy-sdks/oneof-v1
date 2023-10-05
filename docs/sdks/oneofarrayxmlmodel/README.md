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
	"context"
	"log"
	oneofv1 "github.com/speakeasy-sdks/oneof-v1"
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

**[*operations.GetGenerate11Response](../../models/operations/getgenerate11response.md), error**


## GetGenerate2

This endpoint returns a 'CatsOrDogs' model as xml.

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

**[*operations.GetGenerate2Response](../../models/operations/getgenerate2response.md), error**


## PostValidate11

This endpoint expects a 'CatsOrDogs' model as xml.

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
    res, err := s.OneofArrayXMLModel.PostValidate11(ctx, []byte("[.$i%e9YEw"))
    if err != nil {
        log.Fatal(err)
    }

    if res.PostValidate11200TextPlainObject != nil {
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

**[*operations.PostValidate11Response](../../models/operations/postvalidate11response.md), error**


## PostValidate2

This endpoint expects a 'CatsOrDogs' model as xml.

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
    res, err := s.OneofArrayXMLModel.PostValidate2(ctx, []byte("o#0;e5bE!{"))
    if err != nil {
        log.Fatal(err)
    }

    if res.PostValidate2200TextPlainObject != nil {
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

**[*operations.PostValidate2Response](../../models/operations/postvalidate2response.md), error**

