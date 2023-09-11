// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"petstore/internal/sdk/pkg/models/shared"
)

type CreatePetsResponse struct {
	ContentType string
	// unexpected error
	Error       *shared.Error
	StatusCode  int
	RawResponse *http.Response
}
