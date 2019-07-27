package startup

import (
	"fmt"
	"reflect"

	"github.com/gin-gonic/gin"
	test "github.com/xdy/gin/api/test/controllers"
	getRoutes "github.com/xdy/gin/utils/getRoutes"
	mongodb "github.com/xdy/gin/utils/mongoDbDriver"
	"go.mongodb.org/mongo-driver/mongo"
)

func Startup(router *gin.Engine) map[string][]getRoutes.RouteInfo {
	route := getRoutes.Routes()        //解析路由
	client := mongodb.MongodbInitial() //链接数据库
	handler := test.BindHandler(client)
	routerFuncLoader(router, route, client, "test", handler)
	return route
}

func routerFuncLoader(router *gin.Engine, routerHandeler map[string][]getRoutes.RouteInfo, client_ *mongo.Client, api string, crMap map[string]reflect.Value) {
	for k, v := range routerHandeler {
		fmt.Println(k, v)
		if k == api {
			groupRoute := router.Group(k)
			for _, u := range v {
				u_ := u
				Handler := []rune(u_.Handler)
				if Handler[0] >= 97 && Handler[0] <= 122 {
					Handler[0] -= 32
				}
				if u_.Method == "GET" {
					groupRoute.GET(u_.Path, func(c *gin.Context) {
						parms := []reflect.Value{reflect.ValueOf(c)}
						crMap[string(Handler)].Call(parms)
					})
				} else if u_.Method == "POST" {
					groupRoute.POST(u_.Path, func(c *gin.Context) {
						parms := []reflect.Value{reflect.ValueOf(c)}
						crMap[string(Handler)].Call(parms)
					})
				} else if u_.Method == "PUT" {
					groupRoute.PUT(u_.Path, func(c *gin.Context) {
						parms := []reflect.Value{reflect.ValueOf(c)}
						crMap[string(Handler)].Call(parms)
					})
				} else if u_.Method == "DELETE" {
					groupRoute.DELETE(u_.Path, func(c *gin.Context) {
						parms := []reflect.Value{reflect.ValueOf(c)}
						crMap[string(Handler)].Call(parms)
					})
				} else if u_.Method == "PATCH" {
					groupRoute.PATCH(u_.Path, func(c *gin.Context) {
						parms := []reflect.Value{reflect.ValueOf(c)}
						crMap[string(Handler)].Call(parms)
					})
				} else {
					fmt.Println("unmatch any method")
					return
				}
			}
		}
	}
}
