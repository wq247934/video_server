package controllers

import (
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
	"video_server/stream_server/common"
)

func StreamHandler(c *gin.Context) {
	vid := c.Param("vid")
	video, err := os.Open(common.VideoPath + vid)
	if err != nil {
		//sendErrorResponse(w, http.StatusNotFound, "File not found!")
		return
	}
	c.Writer.Header().Set("content-type", "video/mp4")
	http.ServeContent(c.Writer, c.Request, "...", time.Now(), video)
	defer video.Close()
}

func UploadHandler(c *gin.Context) {
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, common.MaxUploadSize)
	if err := c.Request.ParseMultipartForm(common.MaxUploadSize); err != nil {
		log.Println("File is to big!")
		log.Println(err)
		common.SendErrorResponse(c.Writer, http.StatusBadRequest, "File is to big!")
		return
	}
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		log.Printf("Error when try to get file: %v\n", err)
		//sendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
		return
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("Error when try to read file: %v\n", err)
		common.SendErrorResponse(c.Writer, http.StatusInternalServerError, "Internal Error")
		return
	}
	vid := c.Param("vid")
	if err := ioutil.WriteFile(common.VideoPath+vid, data, 0666); err != nil {
		log.Printf("Error when try to write file: %v\n", err)
		common.SendErrorResponse(c.Writer, http.StatusInternalServerError, "Internal Error")
		return
	}
	c.Writer.WriteHeader(http.StatusCreated)
	io.WriteString(c.Writer, "Uploaded successfully")

}
