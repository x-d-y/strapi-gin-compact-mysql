package main

import (
	"github.com/gin-gonic/gin"
	startup "github.com/xdy/gin/config/function"
)

func main() {
	router := gin.Default()
	routerHandeler := startup.Startup(router)
	_ = routerHandeler
	router.Run(":8080")

}
