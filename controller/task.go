package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/takadev15/kanban-api/models"
	"github.com/takadev15/kanban-api/repository"
)

func (db Handlers) GetAllTask(c *gin.Context) {
	res, err := repository.GetAllTask(db.Connect)
	var result gin.H
	if err != nil {
		result = gin.H{
			"message": err,
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, result)
	}
	result = gin.H{
		"data": res,
	}
	c.JSON(http.StatusOK, result)
}

func (db Handlers) CreateTask(c *gin.Context) {
	var (
		task   models.Task
		result gin.H
	)
	if err := c.ShouldBindJSON(&task); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	err := repository.CreateTask(&task, db.Connect)
	if err != nil {
		result = gin.H{
			"message": err,
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, result)
	}
	result = gin.H{
		"id":          task.ID,
		"title":       task.Title,
		"status":      task.Status,
		"description": task.Description,
		"user_id":     task.UserId,
		"category_id": task.CategoryId,
		"created_at":  task.CreatedAt,
	}
	c.JSON(http.StatusOK, result)
}

func (db Handlers) UpdateTask(c *gin.Context) {
	var (
		task   models.Task
		result gin.H
	)
	if err := c.ShouldBindJSON(&task); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	taskId, _ := strconv.Atoi(c.Param("id"))
	_, err := repository.UpdateTask(taskId, &task, db.Connect)
	if err != nil {
		result = gin.H{
			"message": err,
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, result)
	}
	result = gin.H{
		"data": task,
	}
	c.JSON(http.StatusCreated, result)
}

func (db Handlers) DeleteTask(c *gin.Context) {
	taskId := c.Param("id")
	id, _ := strconv.Atoi(taskId)
	err := repository.DeleteTask(id, db.Connect)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "task has been succesfully deleted",
	})
}
