<!-- Start SDK Example Usage -->


```go
package main

import (
	"context"
	oneofv1 "github.com/speakeasy-sdks/oneof-v1"
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
<!-- End SDK Example Usage -->