package controllers

import (
	"belajar-go/config"
	"belajar-go/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	var todo []model.Todo
	config.DB.Find(&todo)
	c.JSON(http.StatusOK, todo)
}

func Create(c *gin.Context) {
	var todo model.Todo
	err := c.BindJSON(&todo)
	if err != nil {
		panic(err)
	}
	res := config.DB.Create(&todo)
	if res.Error != nil {
		c.JSON(http.StatusBadRequest, "failed create todo")
		return
	}
	c.JSON(http.StatusOK, todo)
}

func Show(c *gin.Context) {
	id := c.Params.ByName("id")
	var todo model.Todo
	config.DB.First(&todo, id)
	c.JSON(http.StatusOK, todo)
}

func Update(c *gin.Context) {
	id := c.Params.ByName("id")
	var todo model.Todo
	err := config.DB.First(&todo, id).Error
	if err != nil {
		panic(err)
	}

	if err := c.BindJSON(&todo); err != nil {
		panic(err)
	}

	config.DB.Updates(&todo)
	c.JSON(http.StatusOK, todo)
}

func Delete(c *gin.Context) {
	id := c.Params.ByName("id")
	var todo model.Todo
	err := config.DB.First(&todo, id).Error
	if err != nil {
		panic(err)
	}
	config.DB.Delete(&todo)
	c.JSON(http.StatusOK, todo)
}