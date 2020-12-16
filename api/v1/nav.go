package v1

import (
	"github.com/gin-gonic/gin"
	"note/global"
	"note/middleware"
	"note/model"
	"note/response"
	"note/view"
)

// 创建导航
func CreateNav(c *gin.Context) {
	l := model.UserNav{}
	user := middleware.GetJwtUser(c)
	l.CreateId = user.ID
	if err := c.ShouldBindJSON(&l); err != nil {
		response.ErrorMsgWithResponseMsg(c, response.CreateNavError)
		return
	}
	if err := view.CreateNav(&l); err != nil {
		response.ErrorWithCustomMsg(c, response.CreateNavError.Code, err.Error())
	} else {
		response.OkData(c, response.Ok, global.NLY_NIL_RES)
	}
}

// 获取导航栏

func GetNavList(c *gin.Context) {
	var navList []model.UserNav
	res, err := view.GetNavList(navList)

	if err != nil {
		response.ErrorMsgWithResponseMsg(c, response.GetNavListError)
		return
	}
	response.OkData(c, response.Ok, res)
}
