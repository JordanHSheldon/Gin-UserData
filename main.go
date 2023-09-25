package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.POST("/user_data", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"IsDisabled": "false",
			"username":   "nadroj",
			"password":   "password",
			"email":      "nadroj@gmail.com",
			"settingsid": "1",
		})
	})
	r.Run()
}
