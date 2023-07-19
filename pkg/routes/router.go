package routes

import (
	"os"

	"github.com/akshayUr04/google-translator/pkg/controller"
	"github.com/gin-gonic/gin"
)

func Routing() {
	r := gin.Default()
	r.POST("/translate", controller.Translate)
	r.Run(":" + os.Getenv("PORT"))
}
