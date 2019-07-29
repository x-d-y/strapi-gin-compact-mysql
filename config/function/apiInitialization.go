package startup

import (
	"database/sql"

	template "github.com/xdy/gin/api/template/controllers"
	//test "github.com/xdy/gin/api/test/controllers"
)

var (
	db           *sql.DB
	handlerSlice map[string]interface{}
	tableColumn  map[string]string
)

func apiInitialization(db_ *sql.DB) (map[string]string, map[string]interface{}) {
	db = db_
	handlerSlice = make(map[string]interface{})
	tableColumn = make(map[string]string)
	jsWriter()
	return tableColumn, handlerSlice
}

func jsWriter() {
	handlerSlice, tableColumn = template.Initialization(db, "template", "template", handlerSlice, tableColumn) //初始化每个api并且将每个api的column以sql创建的形式返回
	//handlerSlice, tableColumn = test.Initialization(db, "test", "test", handlerSlice, tableColumn)             //初始化每个api并且将每个api的column以sql创建的形式返回

}
