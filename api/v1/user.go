package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"note/global"
	"note/utils"
	"note/view"
	"time"

	//"note/model"
	"note/response"
)

// @Tags SysUser
// @Summary 分页获取用户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取用户列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /user/get_user_list [get]
func GetUserList(c *gin.Context) {
	type User struct {
		User string
	}
	user := User{
		User: "ssss",
	}
	response.OkData(c, response.Ok, user)
}

func LoginAdmin(c *gin.Context) {
	type L struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	l := L{}
	_ = c.ShouldBindJSON(&l)

}

// 注册用户
func Register(c *gin.Context) {
	r := view.RegisterStruct{}
	_ = c.ShouldBindJSON(&r)

	err := view.Register(&r)
	if err != nil {
		response.ErrorMsg(c, response.RegisterError.Code, err.Error())
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
		response.ErrorMsg(c, 200003, err.Error())
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
