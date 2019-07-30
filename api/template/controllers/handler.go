package template

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	parse "github.com/xdy/gin/utils/gin-parser"
	mysql "github.com/xdy/gin/utils/mysqlDbDriver"
)

/**************************************custom******************************************/

func (r *Routers) InsertOne(c *gin.Context) {
	data := parse.GetBody(c)
	lastId, affected := mysql.Insert(db, table, data)
	if affected >= 1 {
		params := map[string]interface{}{"_id": strconv.FormatInt(lastId, 10)}
		mJson, _ := json.Marshal(mysql.Get(db, table, params, model)[0])
		c.JSON(200, string(mJson))
	} else {
		c.String(400, "error in insert data")
	}
}
func (r *Routers) UpdateOne(c *gin.Context) {
	params := parse.GetParams(c)
	data := parse.GetBody(c)
	_, affected := mysql.Update(db, table, data, params, model)
	if affected >= 1 {
		//params := map[string]interface{}{"_id": strconv.FormatInt(lastId, 10)}
		mJson, _ := json.Marshal(mysql.Get(db, table, params, model)[0])
		c.JSON(200, string(mJson))
	} else {
		c.String(400, "error in insert data")
	}
}
func (r *Routers) FindOne(c *gin.Context) {
	params := parse.GetParams(c)
	fmt.Println(params)
	mJson, _ := json.Marshal(mysql.Get(db, table, params, model)[0])
	c.JSON(200, string(mJson))
}
func (r *Routers) FindAll(c *gin.Context) {
	query := parse.GetQuery(c)
	mJson, _ := json.Marshal(mysql.Get(db, table, query, model))
	c.JSON(200, string(mJson))
}
func (r *Routers) DeleteOne(c *gin.Context) {
	params := parse.GetParams(c)
	mJson, _ := json.Marshal(mysql.Get(db, table, params, model)[0])
	_, affected := mysql.Delete(db, table, params, model)
	if affected >= 1 {
		c.JSON(200, string(mJson))
	} else {
		c.JSON(400, "delete failed")
	}

}
