package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"note/global"
	"note/middleware"
	"note/model"
	"note/utils"
	"note/view"
	"time"
	//"note/model"
	"note/response"
)

func GetUserList(c *gin.Context) {
	var user []model.User
	global.NLY_DB.Find(&user)
	response.OkData(c, response.Ok, user)
}

func LoginAdmin(c *gin.Context) {
	l := view.LoginStruct{}
	if err := c.ShouldBindJSON(&l); err != nil {
		response.ErrorMsgWithResponseMsg(c, response.LoginError)
	}
	user := &model.User{Username: l.Username, Password: l.Password}
	if err, userInfo := view.Login(user); err != nil {
		response.ErrorWithCustomMsg(c, response.LoginError.Code, err.Error())
	} else {
		middleware.DispenseToken(c, userInfo)
	}
}

// 注册用户
func Register(c *gin.Context) {
	var r model.User
	_ = c.ShouldBindJSON(&r)

	err := view.Register(&r)
	if err != nil {
		response.ErrorWithCustomMsg(c, response.RegisterError.Code, err.Error())
	} else {
		response.OkData(c, response.Ok, global.NLY_NIL_RES)
	}
}

// 上传图片测试
func UploadImg(c *gin.Context) {
	f, err := c.FormFile("img")

	fileName := utils.NewMD5([]byte(fmt.Sprintf("%s%s", time.Now().String())))
	fildDir := fmt.Sprintf("%s/%d/%s/", global.NLY_CONFIG.Static.Path, time.Now().Year(), time.Now().Month().String())
	err = utils.CreateDir(fildDir)
	if err != nil {
		response.ErrorWithCustomMsg(c, 200003, err.Error())
	}
	filepath := fmt.Sprintf("%s%s%s", fildDir, fileName, ".png")

	c.SaveUploadedFile(f, filepath)

	type Data struct {
		Path string
	}
	data := Data{
		Path: fmt.Sprintf("%s/%s", c.Request.Host, filepath),
	}
	response.OkData(c, response.Ok, data)
}
