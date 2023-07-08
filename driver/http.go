package driver

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type httpConfig struct {
	port int
}

func LoadHttpConfig(port int) httpConfig {
	return httpConfig{
		port: port,
	}
}

func RunGinHttpServer(ctx context.Context, router *gin.Engine, config httpConfig) (cleanup func(), err error) {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.port),
		Handler: router,
	}

	go func() {
		log.Printf("Run service with %s mode", gin.Mode())

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("listen: %s\n", err)
		}
	}()

	if err != nil {
		return
	}

	cleanup = func() {
		if err := srv.Shutdown(ctx); err != nil {
			log.Println("Server Shutdown:", err)
			return
		}

		fmt.Println("shutdown http server success")
	}

	return
}
