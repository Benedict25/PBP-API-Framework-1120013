package main

import (
	c "Gin/Controllers"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	r := gin.Default()
	r.GET("/users", c.GetProducts)
	r.Run()
}
