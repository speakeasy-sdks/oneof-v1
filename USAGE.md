<!-- Start SDK Example Usage -->


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