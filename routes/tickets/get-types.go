package tickets

import (
	"register-backend/internal/database"
	"register-backend/types"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/gin-gonic/gin"
)

func GetTypes(c *gin.Context) {
	query, err := database.Queries.Raw("get-ticket-types")
	if err != nil {
		c.Abort()
		_ = c.Error(err)
		return
	}

	var ticketTypes []types.TicketType
	err = pgxscan.Select(c, database.Pool, &ticketTypes, query)
	if err != nil {
		c.Abort()
		_ = c.Error(err)
		return
	}

	if len(ticketTypes) == 0 {
		c.Status(204)
		return
	}

	c.JSON(200, ticketTypes)
}
