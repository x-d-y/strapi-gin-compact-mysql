package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

func where(query map[string]interface{}, dataform interface{}) string {
	s := reflect.ValueOf(dataform)
	typeOfT := s.Type()
	var queryString string
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		for k, v := range query {
			if k == typeOfT.Field(i).Name {
				queryString += typeOfT.Field(i).Name + "=" + formater(k, v, f.Type())
				if i < len(query)-1 {
					queryString += " AND "
				}
			}
		}
	}
	return queryString
}

func formater(k string, v interface{}, pro interface{}) string {
	string_ := reflect.ValueOf("").Type()
	int_ := reflect.ValueOf(0).Type()
	bool_ := reflect.ValueOf(true).Type()
	if pro == string_ {
		return_ := "'" + v.(string) + "'"
		return return_
	} else if pro == int_ {
		return strconv.Itoa(v.(int))
	} else if pro == bool_ {
		return "true"
	} else {
		return "error"
	}
}

func Insert(db *sql.DB, table string, data map[string]interface{}) {
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
	stmt, err := db.Prepare("INSERT INTO " + table + "(" + key + ") VALUES(" + questionMark + ");")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(value...)
	if err != nil {
		log.Fatal(err)
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID=%d, affected=%d\n", lastId, rowCnt)

}

func Get(db *sql.DB, table string, query map[string]interface{}, dataform interface{}) {
	queryString := where(query, dataform)
	rows, err := db.Query("select * from " + table + " WHERE " + queryString + ";")

	if err != nil {
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

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			log.Fatal(err)
		}
		var value string
		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			fmt.Println(columns[i], ":", value)
		}
		fmt.Println("--------------")
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
}

func Update(db *sql.DB, table string, data map[string]interface{}, query map[string]interface{}, dataform interface{}) {
	queryString := where(query, dataform)
	var column string
	var value []interface{}
	for k, v := range data {
		column += k + "=?,"
		value = append(value, v)
	}
	column = column[0 : len(column)-1]

	stmt, err := db.Prepare("UPDATE " + table + " SET " + column + " WHERE " + queryString + ";")
	if err != nil {
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
}

func Delete(db *sql.DB, table string, query map[string]interface{}, dataform interface{}) {
	queryString := where(query, dataform)
	stmt, err := db.Prepare("DELETE FROM " + table + " WHERE " + queryString + ";")
	if err != nil {
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
}

func MysqlClient() *sql.DB {
	table := "test"
	db, err := sql.Open("mysql", "root:123456@/test1")
	fmt.Println(err)
	if err != nil {
		log.Fatal(err)
	}
	type dataForm struct {
		name   string
		salary int
		deptId int
	}
	dataform := dataForm{}
	data_ := make(map[string]interface{})
	data_["name"] = "xie"
	data_["salary"] = 111
	data_["deptId"] = 23
	Get(db, table, data_, dataform)
	return db
}

func main_() {

	table := "test"
	db, err := sql.Open("mysql", "root:123456@/test1")
	fmt.Println(err)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	data_ := make(map[string]interface{})

	data_["name"] = "xie"
	data_["salary"] = 111
	data_["deptId"] = 23
	Insert(db, table, data_)

	data := make(map[string]interface{})
	data["name"] = "sbt"
	data["salary"] = 0
	type dataForm struct {
		name   string
		salary int
		deptId int
	}

	query := make(map[string]interface{})
	query["name"] = "xie"
	query["salary"] = 111
	dataform := dataForm{}
	Get(db, table, data_, dataform)
	Update(db, table, data, query, dataform)
	Delete(db, table, data, dataform)
}
