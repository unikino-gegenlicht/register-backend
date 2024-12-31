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
		if pgxscan.NotFound(err) {
			c.Abort()
			ErrUnknownArticle.Emit(c)
			return
		}
		c.Abort()
		_ = c.Error(err)
		return
	}

	c.JSON(200, article)
}
