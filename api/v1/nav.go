package v1

import (
	"github.com/gin-gonic/gin"
	"note/global"
	"note/response"
	"note/view"
)

// 创建导航
func CreateUserNav(c *gin.Context) {
	l := view.UserNav{}
	if err := c.ShouldBindJSON(&l); err != nil {
		response.ErrorMsgWithResponseMsg(c, response.CreateNavError)
		return
	}
	if err := l.CreateUserNav(); err != nil {
		response.ErrorWithCustomMsg(c, response.CreateNavError.Code, err.Error())
	} else {
		response.OkData(c, response.Ok, global.NLY_NIL_RES)
	}
}
