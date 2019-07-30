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
	params := make(map[string]interface{})
	for k, v := range c.Params {
		fmt.Println(k, reflect.TypeOf(v))
		params[v.Key] = []string{v.Value}[0]
	}
	data := make(map[string]interface{})
	body, _ := ioutil.ReadAll(c.Request.Body)
	if err := json.Unmarshal(body, &data); err != nil {
		log.Println("data error", err)
	}
	mysql.Update(db, table, data, params, model)
	fmt.Println("this is test2Handler")
	c.String(200, "pong")
}
func (r *Routers) FindOne(c *gin.Context) {
	params := make(map[string]interface{})
	for k, v := range c.Params {
		fmt.Println(k, reflect.TypeOf(v))
		params[v.Key] = []string{v.Value}[0]
	}
	mysql.Get(db, table, params, model)

	fmt.Println("this is test2Handler")
}
func (r *Routers) FindAll(c *gin.Context) {
	query := make(map[string]interface{})
	for k, v := range c.Request.URL.Query() {
		fmt.Println(k, v)
		fmt.Println(v[0])
		query[k] = v[0]
	}
	mysql.Get(db, table, query, model)
}
func (r *Routers) DeleteOne(c *gin.Context) {
	params := make(map[string]interface{})
	for k, v := range c.Params {
		fmt.Println(k, reflect.TypeOf(v))
		params[v.Key] = []string{v.Value}[0]
	}
	mysql.Delete(db, table, params, model)
	fmt.Println("this is test2Handler")
}
