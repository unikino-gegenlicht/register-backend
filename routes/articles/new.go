package articles

import (
	"net/http"
	"register-backend/internal/database"
	"register-backend/types"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/gin-gonic/gin"
)

func New(c *gin.Context) {
	var article types.Article
	err := c.ShouldBindBodyWithJSON(article)
	if err != nil {
		c.AbortWithError(400, err)
		return
	}

	query, err := database.Queries.Raw("insert-article")
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	err = pgxscan.Get(c, database.Pool, &article, query, article.Name, article.Prices.Members, article.Prices.Guests, article.Color)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	c.JSON(http.StatusCreated, article)

}
