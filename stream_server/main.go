package main

import (
	"os"
	"os/signal"
	"syscall"
	"video_server/stream_server/routers"
)

func main() {
	routers.HttpServerRun()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	routers.HttpServerStop()
}
