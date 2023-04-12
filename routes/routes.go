package routes

import(
	"github.com/joy/go-ecommerce/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine)  {
	incomingRoutes.POST("/users/signup")
	incomingRoutes.POST("/users/login")
}