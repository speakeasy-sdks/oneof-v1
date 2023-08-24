// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type GetGenerate11Response struct {
	Body        []byte
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}

func (o *GetGenerate11Response) GetBody() []byte {
	if o == nil {
		return nil
	}
	return o.Body
}

func (o *GetGenerate11Response) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetGenerate11Response) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetGenerate11Response) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
