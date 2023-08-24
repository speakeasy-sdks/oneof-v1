// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type Generate1Response struct {
	Body        []byte
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}

func (o *Generate1Response) GetBody() []byte {
	if o == nil {
		return nil
	}
	return o.Body
}

func (o *Generate1Response) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *Generate1Response) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *Generate1Response) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
