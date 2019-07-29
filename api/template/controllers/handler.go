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
var modelCfg map[string]interface{}
var model map[string]interface{}
var column string

//tableC string, model map[string]string
func checkForTable(column string) {
	//a := "name varchar(100) NULL,salary int NULL,deptId int NULL,"
	sql := fmt.Sprintf(mysql.CreatTable(), "test2", column)
	fmt.Println(sql)
	_, err := db.Exec(sql)
	if err != nil {
		fmt.Println("create database failed")
	}
}

func getModel(item map[string]interface{}) (map[string]interface{}, string) {
	fmt.Println(item)
	model_ := make(map[string]interface{})
	createColumn := ""
	_ = createColumn
	for k, v := range item {
		createColumn_ := `%s %s(%s) %s,`
		fmt.Println(k, v)
		dataType := ""
		dataLength := ""
		notNull := ""
		v__ := v.(map[string]interface{})
		for k_, v_ := range v__ {
			switch k_ {
			case "dataType":
				model_[k] = v_
				dataType = v_.(string)
			case "dataLength":
				dataLength = v_.(string)
			case "notNull":
				notNull = v_.(string)
			default:
				fmt.Println("can not parse the agrs")
			}
		}
		createColumn_ = fmt.Sprintf(createColumn_, k, dataType, dataLength, notNull)
		createColumn += createColumn_

	}
	fmt.Println(createColumn, "!!!!!!!!!!!!!!!!!!!!")

	return model_, createColumn
}

func Intatial(db_ *sql.DB, table_ string, apiFolder string) map[string]reflect.Value {

	db = db_
	table = table_
	modelJson := path.Join("api", apiFolder, "models", "models.json")
	modelString, err := ioutil.ReadFile(modelJson)
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(modelString, &modelCfg)
	item := modelCfg["item"].(map[string]interface{})
	model, column = getModel(item)
	fmt.Println(model, "~~~~~~")
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
	checkForTable(column)
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
