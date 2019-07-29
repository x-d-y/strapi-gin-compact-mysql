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

var (
	db       *sql.DB
	table    string
	modelCfg map[string]interface{}
	model    map[string]interface{}
	column   string
)

func getModel(item map[string]interface{}) (map[string]interface{}, string) {
	model_ := make(map[string]interface{})
	createColumn := ""
	_ = createColumn
	for k, v := range item {
		createColumn_ := `%s %s(%s) %s,`
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

	return model_, createColumn
}

func Initialization(db_ *sql.DB, table_ string, apiFolder string) (map[string]reflect.Value, string) {
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
	mysql.CheckCreatTable(db, table, column)
	return crMap, column
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
