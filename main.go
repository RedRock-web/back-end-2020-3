//Dte: 2020-03-27  1:47 AM
//Author : tuling

package main

import (
	"back-end/gee"
	"net/http"
)

func main() {
	r := gee.New()
	r.GET("/hello", func(c *gee.Context) {
		c.String(http.StatusOK, "Hello World!")
	})

	r.POST("/login", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":9999")
}
