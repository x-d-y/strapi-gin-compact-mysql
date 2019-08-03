package main

import (
	"github.com/gin-gonic/gin"
	startup "github.com/xdy/gin/config/function"
)

func main() {
	router := gin.Default()
	port := startup.Startup(router)
	router.Run(":" + port["serverPort"].(string))

}
