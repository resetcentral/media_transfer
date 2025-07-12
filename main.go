package main

import (
	"github.com/gin-gonic/gin"
	"github.com/resetcentral/media_transfer/handlers"
	"github.com/resetcentral/media_transfer/storage/orm"
)

func main() {
	orm.NewGormStorage()

	router := gin.Default()
	router.POST("download", handlers.PostDownload)
}
