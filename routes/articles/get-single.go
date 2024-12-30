package articles

import (
	"register-backend/internal/database"
	"register-backend/types"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/gin-gonic/gin"
)

func GetSingle(c *gin.Context) {
	articleID := c.Param("articleID")
	query, err := database.Queries.Raw("get-article")
	if err != nil {
		c.Abort()
		_ = c.Error(err)
		return
	}

	var article types.Article
	err = pgxscan.Get(c, database.Pool, &article, query, articleID)
	if err != nil {
		// TODO: Enhance error handling by filtering out errors which indicate
		//   that a article has not been found
		c.AbortWithError(500, err)
		return
	}

	c.JSON(200, article)
}
