package private

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/qhai-dev/kaka/apiserver/rest/router/private/v1"
)

func NewPrivateRouterGroup(router *gin.RouterGroup) {
	r := router.Group("/private")
	{
		v1.NewRouterGroup(r)
	}
}
