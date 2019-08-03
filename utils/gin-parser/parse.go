package parse

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
)

/*******************************数据库参数解析************************************/
func GetBody(c *gin.Context) map[string]interface{} {
	body, _ := ioutil.ReadAll(c.Request.Body)
	data := make(map[string]interface{})
	if err := json.Unmarshal(body, &data); err != nil {
		log.Println("data error", err)
	}
	return data
}

func GetParams(c *gin.Context) map[string]interface{} {
	params := make(map[string]interface{})
	for _, v := range c.Params {
		params[v.Key] = []string{v.Value}[0]
	}
	return params
}

func GetQuery(c *gin.Context) map[string]interface{} {
	query := make(map[string]interface{})
	for k, v := range c.Request.URL.Query() {
		fmt.Println(k, v)
		fmt.Println(v[0])
		query[k] = v[0]
	}
	return query
}

/************************************配置文件参数解析******************************************/
func CfgParse(path_ string) map[string]interface{} {
	mCfg := make(map[string]interface{})
	modelString, err := ioutil.ReadFile(path_)
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(modelString, &mCfg)
	return mCfg
}
