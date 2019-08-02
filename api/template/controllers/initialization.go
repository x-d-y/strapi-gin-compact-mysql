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

// func tableManagerCheck(item map[string]interface{}, modelString string) {
// 	model := make(map[string]interface{})
// 	model["_Table"] = "varchar(100)"
// 	model["_Property"] = "text"
// 	query := make(map[string]interface{})
// 	query["_Table"] = table
// 	modelString_ := mysql.Get(db, "tableManager", query, model)
// 	if len(modelString_) == 0 {
// 		goto insertTable
// 	}
// 	if modelString_[0].(map[string]interface{})["_Property"] == modelString {
// 		return
// 	} else {
// 		//fmt.Println(modelString_[0].(map[string]interface{})["_Property"], "~~~~~~")
// 		Mms_ := make(map[string]interface{})
// 		Mms := make(map[string]interface{})
// 		//fmt.Println(reflect.TypeOf(modelString_[0].(map[string]interface{})["_Property"]))
// 		json.Unmarshal([]byte(modelString_[0].(map[string]interface{})["_Property"].(string)), &Mms_)
// 		json.Unmarshal([]byte(modelString), &Mms)
// 		fmt.Println(Mms["item"])
// 		fmt.Println(Mms_["item"])
// 		for k, v := range Mms["item"].(map[string]interface{}) {
// 			fmt.Println(v.(map[string]interface{}), "~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
// 			fmt.Println(Mms_["item"].(map[string]interface{})[k], "@@@@@@@@@@@@@@@@@")
// 			// if v.(map[string]interface{}) == Mms_["item"].(map[string]interface{})[k].(map[string]interface{}) {
// 			// 	fmt.Println("ok")
// 			// } else {
// 			// 	fmt.Println("no ok")
// 			// }
// 		}
// 		return
// 	}
// insertTable:
// 	data := make(map[string]interface{})
// 	data["_Table"] = "template"
// 	data["_Property"] = string(modelString)
// 	mysql.Insert(db, "tableManager", data)
// 	add := make(map[string]interface{})
// 	remove := make(map[string]interface{})
// 	modify := make(map[string]interface{})
// 	_, _, _ = add, remove, modify

// 	return
// }

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
