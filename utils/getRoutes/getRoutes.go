package getRoutes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	//getRoutes "github.com/xdy/gin/utils"
)

type RouteInfo struct {
	Method  string `json:"method"`
	Path    string `json:"path"`
	Handler string `json:"handler"`
}

func parseRoutes(data []byte) []RouteInfo {
	var routeInfo_ interface{}
	var routeSlice []RouteInfo
	err_ := json.Unmarshal(data, &routeInfo_) //获取得到route.json内容
	if err_ != nil {
		fmt.Println(err_)
		return nil
	}
	item := routeInfo_.(map[string]interface{})
	for k, v := range item {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case int:
			fmt.Println(k, "is int", vv)
		case []interface{}:
			//fmt.Println(k, "is an array:")
			route := RouteInfo{}
			for _, u := range vv {
				u_ := u.(map[string]interface{})
				if method, ok := u_["method"]; ok {
					route.Method = method.(string)
				}
				if path, ok := u_["path"]; ok {
					route.Path = path.(string)
				}
				if handler, ok := u_["handler"]; ok {
					route.Handler = handler.(string)
				}
				routeSlice = append(routeSlice, route)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}
	// fmt.Println(routeSlice)
	return routeSlice
}

func Routes() map[string][]RouteInfo {
	apiFolder := "api"
	apis, _ := ioutil.ReadDir(apiFolder)
	groupRoutes := make(map[string][]RouteInfo)
	for _, api := range apis {
		path_ := apiFolder + "/" + api.Name()
		apiConfig, _ := ioutil.ReadDir(path_ + "/" + "config")
		for _, routesJson := range apiConfig {
			data, err := ioutil.ReadFile(path_ + "/" + "config" + "/" + routesJson.Name())
			//fmt.Println(routesJson.Name())
			if err != nil {
				return groupRoutes
			}
			res := parseRoutes(data)
			groupRoutes[api.Name()] = res
		}

	}
	return groupRoutes
}
