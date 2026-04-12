package v1

import (
	"errors"

	"github.com/gin-gonic/gin"
	"k8s.io/klog/v2"
)

func NewRouterGroup(router *gin.RouterGroup) {
	r := router.Group("/v1")
	{
		r.GET("/test", func(c *gin.Context) {
			klog.ErrorS(errors.New("error log"), "test v")
			c.JSON(200, gin.H{
				"message": "test",
			})
		})
	}
}
