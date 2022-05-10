package main

import (
	"video_server/stream_server/routers"
)

func main() {
	r := routers.SetRouters()
	err := r.Run(":9000")
	if err != nil {
		panic(err)
	}

}
