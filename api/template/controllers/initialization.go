package template

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path"
	"reflect"

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

func Initialization(db_ *sql.DB, table_ string, apiFolder string, handlerSlice map[string]interface{}, tableColumn map[string]string) (map[string]interface{}, map[string]string) {
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
	crMap := make(map[string]reflect.Value, 0)
	vf := reflect.ValueOf(&ruTest)
	mNum := vf.NumMethod()
	vft := vf.Type()
	for i := 0; i < mNum; i++ {
		mName := vft.Method(i).Name
		crMap[mName] = vf.Method(i)
	}
	mysql.CheckCreatTable(db, table, column, "")
	handlerSlice[table] = crMap
	tableColumn[table] = column
	return handlerSlice, tableColumn
}

func getModel(item map[string]interface{}) (map[string]interface{}, string) {
	model_ := make(map[string]interface{})
	createColumn := ""
	_ = createColumn
	for k, v := range item {
		createColumn_ := `%s %s %s,`
		Type := ""
		Null := ""
		v__ := v.(map[string]interface{})
		for k_, v_ := range v__ {
			switch k_ {
			case "Type":
				model_[k] = v_
				Type = v_.(string)
			case "Null":
				Null = v_.(string)
			case "Key":
				//Key
			case "Default":
				//Default
			case "Extra":
				//Extra
			default:
				fmt.Println("can not parse the agrs")
			}
		}
		createColumn_ = fmt.Sprintf(createColumn_, k, Type, Null)
		createColumn += createColumn_
	}
	model_["_id"] = "int"
	return model_, createColumn
}
