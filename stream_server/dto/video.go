package dto

import (
	"github.com/gin-gonic/gin"
	"video_server/stream_server/common"
)

type VideoUploadInput struct {
	Title        string `json:"title" form:"title" comment:"视频标题" example:"标题1" validate:"required,require_title"` //视频标题
	Introduction string `json:"introduction" form:"introduction" comment:"视频简介" example:"视频简介..."`                 //视频简介
	Public       bool   `json:"public" form:"public" comment:"是否公开" example:"true"`
}

type VideoIndexOutput struct {
	Title      string `json:"title" form:"title" comment:"视频标题" example:"标题1" validate:""`              //视频标题
	VideoID    uint   `json:"video_id" form:"video_id" comment:"视频ID" example:"1" validate:""`          //视频ID
	AuthorName string `json:"author_name" form:"author_name" comment:"视频作者名称" example:"张三" validate:""` //视频作者名称
}

func (params *VideoUploadInput) BindingValidParams(c *gin.Context) error {
	return common.DefaultGetValidParams(c, params)
}
