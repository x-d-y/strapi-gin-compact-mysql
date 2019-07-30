package template

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	parse "github.com/xdy/gin/utils/gin-parser"
	mysql "github.com/xdy/gin/utils/mysqlDbDriver"
)

/**************************************custom******************************************/

func (r *Routers) InsertOne(c *gin.Context) {
	data := parse.GetBody(c)
	mysql.Insert(db, table, data)
}
func (r *Routers) UpdateOne(c *gin.Context) {
	params := parse.GetParams(c)
	data := parse.GetBody(c)
	mysql.Update(db, table, data, params, model)
	c.String(200, "pong")
}
func (r *Routers) FindOne(c *gin.Context) {
	params := parse.GetParams(c)
	mysql.Get(db, table, params, model)
}
func (r *Routers) FindAll(c *gin.Context) {
	query := parse.GetQuery(c)
	res := mysql.Get(db, table, query, model)
	fmt.Println(res)
	mJson, _ := json.Marshal(res)
	mString := string(mJson)
	fmt.Println(mString)
	c.JSON(200, mString)
}
func (r *Routers) DeleteOne(c *gin.Context) {
	params := parse.GetParams(c)
	mysql.Delete(db, table, params, model)
}
