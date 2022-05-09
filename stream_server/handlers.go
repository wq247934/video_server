package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"os"
	"time"
)

func StreamHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	vid := p.ByName("vid-id")
	video, err := os.Open(VideoPath + vid + ".mp4")
	if err != nil {
		sendErrorResponse(w, http.StatusNotFound, "File not found!")
		return
	}
	w.Header().Set("content-type", "video/mp4")
	http.ServeContent(w, r, "...", time.Now(), video)
	defer video.Close()
}

func UploadHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}
