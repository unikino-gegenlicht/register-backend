package articles

import (
	"errors"
	"fmt"
	"net/http"
	"register-backend/internal/database"
	"register-backend/routes"
	"register-backend/types"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
		var validationError validator.ValidationErrors
		if errors.As(err, &validationError) {
			res := routes.ErrInvalidRequestBody
			for _, fieldError := range validationError {
				if fieldError.Param() != "" {
					err := fmt.Errorf("unexpected input value in field '%s'. expected value matching: %s(%s), but got: %v", fieldError.Field(), fieldError.ActualTag(), fieldError.Param(), fieldError.Value())
					res.Errors = append(res.Errors, err)
					continue
				}
				err := fmt.Errorf("unexpected input value in field '%s'. expected value matching: %s, but got: %v", fieldError.Field(), fieldError.ActualTag(), fieldError.Value())
				res.Errors = append(res.Errors, err)
			}
			res.Emit(c)
			return
		}

		res := routes.ErrInvalidRequestBody
		res.Errors = append(res.Errors, err)
		res.Emit(c)
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
				ErrDuplicateArticleName.Emit(c)
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
