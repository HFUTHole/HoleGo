package controllers

import (
	"github.com/gin-gonic/gin"
	"hole/src/config"
	"net/http"
)

func Version() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, config.GetVersion())
	}
}
