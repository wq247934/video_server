package dto

import (
	"github.com/gin-gonic/gin"
	"video_server/stream_server/common"
)

type VideoInput struct {
	Title        string `json:"title" form:"title" comment:"视频标题" example:"标题1" validate:"required,require_title"` //视频标题
	Introduction string `json:"introduction" form:"introduction" comment:"视频简介" example:"视频简介..."`                 //视频简介
	Public       bool   `json:"public" form:"public" comment:"是否公开" example:"true"`
}

func (params *VideoInput) BindingValidParams(c *gin.Context) error {
	return common.DefaultGetValidParams(c, params)
}
