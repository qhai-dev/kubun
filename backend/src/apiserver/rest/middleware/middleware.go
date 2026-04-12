package middleware

import (
	"github.com/gin-gonic/gin"
)

func SetupGlobalMiddleware(r *gin.Engine) {
	r.Use(gin.Recovery())
	r.Use(WithLogging())
}
