package startup

import (
	"github.com/gin-gonic/gin"
	template "github.com/xdy/gin/api/template/controllers"
	test "github.com/xdy/gin/api/test/controllers"
	getRoutes "github.com/xdy/gin/utils"
)

func Startup(router *gin.Engine) map[string][]getRoutes.RouteInfo {
	route := getRoutes.Routes() //解析路由
	//route_ := route.(map[string][]interface{})
	template.LoadHandler(router, route)
	test.LoadHandler(router, route)
	return route
}
