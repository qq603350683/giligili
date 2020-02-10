package middleware

import (
	"github.com/gin-gonic/gin"
)

func CurrentUser() gin.HandlerFunc {
	return func(c *gin.Context) {

		//c.JSON(202, serializer.JSON{
		//	Status:  http.StatusOK,
		//	Message: "OK",
		//	Data:    time.Now(),
		//})

		//panic("GG")

		//c.Next()

		//c.JSON(203, serializer.JSON{
		//	Status:  http.StatusOK,
		//	Message: "OK",
		//	Data:    nil,
		//})
	}
}
