package controllers

import (
	"github.com/gin-gonic/gin"
	"hole/pkgs/config"
	"net/http"
)

func Version() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, config.GetVersion())
	}
}
