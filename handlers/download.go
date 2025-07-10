package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/resetcentral/media_transfer/models"
	"github.com/resetcentral/media_transfer/storage"
)

func PostDownload(c *gin.Context) {
	var task models.Task
	storage.DB.CreateTask(task)
}
