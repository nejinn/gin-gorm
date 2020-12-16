package v1

import (
	"github.com/gin-gonic/gin"
	"note/global"
	"note/middleware"
	"note/model"
	"note/response"
	"note/utils"
	"note/view"
)

// 创建导航
func CreateNav(c *gin.Context) {
	l := model.Nav{}
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
	var navList []model.Nav
	res, err := view.GetNavList(navList)

	if err != nil {
		response.ErrorMsgWithResponseMsg(c, response.GetNavListError)
		return
	}
	var navListTree []*utils.NavListJson
	for _, item := range res {
		navListItem := utils.NavListJson{
			Id:       item.ID,
			Name:     item.Name,
			Icon:     item.Icon,
			ParentId: item.ParentId,
			Exact:    true,
		}
		if item.Type == 1 {
			navListItem.NavType = "nly-sidebar-nav-item"
			navListItem.Router = item.Router
		} else if item.Type == 2 {
			navListItem.NavType = "nly-sidebar-nav-tree"
		} else {
			navListItem.NavType = "nly-sidebar-nav-header"
		}
		navListTree = append(navListTree, &navListItem)
	}

	result := utils.NavListTree(navListTree)

	response.OkData(c, response.Ok, result)
}
