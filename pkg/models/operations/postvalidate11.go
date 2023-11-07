// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type PostValidate11Response struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	Res         *string
}

func (o *PostValidate11Response) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PostValidate11Response) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PostValidate11Response) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PostValidate11Response) GetRes() *string {
	if o == nil {
		return nil
	}
	return o.Res
}
