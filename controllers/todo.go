package controllers

import (
	"net/http"
	"todo-app/database"
	"todo-app/models"

	"github.com/gin-gonic/gin"
)

// GetTodos retrieves all todos from the database
func GetTodos(c *gin.Context) {
	var todos []models.Todo
	if err := database.DB.Find(&todos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve todos"})
		return
	}
	c.JSON(http.StatusOK, todos)
}

// CreateTodo adds a new todo to the database
func CreateTodo(c *gin.Context) {
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := database.DB.Create(&todo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create todo"})
		return
	}
	c.JSON(http.StatusCreated, todo)
}

// UpdateTodo modifies an existing todo in the database
func UpdateTodo(c *gin.Context) {
	id := c.Param("id")
	var todo models.Todo
	if err := database.DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	database.DB.Save(&todo)
	c.JSON(http.StatusOK, todo)
}

// DeleteTodo removes a todo from the database
func DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	if err := database.DB.Delete(&models.Todo{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
