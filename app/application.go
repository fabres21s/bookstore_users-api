package app

import (
	"github.com/fabres21s/bookstore_users-api/logger"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()

	logger.Info("about the start application")
	router.Run(":8080")
}
