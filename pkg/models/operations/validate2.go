// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type Validate2Response struct {
	ContentType                 string
	StatusCode                  int
	RawResponse                 *http.Response
	Validate2200TextPlainObject *string
}

func (o *Validate2Response) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *Validate2Response) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *Validate2Response) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *Validate2Response) GetValidate2200TextPlainObject() *string {
	if o == nil {
		return nil
	}
	return o.Validate2200TextPlainObject
}