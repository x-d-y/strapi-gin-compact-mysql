package template

import (
	"fmt"
	"reflect"

	"github.com/gin-gonic/gin"
)

type Routers struct {
}

func ReturnRouters() Routers {
	var ruTest Routers
	return ruTest
}

func BindHandler() map[string]reflect.Value {
	var ruTest Routers
	type ControllerMapsType map[string]reflect.Value
	crMap := make(ControllerMapsType, 0)
	vf := reflect.ValueOf(&ruTest)
	mNum := vf.NumMethod()
	vft := vf.Type()
	for i := 0; i < mNum; i++ {
		mName := vft.Method(i).Name
		crMap[mName] = vf.Method(i)
	}
	return crMap
}

// var client *mongo.Client

// func LoadHandler(router *gin.Engine, routerHandeler map[string][]getRoutes.RouteInfo, client_ *mongo.Client, api string) {
// 	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
// 	fmt.Println(dir)
// 	client = client_
// 	crMap := BindHandler()
// 	for k, v := range routerHandeler {
// 		if k == api {
// 			groupRoute := router.Group(k)
// 			for _, u := range v {
// 				u_ := u
// 				Handler := []rune(u_.Handler)
// 				if Handler[0] >= 97 && Handler[0] <= 122 {
// 					Handler[0] -= 32
// 				}
// 				if u_.Method == "GET" {
// 					groupRoute.GET(u_.Path, func(c *gin.Context) {
// 						parms := []reflect.Value{reflect.ValueOf(c)}
// 						crMap[string(Handler)].Call(parms)
// 					})
// 				} else if u_.Method == "POST" {
// 					groupRoute.POST(u_.Path, func(c *gin.Context) {
// 						parms := []reflect.Value{reflect.ValueOf(c)}
// 						crMap[string(Handler)].Call(parms)
// 					})
// 				} else if u_.Method == "PUT" {
// 					groupRoute.PUT(u_.Path, func(c *gin.Context) {
// 						parms := []reflect.Value{reflect.ValueOf(c)}
// 						crMap[string(Handler)].Call(parms)
// 					})
// 				} else if u_.Method == "DELETE" {
// 					groupRoute.DELETE(u_.Path, func(c *gin.Context) {
// 						parms := []reflect.Value{reflect.ValueOf(c)}
// 						crMap[string(Handler)].Call(parms)
// 					})
// 				} else if u_.Method == "PATCH" {
// 					groupRoute.PATCH(u_.Path, func(c *gin.Context) {
// 						parms := []reflect.Value{reflect.ValueOf(c)}
// 						crMap[string(Handler)].Call(parms)
// 					})
// 				} else {
// 					fmt.Println("unmatch any method")
// 					return
// 				}
// 			}
// 		}
// 	}
// }

/**************************************custom******************************************/

func (r *Routers) InsertOne(c *gin.Context) {
	fmt.Println("this is test1Handler")
}
func (r *Routers) UpdateOne(c *gin.Context) {
	fmt.Println("this is test2Handler")
}
func (r *Routers) FindOne(c *gin.Context) {
	fmt.Println("this is test2Handler")
}
func (r *Routers) DeleteOne(c *gin.Context) {
	fmt.Println("this is test2Handler")
}
