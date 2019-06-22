package test

import (
	"fmt"
	"reflect"

	"github.com/gin-gonic/gin"
	getRoutes "github.com/xdy/gin/utils/getRoutes"
	"go.mongodb.org/mongo-driver/mongo"
)

type Routers struct {
}

var cPath = "test"

func bindHandler() map[string]reflect.Value {
	var ruTest Routers
	type ControllerMapsType map[string]reflect.Value
	crMap := make(ControllerMapsType, 0)
	vf := reflect.ValueOf(&ruTest)
	mNum := vf.NumMethod()
	vft := vf.Type()
	//fmt.Println("NumMethod:", mNum)
	for i := 0; i < mNum; i++ {
		mName := vft.Method(i).Name
		//fmt.Println("index:", i, " MethodName:", mName)
		crMap[mName] = vf.Method(i)
	}
	return crMap
}

var client *mongo.Client

func LoadHandler(router *gin.Engine, routerHandeler map[string][]getRoutes.RouteInfo, client_ *mongo.Client) {
	client = client_
	crMap := bindHandler()
	//fmt.Println(crMap)
	for k, v := range routerHandeler {
		if k == cPath {
			groupRoute := router.Group(k)
			for i, u := range v {
				u_ := u
				fmt.Println(i)
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

/**************************************custom******************************************/

type Model struct {
	Name string
	Age  int
	City string
}

func (r *Routers) InsertOne(c *gin.Context) {
	buf := make([]byte, 1024)
	n, _ := c.Request.Body.Read(buf)
	fmt.Println(string(buf[0:n]))
	//fmt.Println(c.PostForm())
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
