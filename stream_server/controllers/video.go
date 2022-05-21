package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/contrib/sessions"
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

}

// Upload godoc
// @Summary 上传视频
// @Description 用户上传视频
// @Tags 视频接口
// @ID /video/upload
// @Accept  json
// @Produce  json
// @Param polygon body dto.VideoInput true "body"
// @Success 200 {object} middleware.Response{data=string} "success"
// @Router /video/upload [post]
func (videoController VideoController) Upload(c *gin.Context) {
	params := &dto.VideoInput{}
	if err := params.BindingValidParams(c); err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}
	session := sessions.Default(c)
	adminInfo, ok := session.Get(common.AdminSessionInfoKey).(string)
	if !ok || adminInfo == "" {
		c.Abort()
		return
	}
	var userInfo dto.AdminSessionInfo
	err := json.Unmarshal([]byte(adminInfo), &userInfo)
	if err != nil {
		fmt.Println(err)
		return
	}
	user := dao.User{Username: userInfo.UserName}
	u := user.Find()
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
