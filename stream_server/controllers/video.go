package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
	"video_server/stream_server/common"
	"video_server/stream_server/dao"
	"video_server/stream_server/dto"
	"video_server/stream_server/middleware"
)

type VideoController struct {
}

func VideoControllerRegister(group *gin.RouterGroup) {
	video := &VideoController{}
	group.POST("/upload", video.Upload)
	group.GET("/videos", video.VideoIndex)
}

// Upload godoc
// @Summary 上传视频
// @Description 用户上传视频
// @Tags 视频接口
// @ID /video/upload
// @Accept  json
// @Produce  json
// @Param polygon body dto.VideoUploadInput true "body"
// @Success 200 {object} middleware.Response{data=string} "success"
// @Router /video/upload [post]
func (videoController VideoController) Upload(c *gin.Context) {
	params := &dto.VideoUploadInput{}
	if err := params.BindingValidParams(c); err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}
	userInfo, err := dto.GetUserInfoFromSession(c)
	if err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}
	user := dao.User{Username: userInfo.UserName}
	u := user.Find()
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, common.MaxUploadSize)
	if err := c.Request.ParseMultipartForm(common.MaxUploadSize); err != nil {
		log.Println("File is to big!")
		log.Println(err)
		middleware.ResponseError(c, 2002, err)
		return
	}
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		log.Printf("Error when try to get file: %v\n", err)
		common.SendErrorResponse(c.Writer, http.StatusInternalServerError, "Internal Error")
		return
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("Error when try to read file: %v\n", err)
		common.SendErrorResponse(c.Writer, http.StatusInternalServerError, "Internal Error")
		return
	}

	fileHash := common.GetFileHash(data)
	hashStr := fmt.Sprintf("%+x", fileHash)
	hash := dao.Hash{HashSum: hashStr}
	h := hash.Find()
	if h.ID == 0 {
		_uuid, _ := uuid.NewV4()
		filePath := common.VideoPath + _uuid.String()
		if err := ioutil.WriteFile(filePath, data, 0666); err != nil {
			log.Printf("Error when try to write file: %v\n", err)
			common.SendErrorResponse(c.Writer, http.StatusInternalServerError, "Internal Error")
			return
		}
		h = &dao.Hash{
			Model:   gorm.Model{},
			Path:    filePath,
			HashSum: hashStr,
			Active:  false,
		}
	}
	video := dao.Video{
		Model:        gorm.Model{},
		Title:        params.Title,
		HashID:       h.ID,
		Hash:         *h,
		User:         *u,
		Introduction: params.Introduction,
		Public:       params.Public,
	}

	if err := video.Save(); err != nil {
		log.Println(err)
		return
	}

	c.Writer.WriteHeader(http.StatusOK)
	io.WriteString(c.Writer, "Uploaded successfully")

}

// VideoIndex godoc
// @Summary 视频列表
// @Description 视频列表
// @Tags 视频接口
// @ID /video/videos
// @Accept  json
// @Produce  json
// @Success 200 {object} middleware.Response{data=[]dto.VideoIndexOutput} "success"
// @Router /video/videos [get]
func (videoController VideoController) VideoIndex(c *gin.Context) {
	userInfo, err := dto.GetUserInfoFromSession(c)
	if err != nil {
		middleware.ResponseError(c, 2003, err)
	}
	user := dao.User{Username: userInfo.UserName}
	user = *user.Find()
	video := dao.Video{AuthorID: user.ID}
	videos := video.Find()
	var videoOutputs []dto.VideoIndexOutput
	for _, video := range *videos {
		videoOutput := dto.VideoIndexOutput{
			Title:      video.Title,
			VideoID:    video.ID,
			AuthorName: video.User.Username,
		}
		videoOutputs = append(videoOutputs, videoOutput)
	}
	middleware.ResponseSuccess(c, videoOutputs)
}

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
