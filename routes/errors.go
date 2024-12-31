package routes

import "github.com/wisdom-oss/common-go/v3/types"

var ErrInvalidRequestBody = types.ServiceError{
	Type:   "https://www.rfc-editor.org/rfc/rfc9110#section-15.5.1",
	Status: 400,
	Title:  "Invalid Request Body",
	Detail: "The request body failed validation. Please check your request and the documentation",
}
