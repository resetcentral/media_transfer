package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/resetcentral/media_transfer/storage"
)

func GetTasks(c *gin.Context) {
	tasks, err := storage.DB.GetTasks()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "An unexpected error occurred"})
		return
	}
	c.IndentedJSON(http.StatusOK, tasks)
}

func GetTaskByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		switch {

		}
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "id must be an integer"})
		return
	}
	task, err := storage.DB.GetTaskByID(id)
	if err != nil {

	}

	c.IndentedJSON(http.StatusOK, task)

}
