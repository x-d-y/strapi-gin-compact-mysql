package startup

import (
	"database/sql"
	"reflect"

	template "github.com/xdy/gin/api/template/controllers"
)

var (
	db           *sql.DB
	handlerSlice map[string]interface{}
	tableColumn  map[string]string
	handler      map[string]reflect.Value
	column       string
)

func apiInitialization(db_ *sql.DB) (map[string]string, map[string]interface{}) {
	db = db_
	handlerSlice = make(map[string]interface{})
	tableColumn = make(map[string]string)
	handler = make(map[string]reflect.Value)
	column = ""
	jsWriter()
	return tableColumn, handlerSlice
}

func handlerColumn(table string) {
	handlerSlice[table] = handler
	tableColumn[table] = column
}

func jsWriter() {
	handler, column = template.Initialization(db, "template", "template") //初始化每个api并且将每个api的column以sql创建的形式返回
	handlerColumn("template")
}
