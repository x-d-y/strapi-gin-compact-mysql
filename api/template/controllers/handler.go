package template

import (
	"fmt"
	"reflect"

	"github.com/gin-gonic/gin"
)

type Routers struct {
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
func (r *Routers) FindAll(c *gin.Context) {
	fmt.Println("this is test2Handler")
}
func (r *Routers) DeleteOne(c *gin.Context) {
	fmt.Println("this is test2Handler")
}
