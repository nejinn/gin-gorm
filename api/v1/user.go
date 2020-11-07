package v1

import (
	"github.com/gin-gonic/gin"
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
