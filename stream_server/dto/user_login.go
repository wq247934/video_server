package dto

import (
	"github.com/gin-gonic/gin"
	"time"
	"video_server/stream_server/common"
)

type AdminSessionInfo struct {
	ID        uint      `json:"id"`
	UserName  string    `json:"username"`
	LoginTime time.Time `json:"login_time"`
}

type UserLoginInput struct {
	UserName string `json:"username" form:"username" comment:"用户名" example:"admin" validate:"required,require_username"` //用户名
	Password string `json:"password" form:"password" comment:"密码" example:"123456" validate:"required"`                  //密码
}

func (params *UserLoginInput) BindingValidParams(c *gin.Context) error {
	return common.DefaultGetValidParams(c, params)
}

type UserLoginOutput struct {
	Token string `json:"token" form:"token" comment:"token" example:"admin" validate:""` //管理员用户名
}
