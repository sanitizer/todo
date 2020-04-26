package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/sanitizer/todo/dao"
	"github.com/sanitizer/todo/server"
)

func main() {
	dao.MigrateSchema()

	server := gin.Default()
	server.GET("/v1/todos", controller.GetAll)
	server.POST("/v1/todos", controller.Create)
	server.PUT("/v1/todos/:id", controller.Update)
	server.PATCH("/v1/todos/:id/complete", controller.Complete)
	server.DELETE("/v1/todos/:id", controller.Delete)
	server.Run(":8080")
}


