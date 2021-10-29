package routes

import (
	"net/http"
	ctrl "slot/controllers"
	M "slot/middleware"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)



func RouterSetup() *gin.Engine {
	router := gin.Default()
	router.ForwardedByClientIP = true

	router.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Events booking app!")
	})
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://google.com"}
	config.MaxAge = 12 * time.Hour 
	
	

	router.Use(cors.New(config))
	router.Use(M.RateLimit())
	router.Use(M.Log)


	router.GET("/slots", ctrl.EventController{}.AvailableSlots)
	router.GET("/events", ctrl.EventController{}.BookedSlots)
	router.POST("/event", ctrl.EventController{}.CreateEvent)
	router.GET("/event/:eventId", ctrl.EventController{}.GetOneEvent)
	router.PUT("/event/:eventId", ctrl.EventController{}.UpdateEvent)
	router.DELETE("/event/:eventId", ctrl.EventController{}.DeleteEvent)
	router.POST("/upload/:eventId", ctrl.EventController{}.UploadFile)


	return router
}


