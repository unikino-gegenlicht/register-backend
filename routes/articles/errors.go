package articles

import "github.com/wisdom-oss/common-go/v3/types"

var ErrDuplicateArticleName = types.ServiceError{
	Type:   "https://www.rfc-editor.org/rfc/rfc9110#section-15.5.1",
	Status: 400,
	Title:  "Duplicate Article Name",
	Detail: "The article you are trying to create/update uses a name already present in the database",
}

var ErrUnknownArticle = types.ServiceError{
	Type:   "https://www.rfc-editor.org/rfc/rfc9110#section-15.5.5",
	Status: 404,
	Title:  "Unknown Article",
	Detail: "The article ID does not match any known article",
}

var ErrUndeletableArticle = types.ServiceError{
	Type:   "https://www.rfc-editor.org/rfc/rfc9110#section-15.5.4",
	Status: 403,
	Title:  "Article Used In Transaction",
	Detail: "The article has already been used in a transaction and may not be deleted",
}
