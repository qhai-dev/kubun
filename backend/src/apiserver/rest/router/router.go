package router

import (
	"github.com/gin-gonic/gin"
	"github.com/qhai-dev/kaka/apiserver/rest/router/private"
	"github.com/qhai-dev/kaka/apiserver/rest/router/public"
)

func SetupRouter(router *gin.Engine) {
	r := router.Group("/api")
	{
		public.NewPublicRouterGroup(r)
		private.NewPrivateRouterGroup(r)
	}
}
