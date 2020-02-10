package routes

import (
	"giligili/controller"
	"giligili/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter() {
	r := gin.Default()

	r.Use(middleware.CurrentUser())

	r.GET("/ping", index)

	r.GET("/video/:v_id", controller.GetVideoInfo)
	r.GET("/videos/:offset/:limit", controller.GetListVideo)
	r.POST("/videos", controller.CreateVideo)
	r.PUT("/video/:v_id", controller.UpdateVideo)
	r.DELETE("/video/:v_id", controller.DeleteVideo)

	err := r.Run()
	if err != nil {
		panic(err)
	}
}

func index(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}
