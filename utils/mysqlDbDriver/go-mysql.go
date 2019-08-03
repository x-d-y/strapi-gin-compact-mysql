package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	parse "github.com/xdy/gin/utils/gin-parser"
)

var (
	tableColumn map[string]string
)

const tableManagerColumn string = `
		_Table varchar(100) NOT NULL,		
		_Property text NOT NULL,
		`

const tableManagerPK string = `CONSTRAINT tableManager_PK PRIMARY KEY (_Table)`
const pk string = "_id INT NOT NULL AUTO_INCREMENT,CONSTRAINT test2_PK PRIMARY KEY (_id)"

func sqlSelect(command string) string {
	createTable := `CREATE TABLE IF NOT EXISTS %s (
		%s
		%s
	)
	ENGINE=InnoDB
	DEFAULT CHARSET=utf8
	COLLATE=utf8_general_ci;
	`
	var return_ string
	switch command {
	case "createTable":
		return_ = createTable
	default:
		fmt.Println("can not parse the agrs")
	}
	return return_
}

func CheckCreatTable(db *sql.DB, table string, column string, pk_ string) string {
	if pk_ == "" {
		pk_ = pk
	}
	sqlStr := sqlSelect("createTable")
	sqlStr = fmt.Sprintf(sqlStr, table, column, pk_)
	res, err := db.Exec(sqlStr)
	if err != nil {
		fmt.Println(table)
		log.Fatal("create database failed", err)
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
		fmt.Println(lastId, rowCnt)
	}

	return sqlStr
}

func where(query map[string]interface{}, dataform map[string]interface{}) string {
	var queryString string
	i := 0
	for formk, _ := range dataform {
		for k, v := range query {
			if k == formk {
				queryString += k + "=" + formater(k, v, dataform)
				if i < len(query)-1 {
					queryString += " AND "
				}
				i++
			}
		}
	}

	return queryString
}

func formater(k string, v interface{}, pro map[string]interface{}) string {
	if k == "_id" {
		return v.(string)
	}
	if strings.Index(pro[k].(string), "varchar") > -1 {
		return_ := "'" + v.(string) + "'"
		return return_
	} else if strings.Index(pro[k].(string), "text") > -1 {
		return_ := "'" + v.(string) + "'"
		return return_
	} else if strings.Index(pro[k].(string), "int") > -1 {
		return v.(string) //strconv.Itoa(v.(int))
	} else if strings.Index(pro[k].(string), "bool") > -1 {
		return "true"
	} else {
		return "error"
	}
}

func Insert(db *sql.DB, table string, data map[string]interface{}) (int64, int64) {
	var key string
	var value []interface{}
	var questionMark string
	for k, v := range data {
		key = key + k + ","
		value = append(value, v)
		questionMark = questionMark + "?,"
	}
	key = key[0 : len(key)-1]
	questionMark = questionMark[0 : len(questionMark)-1]
retry:
	stmt, err := db.Prepare("INSERT INTO " + table + "(" + key + ") VALUES(" + questionMark + ");")
	if err != nil {
		errString := fmt.Sprintf("%s", err)
		index := strings.Index(errString, "doesn't exist")
		if index > -1 {
			column := tableColumn[table]
			CheckCreatTable(db, table, column, pk)
			goto retry
		}
		log.Fatal(err)
	}
	res, err := stmt.Exec(value...)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID=%d, affected=%d\n", lastId, rowCnt)
	return lastId, rowCnt
}

func Get(db *sql.DB, table string, query map[string]interface{}, dataform map[string]interface{}) []interface{} {
	queryString := where(query, dataform)
retry:
	rows, err := db.Query("select * from " + table + " WHERE " + queryString + ";")
	if err != nil {
		errString := fmt.Sprintf("%s", err)
		index := strings.Index(errString, "doesn't exist")
		if index > -1 {
			column := tableColumn[table]
			CheckCreatTable(db, table, column, pk)
			goto retry
		}
		log.Fatal(err)
	}

	defer rows.Close()
	columns, err := rows.Columns()

	if err != nil {
		log.Fatal(err)
	}

	values := make([]sql.RawBytes, len(columns))

	scanArgs := make([]interface{}, len(values))

	for i := range values {
		scanArgs[i] = &values[i]
	}
	var return_ []interface{}
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			log.Fatal(err)
		}
		var value string
		column := make(map[string]interface{})
		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}

			if strings.Index(dataform[columns[i]].(string), "varchar") > -1 {
				column[columns[i]] = value
			} else if strings.Index(dataform[columns[i]].(string), "text") > -1 {
				column[columns[i]] = value
			} else if strings.Index(dataform[columns[i]].(string), "int") > -1 {
				int_, err := strconv.Atoi(value)
				if err == nil {
					column[columns[i]] = int_
				}
			} else {
			}
		}
		return_ = append(return_, column)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return return_
}

func Update(db *sql.DB, table string, data map[string]interface{}, query map[string]interface{}, dataform map[string]interface{}) (int64, int64) {
	queryString := where(query, dataform)
	if queryString == "" {
		return -1, -1
	}
	var column string
	var value []interface{}
	for k, v := range data {
		column += k + "=?,"
		value = append(value, v)
	}
	column = column[0 : len(column)-1]
retry:
	stmt, err := db.Prepare("UPDATE " + table + " SET " + column + " WHERE " + queryString + ";")
	if err != nil {
		errString := fmt.Sprintf("%s", err)
		index := strings.Index(errString, "doesn't exist")
		if index > -1 {
			column := tableColumn[table]
			CheckCreatTable(db, table, column, pk)
			goto retry
		}
		log.Fatal(err)
	}

	res, err := stmt.Exec(value...)
	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID=%d, affected=%d\n", lastId, rowCnt)
	return lastId, rowCnt
}

func Delete(db *sql.DB, table string, query map[string]interface{}, dataform map[string]interface{}) (int64, int64) {
	queryString := where(query, dataform)
retry:
	stmt, err := db.Prepare("DELETE FROM " + table + " WHERE " + queryString + ";")
	if err != nil {
		errString := fmt.Sprintf("%s", err)
		index := strings.Index(errString, "doesn't exist")
		if index > -1 {
			column := tableColumn[table]
			CheckCreatTable(db, table, column, pk)
			goto retry
		}
		log.Fatal(err)
	}
	res, err := stmt.Exec()
	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	rowCnt, err := res.RowsAffected()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID=%d, affected=%d\n", lastId, rowCnt)
	return lastId, rowCnt
}

func ConnectClient() *sql.DB {
	dbCfg := path.Join("config", "environments", "developments", "database.json")
	mDbCfg := parse.CfgParse(dbCfg)
	dbInfo := fmt.Sprintf("%s:%s@/%s", mDbCfg["databaseAdmin"], os.Getenv(mDbCfg["databasePasswd"].(string)), mDbCfg["database"])
	db, err := sql.Open("mysql", dbInfo) //"root:123456@/test1")
	if err != nil {
		log.Fatal(err)
	}
	//CheckCreatTable(db, "tableManager", tableManagerColumn, tableManagerPK)
	return db
}

func TableColumn(tableColumn_ map[string]string) {
	tableColumn = tableColumn_
	//fmt.Println(tableColumn)
}
