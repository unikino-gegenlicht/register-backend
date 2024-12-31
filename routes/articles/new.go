package articles

import (
	"errors"
	"net/http"
	"register-backend/internal/database"
	"register-backend/types"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgconn"

	"github.com/jackc/pgerrcode"
)

// DefaultArticleColor is the default color applied to the article cards in the
// frontend/app if no explicit color has been set during the creation
const DefaultArticleColor = "#00a2ff"

func New(c *gin.Context) {
	var article types.Article
	article.Color = DefaultArticleColor

	err := c.ShouldBind(&article)
	if err != nil {
		c.Abort()
		_ = c.Error(err)
		return
	}

	query, err := database.Queries.Raw("insert-article")
	if err != nil {
		c.Abort()
		_ = c.Error(err)
		return
	}

	err = pgxscan.Get(c, database.Pool, &article, query, article.Name, article.Prices.Members, article.Prices.Guests, article.Color)
	if err != nil {
		if pgxscan.NotFound(err) {
			panic("no result found after inserting")
		}
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case pgerrcode.UniqueViolation:
				c.Abort()
				ErrDuplicateArticle.Emit(c)
				return
			default:
				c.Abort()
				_ = c.Error(err)
				return
			}
		}
	}

	c.JSON(http.StatusCreated, article)
}
