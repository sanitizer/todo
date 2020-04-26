package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sanitizer/todo/dao"
	"github.com/sanitizer/todo/model"
	"net/http"
	"strconv"
)

func GetAll(context *gin.Context) {
	context.JSON(
		http.StatusOK,
		dao.GetAll(context.Query("status")),
	)
}

func Delete(context *gin.Context) {
	id, _ := strconv.ParseInt(context.Param("id"), 10, 64)
	dao.Delete(id)
	context.String(204, "")
}

func Create(context *gin.Context) {
	rawBody, _ := context.GetRawData()
	body := new(model.Todo)
	body.FromJson(rawBody)

	context.JSON(
		http.StatusCreated,
		dao.Create(*body),
	)
}

func Update(context *gin.Context) {
	rawBody, _ := context.GetRawData()
	body := new(model.Todo)
	body.FromJson(rawBody)
	id, _ := strconv.ParseInt(context.Param("id"), 10, 64)
	context.JSON(
		http.StatusOK,
		dao.Update(id, *body),
	)
}

func Complete(context *gin.Context) {
	id, _ := strconv.ParseInt(context.Param("id"), 10, 64)
	context.JSON(
		http.StatusOK,
		dao.Complete(id),
	)
}
