package template

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"reflect"

	"github.com/gin-gonic/gin"
	mysql "github.com/xdy/gin/utils/mysqlDbDriver"
)

/**************************************custom******************************************/

func (r *Routers) InsertOne(c *gin.Context) {
	body, _ := ioutil.ReadAll(c.Request.Body)
	data := make(map[string]interface{})
	if err := json.Unmarshal(body, &data); err != nil {
		log.Println("data error", err)
	}
	mysql.Insert(db, table, data)
	fmt.Println("this is test1Handler")
}
func (r *Routers) UpdateOne(c *gin.Context) {
	fmt.Println(reflect.TypeOf(c.Params), "~~~~~~~~")
	data := make(map[string]interface{})
	data["name"] = "xie"
	data["salary"] = 111
	data["deptId"] = 23
	data_ := make(map[string]interface{})
	data_["name"] = "taosb"
	data_["salary"] = 0
	data_["deptId"] = 0
	mysql.Update(db, table, data_, data, model)
	fmt.Println("this is test2Handler")
}
func (r *Routers) FindOne(c *gin.Context) {
	fmt.Println("this is test2Handler")
}
func (r *Routers) FindAll(c *gin.Context) {
	query := c.Request.URL.Query()
	_ = query
	data_ := make(map[string]interface{})
	data_["name"] = "xie"
	data_["salary"] = 111
	data_["deptId"] = 23
	fmt.Println("this is test2Handler")

	mysql.Get(db, table, data_, model)
}
func (r *Routers) DeleteOne(c *gin.Context) {
	data_ := make(map[string]interface{})
	data_["name"] = "taosb"
	data_["salary"] = 0
	data_["deptId"] = 0
	mysql.Delete(db, table, data_, model)
	fmt.Println("this is test2Handler")
}
