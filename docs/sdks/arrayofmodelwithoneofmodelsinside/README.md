# ArrayOfModelWithOneofModelsInside
(*ArrayOfModelWithOneofModelsInside*)

### Available Operations

* [GetGenerate13](#getgenerate13) - Generate1
* [GetGenerate4](#getgenerate4) - Generate
* [PostValidate13](#postvalidate13) - Validate1
* [PostValidate4](#postvalidate4) - Validate

## GetGenerate13

This endpoint returns a 'ArrayOfCatOrDogObjects' model as xml.

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
    res, err := s.ArrayOfModelWithOneofModelsInside.GetGenerate13(ctx)
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

**[*operations.GetGenerate13Response](../../pkg/models/operations/getgenerate13response.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetGenerate4

This endpoint returns a 'ArrayOfCatOrDogObjects' model as xml.

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
    res, err := s.ArrayOfModelWithOneofModelsInside.GetGenerate4(ctx)
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

**[*operations.GetGenerate4Response](../../pkg/models/operations/getgenerate4response.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PostValidate13

This endpoint expects a 'ArrayOfCatOrDogObjects' model as xml.

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
    res, err := s.ArrayOfModelWithOneofModelsInside.PostValidate13(ctx, []byte("0xDd9028aE08"))
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

**[*operations.PostValidate13Response](../../pkg/models/operations/postvalidate13response.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## PostValidate4

This endpoint expects a 'ArrayOfCatOrDogObjects' model as xml.

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
    res, err := s.ArrayOfModelWithOneofModelsInside.PostValidate4(ctx, []byte("0x74d81dbA30"))
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

**[*operations.PostValidate4Response](../../pkg/models/operations/postvalidate4response.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
