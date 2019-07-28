package template

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path"
	"reflect"

	"github.com/gin-gonic/gin"
	mysql "github.com/xdy/gin/utils/mysqlDbDriver"
)

type Routers struct {
}

var db *sql.DB
var table string
var model map[string]string

func Intatial(db_ *sql.DB, table_ string, apiFolder string) map[string]reflect.Value {
	db = db_
	table = table_
	modelJson := path.Join("api", apiFolder, "models", "models.json")
	modelString, err := ioutil.ReadFile(modelJson)
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(modelString, &model)
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
	data_ := make(map[string]interface{})
	data_["name"] = "xie"
	data_["salary"] = 111
	data_["deptId"] = 23
	mysql.Insert(db, table, data_)
	fmt.Println("this is test1Handler")
}
func (r *Routers) UpdateOne(c *gin.Context) {
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
	data_ := make(map[string]interface{})
	data_["name"] = "xie"
	data_["salary"] = 111
	data_["deptId"] = 23
	mysql.Get(db, table, data_, model)
	fmt.Println("this is test2Handler")
}
func (r *Routers) DeleteOne(c *gin.Context) {
	data_ := make(map[string]interface{})
	data_["name"] = "taosb"
	data_["salary"] = 0
	data_["deptId"] = 0
	mysql.Delete(db, table, data_, model)
	fmt.Println("this is test2Handler")
}
