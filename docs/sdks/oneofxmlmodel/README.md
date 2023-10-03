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

**[*operations.GetGenerateResponse](../../models/operations/getgenerateresponse.md), error**


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

**[*operations.GetGenerate1Response](../../models/operations/getgenerate1response.md), error**


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
    res, err := s.OneofXMLModel.PostValidate(ctx, []byte("Vb96?bTg=M"))
    if err != nil {
        log.Fatal(err)
    }

    if res.PostValidate200TextPlainObject != nil {
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

**[*operations.PostValidateResponse](../../models/operations/postvalidateresponse.md), error**


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
    res, err := s.OneofXMLModel.PostValidate1(ctx, []byte("cW>2F^@u/|"))
    if err != nil {
        log.Fatal(err)
    }

    if res.PostValidate1200TextPlainObject != nil {
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

**[*operations.PostValidate1Response](../../models/operations/postvalidate1response.md), error**

