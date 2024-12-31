package articles

import (
	"errors"
	"fmt"
	"register-backend/internal/database"
	"register-backend/routes"
	"register-backend/types"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

func Edit(c *gin.Context) {
	var updateableFields struct {
		Name        *string  `json:"name"`
		Enabled     *bool    `json:"enabled"`
		MemberPrice *float64 `binding:"omitnil,gte=0"     json:"memberPrice"`
		GuestPrice  *float64 `binding:"omitnil,gte=0"     json:"guestPrice"`
		Color       *string  `binding:"omitempty,iscolor" json:"color"`
	}
	err := c.ShouldBind(&updateableFields)
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

	if updateableFields.Name != nil {
		article.Name = *updateableFields.Name
	}

	if updateableFields.Enabled != nil {
		article.Enabled = *updateableFields.Enabled
	}

	if updateableFields.MemberPrice != article.Prices.Members {
		article.Prices.Members = updateableFields.MemberPrice
	}

	if updateableFields.GuestPrice != nil {
		article.Prices.Guests = *updateableFields.GuestPrice
	}

	if updateableFields.Color != nil {
		article.Color = *updateableFields.Color
	}

	query, err = database.Queries.Raw("update-article")
	if err != nil {
		c.Abort()
		_ = c.Error(err)
		return
	}
	err = pgxscan.Get(c, database.Pool, &article, query,
		article.ID, article.Name, article.Enabled, article.Prices.Members,
		article.Prices.Guests, article.Color)
	if err != nil {
		c.Abort()
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

	c.JSON(200, article)
}
