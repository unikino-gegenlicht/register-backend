package articles

import (
	"register-backend/internal/database"
	"register-backend/types"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/gin-gonic/gin"
)

func GetAll(c *gin.Context) {
	query, err := database.Queries.Raw("get-articles")
	if err != nil {
		c.Abort()
		_ = c.Error(err)
		return
	}

	var articles []types.Article
	err = pgxscan.Select(c, database.Pool, &articles, query)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	if len(articles) == 0 {
		c.Status(204)
		return
	}

	c.JSON(200, articles)
}
