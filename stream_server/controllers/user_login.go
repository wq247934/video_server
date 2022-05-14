package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"time"
	"video_server/stream_server/common"
	"video_server/stream_server/dao"
	"video_server/stream_server/dto"
	"video_server/stream_server/middleware"
)

type UserLoginController struct {
}

func UserLoginControllerRegister(group *gin.RouterGroup) {
	userLogin := &UserLoginController{}
	group.POST("/login", userLogin.UserLogin)
	group.GET("/logout", userLogin.UserLogout)
}

// UserLogin godoc
// @Summary 用户登录
// @Description 用户登录
// @Tags 用户接口
// @ID /user_admin/login
// @Accept  json
// @Produce  json
// @Param polygon body dto.AdminLoginInput true "body"
// @Success 200 {object} middleware.Response{data=dto.AdminLoginOutput} "success"
// @Router /admin_login/login [post]
func (adminLogin UserLoginController) UserLogin(c *gin.Context) {
	params := &dto.AdminLoginInput{}
	fmt.Println(params)
	if err := params.BindingValidParams(c); err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}
	//1. 通过params.Username 拿到用户信息adminInfo
	//2.adminInfo.salt + params.Password sha256 =>saltPassword
	//3.saltPassword == adminInfo.password
	user := &dao.User{
		Username: params.UserName,
		Password: params.Password,
	}
	u := user.LoginCheck()
	if u == nil {
		middleware.ResponseError(c, 2002, errors.New("login failed"))
		return
	}
	//设置session
	sessionInfo := &dto.AdminSessionInfo{
		ID:        u.ID,
		UserName:  u.Username,
		LoginTime: time.Now(),
	}
	sessBts, err := json.Marshal(sessionInfo)
	if err != nil {
		middleware.ResponseError(c, 2003, err)
		return
	}
	sess := sessions.Default(c)
	sess.Set(common.AdminSessionInfoKey, string(sessBts))
	fmt.Println(common.AdminSessionInfoKey, string(sessBts))
	err = sess.Save()
	if err != nil {
		middleware.ResponseError(c, 2004, err)
		return
	}
	out := &dto.AdminLoginOutput{Token: u.Username}
	middleware.ResponseSuccess(c, out)
}

// UserLogout godoc
// @Summary 管理员退出
// @Description 管理员退出
// @Tags 管理员接口
// @ID /admin_login/logout
// @Accept  json
// @Produce  json
// @Success 200 {object} middleware.Response{data=string} "success"
// @Router /admin_login/logout [get]
func (adminLogin UserLoginController) UserLogout(c *gin.Context) {
	sess := sessions.Default(c)
	sess.Delete(common.AdminSessionInfoKey)

	if err := sess.Save(); err != nil {
		middleware.ResponseError(c, 2004, err)
		return
	}
	middleware.ResponseSuccess(c, "")
}
