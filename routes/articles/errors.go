package articles

import "github.com/wisdom-oss/common-go/v3/types"

var ErrDuplicateArticle = types.ServiceError{
	Type:   "https://www.rfc-editor.org/rfc/rfc9110#section-15.5.1",
	Status: 400,
	Title:  "Duplicate Article",
	Detail: "The article you are trying to create already exists. Please choose another name.",
}

var ErrUnknownArticle = types.ServiceError{
	Type:   "https://www.rfc-editor.org/rfc/rfc9110#section-15.5.5",
	Status: 404,
	Title:  "Unknown Article",
	Detail: "The article ID does not match any known article",
}
