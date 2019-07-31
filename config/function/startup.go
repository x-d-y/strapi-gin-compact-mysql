package startup

import (
	"reflect"

	"github.com/gin-gonic/gin"
	getRoutes "github.com/xdy/gin/utils/getRoutes"
	mysql "github.com/xdy/gin/utils/mysqlDbDriver"
)

func Startup(router *gin.Engine) map[string][]getRoutes.RouteInfo {
	route, _ := getRoutes.Routes()                     //解析路由,获取每个api
	db := mysql.ConnectClient()                        //连接数据库
	tableColumn, handlerSlice := apiInitialization(db) //初始化每个api
	for k, v := range handlerSlice {
		//fmt.Println(k, v)
		routerFuncLoader(router, route, k, v.(map[string]reflect.Value)) //将每个api的函数与路由绑定
	}
	mysql.TableColumn(tableColumn) //将不同的api对应的column保留，以防在运行过程中数据表被删除，自动恢复数据表，但是数据无法恢复
	return route
}
