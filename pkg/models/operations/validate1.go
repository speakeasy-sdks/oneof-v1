// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type Validate1Response struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse                 *http.Response
	Validate1200TextPlainObject *string
}

func (o *Validate1Response) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *Validate1Response) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *Validate1Response) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *Validate1Response) GetValidate1200TextPlainObject() *string {
	if o == nil {
		return nil
	}
	return o.Validate1200TextPlainObject
}
