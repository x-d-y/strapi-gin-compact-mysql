package startup

import (
	getRoutes "github.com/xdy/gin/utils"
)

func Startup() {

	getRoutes.Routes() //解析路由

}
