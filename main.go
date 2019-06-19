package main

import (
    "fmt"
	"github.com/gin-gonic/gin"
	"github.com/xdy/gin/config/function"
)

func main(){
    startup.Startup()
    router := gin.Default()
    v1 := router.Group("v1")
    {
        step := 1
        for ; step > 0; step-- {
            fmt.Println(step)
        
            v1.GET("/login",func(c *gin.Context){
                fmt.Println("login")
                c.JSON(200,gin.H{
                    "status":"ok",
                })
            })

        }

    router.Run(":8080")
    }
}
