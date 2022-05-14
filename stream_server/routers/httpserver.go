package routers

import (
	"context"
	"log"
	"net/http"
	"time"
	"video_server/stream_server/common"
)

var (
	HttpSrvHandler *http.Server
)

func HttpServerRun() {
	r := InitRouter()
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
