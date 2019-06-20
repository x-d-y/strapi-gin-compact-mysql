package startup

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type routeInfo struct {
	Method  interface{} `json:"method"`
	Path    interface{} `json:"path"`
	Handler interface{} `json:"handler"`
}

func Startup() {
	apiFolder := "api"
	var route [100]string
	_ = route
	var groupMap map[string][]string
	_ = groupMap
	files, _ := ioutil.ReadDir(apiFolder)
	for num, file := range files {
		fmt.Println(num)
		path_ := apiFolder + "/" + file.Name()
		route_, _ := ioutil.ReadDir(path_ + "/" + "config")
		for _, config := range route_ {
			data, err := ioutil.ReadFile(path_ + "/" + "config" + "/" + config.Name())
			if err != nil {
				return
			}
			var routeInfo_ interface{}
			err_ := json.Unmarshal(data, &routeInfo_)
			if err_ != nil {
				fmt.Println(err_)
				return
			}
			fmt.Println(routeInfo_)
			m := routeInfo_.(map[string]interface{})
			for k, v := range m {
				switch vv := v.(type) {

				case string:
					fmt.Println(k, "is string", vv)
				case int:
					fmt.Println(k, "is int", vv)
				case []interface{}:
					fmt.Println(k, "is an array:")
					for i, u := range vv {
						u_ := u.(map[string]interface{})
						if method, ok := u_["method"]; ok {
							fmt.Println(method.(string))
						}
						fmt.Println(i, u)
						//fmt.Println(routeAndHandler)
					}
				default:
					fmt.Println(k, "is of a type I don't know how to handle")
				}
			}
		}
	}
}
