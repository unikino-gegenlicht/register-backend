package tickets

import "github.com/gin-gonic/gin"

func ConvertReservation(c *gin.Context) {
	_ = c.Param("reservationID")

}
