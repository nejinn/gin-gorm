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
	l := utils.PageInfo{}
	_ = c.ShouldBindJSON(&l)
	err, data := view.UserList(&l, c)
	if err != nil {
		response.ErrorWithCustomMsg(c, response.GetUserListError.Code, err.Error())
		return
	}
	response.OkData(c, response.Ok, data)
}

func LoginAdmin(c *gin.Context) {
	l := view.LoginStruct{}
	if err := c.ShouldBindJSON(&l); err != nil {
		response.ErrorMsgWithResponseMsg(c, response.LoginError)
		return
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

// 用户信息
func GetUserInfo(c *gin.Context) {
	var navList []model.Nav
	res, err := view.GetNavList(navList)
	user := middleware.GetJwtUser(c)
	user.UserPic = fmt.Sprintf("%s%s/%s", "http://", c.Request.Host, user.UserPic)
	if err != nil {
		response.ErrorMsgWithResponseMsg(c, response.GetNavListError)
		return
	}
	var navListTree []*utils.NavListJson
	for _, item := range res {
		navListItem := utils.NavListJson{
			Id:       item.ID,
			Icon:     item.Icon,
			ParentId: item.ParentId,
			Exact:    true,
		}
		if item.Type == 1 {
			navListItem.NavType = "nly-sidebar-nav-item"
			navListItem.Router = item.Router
			navListItem.Name = item.Name
		} else if item.Type == 2 {
			navListItem.NavType = "nly-sidebar-nav-tree"
			navListItem.Target = fmt.Sprintf("%d", navListItem.Id)
			navListItem.Text = item.Name
		} else {
			navListItem.NavType = "nly-sidebar-nav-header"
			navListItem.Text = item.Name
		}
		navListTree = append(navListTree, &navListItem)
	}

	nav := utils.NavListTree(navListTree)

	type result struct {
		User    model.User           `json:"user_info"`
		UserNav []*utils.NavListJson `json:"user_sidebar"`
	}
	userInfo := result{
		User:    user,
		UserNav: nav,
	}
	response.OkData(c, response.Ok, userInfo)

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

// 检验用户名是否存在
func CheckUsernameExist(c *gin.Context) {
	u := view.CheckUsername{}
	_ = c.ShouldBindJSON(&u)
	res := view.CheckUsernameExist(&u)
	response.OkData(c, response.Ok, res)
}

// 新增用户
func AddUser(c *gin.Context) {
	a := view.AddUserInfo{}
	if err := c.ShouldBindJSON(&a); err != nil {
		response.ErrorMsgWithResponseMsg(c, response.AddUserError)
		return
	}
	_ = view.AddUser(&a)

}
