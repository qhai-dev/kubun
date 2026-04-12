package middleware

import (
	"github.com/gin-gonic/gin"
	"k8s.io/klog/v2"
)

func WithLogging() gin.HandlerFunc {
	return func(c *gin.Context) {
		keysAndValues := []interface{}{
			"HTTP", c.Request.Proto,
			"method", c.Request.Method,
			"srcIP", c.ClientIP(),
			"status", c.Writer.Status(),
		}
		klog.V(0).InfoDepth(1, keysAndValues...)
		c.Next()
	}
}
