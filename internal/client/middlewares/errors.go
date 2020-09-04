package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Errors logs to stdout
func Errors(context *gin.Context) {
	context.Next()

	for _, err := range context.Errors {
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Error("[ Client ] Handler Error")
	}
}
