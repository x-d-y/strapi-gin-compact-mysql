package startup

import (
    "fmt"
    "io/ioutil"
)

func Startup() {
    apiFolder := "api"
    var route [100]string
    _ = route
    var groupMap map[string][]string
    _ = groupMap
    files, _ := ioutil.ReadDir(apiFolder)
    for num,file := range files{
        //var route_ = 
        fmt.Println(num)
        path_ := apiFolder+"/"+file.Name()
        //fmt.Println(path_)
        route_,_ := ioutil.ReadDir(path_+"/"+"config")
        for _,config := range route_{

            fmt.Println(config.Name())
        }
    }
}
