package handlers

import (
	"time"

	"github.com/gin-gonic/gin"
)

func ping(c *gin.Context) {
	currentTime := time.Now()
	c.String(200, "PING OK, Time: "+currentTime.String())
}
