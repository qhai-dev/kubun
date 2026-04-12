package rest

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/qhai-dev/kaka/apiserver/rest/middleware"
	"github.com/qhai-dev/kaka/apiserver/rest/router"
	"k8s.io/klog/v2"
)

func Run() {
	ctx := context.Background()

	r := gin.New()
	// setup global middleware
	middleware.SetupGlobalMiddleware(r)
	// setup routerGroup
	router.SetupRouter(r)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r.Handler(),
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			klog.ErrorS(err, "HTTP start failed")
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	klog.Info("Start Shutdown Server")

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		klog.ErrorS(err, "Shutdown Server failed")
	}
	klog.Info("End Shutdown Server")
}
