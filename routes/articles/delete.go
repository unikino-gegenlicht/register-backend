package articles

import (
	"errors"
	"register-backend/internal/database"
	"register-backend/types"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

func Delete(c *gin.Context) {
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
		c.Abort()
		if pgxscan.NotFound(err) {
			ErrUnknownArticle.Emit(c)
			return
		}
		_ = c.Error(err)
		return
	}

	query, err = database.Queries.Raw("delete-article")
	if err != nil {
		c.Abort()
		_ = c.Error(err)
		return
	}

	_, err = database.Pool.Exec(c, query, articleID)
	if err != nil {
		c.Abort()
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case pgerrcode.ForeignKeyViolation:
				ErrUndeletableArticle.Emit(c)
				return
			default:
				_ = c.Error(err)
				return
			}
		}
		_ = c.Error(err)
		return
	}

	c.Status(204)
}
