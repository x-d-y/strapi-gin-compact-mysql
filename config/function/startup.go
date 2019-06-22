package startup

import (
	"github.com/gin-gonic/gin"
	template "github.com/xdy/gin/api/template/controllers"
	test "github.com/xdy/gin/api/test/controllers"
	getRoutes "github.com/xdy/gin/utils/getRoutes"
	mongodb "github.com/xdy/gin/utils/mongoDbDriver"
)

func Startup(router *gin.Engine) map[string][]getRoutes.RouteInfo {
	route := getRoutes.Routes() //解析路由
	//route_ := route.(map[string][]interface{})
	client := mongodb.MongodbInitial()
	template.LoadHandler(router, route, client)
	test.LoadHandler(router, route, client)
	return route
}
