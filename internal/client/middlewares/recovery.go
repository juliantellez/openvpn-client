package middlewares

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Recovery logs to stdout
func Recovery(context *gin.Context) {
	context.Next()

	defer func() {
		if r := recover(); r != nil {
			logrus.WithFields(logrus.Fields{
				"error": fmt.Sprintf("%s", r),
			}).Error("[ Client ]: Recovery")

			context.Writer.WriteHeader(http.StatusInternalServerError)
		}
	}()
}
