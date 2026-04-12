package public

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/qhai-dev/kaka/apiserver/rest/router/public/v1"
)

func NewPublicRouterGroup(router *gin.RouterGroup) {
	r := router.Group("/public")
	{
		v1.NewRouterGroup(r)
	}
}
