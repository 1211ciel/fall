package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		name := c.DefaultQuery("name", "ciel")
		c.JSON(http.StatusOK, gin.H{"msg": fmt.Sprint("hello,", name)})
	})
	r.POST("/add", func(c *gin.Context) {
		token := c.GetHeader("token")
		fmt.Println(token)
		var req struct {
			Name string `json:"name"`
		}
		err := c.ShouldBind(&req)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"msg": "err"})
		}
		c.JSON(http.StatusOK, gin.H{"msg": fmt.Sprintf("you added name is %s", req.Name)})
	})
	err := r.Run(":12011")
	if err != nil {
		panic(err.Error())
	}
}
