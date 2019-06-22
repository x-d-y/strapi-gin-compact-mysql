package template

import (
	"fmt"
	"reflect"

	"github.com/gin-gonic/gin"
	getRoutes "github.com/xdy/gin/utils"
)

type Routers struct {
}

var cPath = "template"

func bindHandler() map[string]reflect.Value {
	var ruTest Routers
	type ControllerMapsType map[string]reflect.Value
	crMap := make(ControllerMapsType, 0)
	vf := reflect.ValueOf(&ruTest)
	mNum := vf.NumMethod()
	vft := vf.Type()
	fmt.Println("NumMethod:", mNum)
	for i := 0; i < mNum; i++ {
		mName := vft.Method(i).Name
		fmt.Println("index:", i, " MethodName:", mName)
		crMap[mName] = vf.Method(i)
	}
	return crMap
}

func LoadHandler(router *gin.Engine, routerHandeler map[string][]getRoutes.RouteInfo) {
	crMap := bindHandler()
	fmt.Println(crMap)
	for k, v := range routerHandeler {
		if k == cPath {
			groupRoute := router.Group(k)
			for _, u := range v {
				u_ := u
				Handler := []rune(u_.Handler)
				if Handler[0] >= 97 && Handler[0] <= 122 {
					Handler[0] -= 32
				}
				if u.Method == "GET" {
					groupRoute.GET(u.Path, func(c *gin.Context) {
						parms := []reflect.Value{reflect.ValueOf(c)}
						crMap[string(Handler)].Call(parms)
					})
				} else if u.Method == "POST" {
					groupRoute.POST(u.Path, func(c *gin.Context) {
						parms := []reflect.Value{reflect.ValueOf(c)}
						crMap[string(Handler)].Call(parms)
					})
				} else if u.Method == "PUT" {
					groupRoute.PUT(u.Path, func(c *gin.Context) {
						parms := []reflect.Value{reflect.ValueOf(c)}
						crMap[string(Handler)].Call(parms)
					})
				} else if u.Method == "DELETE" {
					groupRoute.DELETE(u.Path, func(c *gin.Context) {
						parms := []reflect.Value{reflect.ValueOf(c)}
						crMap[string(Handler)].Call(parms)
					})
				} else if u.Method == "PATCH" {
					groupRoute.PATCH(u.Path, func(c *gin.Context) {
						parms := []reflect.Value{reflect.ValueOf(c)}
						crMap[string(Handler)].Call(parms)
					})
				} else {
					return
				}
			}
		}

	}
}

func (r *Routers) Test1Hadler(c *gin.Context) {
	fmt.Println("this is test1Handler")
}

func (r *Routers) Test2Hadler(c *gin.Context) {
	fmt.Println("this is test2Handler")
}
