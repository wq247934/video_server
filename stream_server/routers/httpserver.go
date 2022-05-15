package routers

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
	"video_server/stream_server/common"
	"video_server/stream_server/middleware"
)

var (
	HttpSrvHandler *http.Server
)

func HttpServerRun() {
	gin.SetMode("debug")
	r := InitRouter(middleware.Cors())
	HttpSrvHandler = &http.Server{
		Addr:         common.HttpAddr,
		Handler:      r,
		ReadTimeout:  time.Duration(common.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(common.WriteTimeout) * time.Second,
	}
	go func() {
		log.Printf(" [INFO] HttpServerRun:%s\n", common.HttpAddr)
		if err := HttpSrvHandler.ListenAndServe(); err != nil {
			log.Fatalf(" [ERROR] HttpServerRun:%s err:%v\n", common.HttpAddr, err)
		}
	}()
}

func HttpServerStop() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := HttpSrvHandler.Shutdown(ctx); err != nil {
		log.Fatalf(" [ERROR] HttpServerStop err:%v\n", err)
	}
	log.Printf(" [INFO] HttpServerStop stopped\n")
}
