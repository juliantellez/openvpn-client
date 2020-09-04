package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Bind attaches engine to handlers
func Bind(engine *gin.Engine) {
	engine.GET("/ping", ping)

	logrus.Info("[ Client ] Binding Handlers")
}
