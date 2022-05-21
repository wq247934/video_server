package dto

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
	"time"
	"video_server/stream_server/common"
)

type UserSessionInfo struct {
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
	Token    string `json:"token" form:"token" comment:"token" example:"admin" validate:""`          //token
	UserName string `json:"username" form:"username" comment:"username" example:"admin" validate:""` //用户名
}

func GetUserInfoFromSession(c *gin.Context) (*UserSessionInfo, error) {
	session := sessions.Default(c)
	adminInfo, ok := session.Get(common.AdminSessionInfoKey).(string)
	if !ok || adminInfo == "" {
		c.Abort()
		return nil, errors.New("admin session is null")
	}
	var userInfo UserSessionInfo
	err := json.Unmarshal([]byte(adminInfo), &userInfo)
	if err != nil {
		log.Println(err)
		return nil, errors.New("parse adminInfo error:" + err.Error())
	}
	return &userInfo, nil
}
