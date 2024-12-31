package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/wisdom-oss/common-go/v3/types"
)

// Panic is the base types.ServiceError used if the ErrorHandler catches a panic
// during the handling of a request.
var Panic = &types.ServiceError{
	Type:   "https://pkg.go.dev/builtin#panic",
	Status: 500,
	Title:  "Internal Panic",
	Detail: "The service encountered a panic state during the handling of your request.",
}

func RecoveryHandler(c *gin.Context, err any) {
	response := Panic
	response.Errors = []error{fmt.Errorf("%v", err)}
	c.Header("Content-Type", "application/problem+json; charset=utf-8")
	c.AbortWithStatusJSON(500, response)
}
