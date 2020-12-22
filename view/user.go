package view

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"note/global"
	"note/middleware"
	"note/model"
	"note/utils"
	"strconv"
	"time"
)

type RegisterStruct struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginStruct struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Fields struct {
	Key          string `json:"key"`
	Label        string `json:"label"`
	Class        string `json:"class"`
	StickyColumn bool   `json:"stickyColumn"`
	Fixed        string `json:"fixed"`
	Variant      string `json:"variant"`
	TdClass      string `json:"tdClass"`
}

type UserListInfo struct {
	ID                int    `json:"id"`
	Username          string `json:"username"`
	UserPic           string `json:"user_pic"`
	UserType          int    `json:"user_type"`
	UserPhone         string `json:"user_phone"`
	UserEmail         string `json:"user_email"`
	UserExp           int    `json:"user_exp"`
	UserLevel         int    `json:"user_level"`
	UserBirthday      string `json:"user_birthday"`
	UserGender        int    `json:"user_gender"` // 1 男 2 女
	UserIp            string `json:"user_ip"`
	UserLastLoginIp   string `json:"user_last_login_ip"`
	UserLastLoginTime string `json:"user_last_login_time"`
	CreatedAt         string `json:"create_date"`
	UpdatedAt         string `json:"update_date"`
	IsDelete          *bool  `json:"is_delete"`
}

type UserListRes struct {
	Items  []UserListInfo `json:"items"`
	Fields []Fields       `json:"fields"`
	Page   utils.Page     `json:"page"`
}

type CheckUsername struct {
	Username string `json:"username"`
}

type CheckUsernameRes struct {
	CheckResult bool   `json:"check_result"`
	Msg         string `json:"msg"`
}

type AddUserInfo struct {
	Username     string `json:"username" binding:"required"`
	Password     string `json:"password" binding:"required"`
	UserType     int    `json:"user_type" binding:"required"`
	UserPhone    string `json:"user_phone" binding:"required"`
	UserEmail    string `json:"user_email" binding:"required,email"`
	UserBirthday string `json:"user_birthday" binding:"required"`
	UserGender   int    `json:"user_gender" binding:"required"` // 1 男 2 女
}

type DeleteUserInfo struct {
	ID int `json:"user_id"`
}

type DeleteUserListInfo struct {
	ID []int `json:"user_id"`
}

//用户注册
func Register(r *model.User) (err error) {
	var registerUser model.User
	global.NLY_DB.Where("username = ?", r.Username).First(&registerUser)
	if !errors.Is(global.NLY_DB.Where("username = ?", r.Username).First(&registerUser).Error, gorm.ErrRecordNotFound) {
		return errors.New("用户名已经注册")
	}
	password := utils.NewMD5([]byte(r.Password))
	avatar, avatarErr := utils.CreateDefaultAvatar(r.Username)
	if avatarErr != nil {
		return errors.New("生成头像失败")
	}
	user := model.User{
		Username: r.Username,
		Password: password,
		UserPic:  avatar,
	}
	err = global.NLY_DB.Create(&user).Error
	return err
}

// 登录校验账号密码 （admin）
func Login(r *model.User) (err error, userInfo *model.User) {
	var loginUser model.User
	if r.Username == "" {
		return errors.New("用户名不能为空"), nil
	}
	if r.Password == "" {
		return errors.New("密码不能为空"), nil
	}
	r.Password = utils.NewMD5([]byte(r.Password))

	global.NLY_DB.Where("username = ? AND password = ?", r.Username, r.Password).First(&loginUser)
	if errors.Is(global.NLY_DB.Where("username = ? AND password = ?", r.Username, r.Password).First(&loginUser).Error, gorm.ErrRecordNotFound) {
		return errors.New("登录信息错误"), nil
	}
	if *loginUser.IsDelete == true {
		return errors.New("用户已停用"), nil
	}
	if loginUser.UserType != 1 {
		return errors.New("非超级用户不能登录"), nil
	}
	nowTime := time.Now()
	loginUser.UserLastLoginTime = &nowTime
	global.NLY_DB.Save(&loginUser)
	return nil, &loginUser
}

// 用户列表
func UserList(c *gin.Context) (err error, res interface{}) {
	//offset := pageInfo.Size * (pageInfo.Page - 1)
	var limit int
	var offset int
	var page int
	var size int
	if querySize, isExist := c.GetQuery("size"); isExist == true {
		size, _ = strconv.Atoi(querySize)
	} else {
		size = utils.DefaultSize
	}

	if queryPage, isExist := c.GetQuery("page"); isExist == true {
		page, _ = strconv.Atoi(queryPage)
	} else {
		page = utils.DefaultPage
	}
	limit = size
	offset = limit * (page - 1)

	var items []model.User
	var count int
	var data UserListRes
	var userDb = global.NLY_DB
	if username, isExist := c.GetQuery("username"); isExist == true {
		userDb = userDb.Where("username LIKE ?", fmt.Sprintf("%s%s%s", "%", username, "%"))
	}
	if userPhone, isExist := c.GetQuery("user_phone"); isExist == true {
		userDb = userDb.Where("user_phone LIKE ?", fmt.Sprintf("%s%s%s", "%", userPhone, "%"))
	}
	if userEmail, isExist := c.GetQuery("user_email"); isExist == true {
		userDb = userDb.Where("user_email LIKE ?", fmt.Sprintf("%s%s%s", "%", userEmail, "%"))
	}
	if userType, isExist := c.GetQuery("user_type"); isExist == true {
		if queryUserType, err := strconv.ParseInt(userType, 10, 64); err == nil {
			userDb = userDb.Where("user_type = ?", queryUserType)
		}
	}
	if isDelete, isExist := c.GetQuery("is_delete"); isExist == true {
		if queryIsDelete, err := strconv.ParseBool(isDelete); err == nil {
			userDb = userDb.Where("is_delete = ?", queryIsDelete)
		}
	}
	err = userDb.Find(&items).Count(&count).Error
	if err != nil {
		return err, nil
	}
	err = userDb.Limit(limit).Offset(offset).Order("id").Find(&items).Error
	data.Page.Page = page
	data.Page.Count = count
	data.Page.Size = size
	data.Fields = GenFields(data.Fields)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			data.Items = []UserListInfo{}
			return nil, data
		}
		return err, nil
	}
	for _, item := range items {
		var itemAppend UserListInfo
		itemAppend.ID = item.ID
		itemAppend.Username = item.Username
		itemAppend.CreatedAt = item.CreatedAt.Format("2006-01-02 15:04:05")
		itemAppend.UpdatedAt = item.UpdatedAt.Format("2006-01-02 15:04:05")
		itemAppend.UserPic = fmt.Sprintf("http://%s/%s", c.Request.Host, item.UserPic)
		itemAppend.IsDelete = item.IsDelete
		itemAppend.UserBirthday = item.UserBirthday
		itemAppend.UserEmail = item.UserEmail
		itemAppend.UserExp = item.UserExp
		itemAppend.UserGender = item.UserGender
		itemAppend.UserIp = item.UserIp
		itemAppend.UserLastLoginIp = item.UserLastLoginIp
		itemAppend.UserLastLoginTime = ""
		if item.UserLastLoginTime != nil {
			itemAppend.UserLastLoginTime = item.UserLastLoginTime.Format("2006-01-02 15:04:05")
		}
		itemAppend.UserLevel = item.UserLevel
		itemAppend.UserPhone = item.UserPhone
		itemAppend.UserType = item.UserType
		data.Items = append(data.Items, itemAppend)
	}
	return nil, data
}

func GenFields(f []Fields) (res []Fields) {
	f = append(f, Fields{
		Key:          "user_pic",
		Label:        "头像",
		Class:        "text-center",
		StickyColumn: true,
		Fixed:        "left",
		TdClass:      "custom-tbody-td-background-color-info",
	})
	f = append(f, Fields{
		Key:          "username",
		Label:        "用户名称",
		Class:        "text-center",
		StickyColumn: true,
		Fixed:        "left",
		TdClass:      "custom-tbody-td-background-color-info",
	})
	f = append(f, Fields{
		Key:          "id",
		Label:        "用户id",
		Class:        "text-center",
		StickyColumn: true,
		Fixed:        "left",
		TdClass:      "custom-tbody-td-background-color-info",
	})
	f = append(f, Fields{
		Key:          "isSelected",
		Label:        "选中",
		Class:        "text-center text-info",
		StickyColumn: true,
		Fixed:        "left",
		TdClass:      "custom-tbody-td-background-color-info",
	})

	f = append(f, Fields{
		Key:   "user_type",
		Label: "类型",
		Class: "text-center",
	})
	f = append(f, Fields{
		Key:   "user_phone",
		Label: "电话",
		Class: "text-center",
	})
	f = append(f, Fields{
		Key:   "user_email",
		Label: "邮箱",
		Class: "text-center",
	})
	f = append(f, Fields{
		Key:   "user_exp",
		Label: "经验",
		Class: "text-center",
	})
	f = append(f, Fields{
		Key:   "user_level",
		Label: "等级",
		Class: "text-center",
	})
	f = append(f, Fields{
		Key:   "user_last_login_time",
		Label: "最后登录时间",
		Class: "text-center",
	})
	f = append(f, Fields{
		Key:   "user_birthday",
		Label: "生日",
		Class: "text-center",
	})
	f = append(f, Fields{
		Key:   "user_gender",
		Label: "性别",
		Class: "text-center",
	})
	f = append(f, Fields{
		Key:   "user_ip",
		Label: "最后登录ip",
		Class: "text-center",
	})
	f = append(f, Fields{
		Key:   "create_date",
		Label: "注册日期",
		Class: "text-center",
	})
	//f = append(f, Fields{
	//	Key:   "update_date",
	//	Label: "更新日期",
	//	Class: "text-center",
	//})
	f = append(f, Fields{
		Key:   "is_delete",
		Label: "已停用",
		Class: "text-center",
	})
	//f = append(f, Fields{
	//	Key:   "create_user",
	//	Label: "创建人",
	//	Class: "text-center",
	//})
	//f = append(f, Fields{
	//	Key:   "update_user",
	//	Label: "更新人",
	//	Class: "text-center",
	//})
	f = append(f, Fields{
		Key:          "action",
		Label:        "操作",
		Class:        "text-center",
		StickyColumn: true,
		Fixed:        "right",
		TdClass:      "custom-tbody-td-background-color-info",
	})

	return f
}

// 检验用户是否存在
func CheckUsernameExist(u *CheckUsername) (res CheckUsernameRes) {
	if !errors.Is(global.NLY_DB.Where("username = ?", u.Username).Find(&model.User{}).Error, gorm.ErrRecordNotFound) {
		return CheckUsernameRes{
			CheckResult: true,
			Msg:         "用户名已存在",
		}
	} else {
		return CheckUsernameRes{
			CheckResult: false,
		}
	}
}

// 新增用户
func AddUser(a *AddUserInfo, c *gin.Context) (res error) {
	if len(a.Password) < 6 {
		return errors.New("密码不能小于六位")
	}

	var username CheckUsername
	username.Username = a.Username
	isExist := CheckUsernameExist(&username)
	if isExist.CheckResult {
		return errors.New(isExist.Msg)
	}
	avatar, avatarErr := utils.CreateDefaultAvatar(a.Username)
	if avatarErr != nil {
		return errors.New("生成头像失败")
	}
	user := middleware.GetJwtUser(c)
	var NewUser model.User
	NewUser.Password = utils.NewMD5([]byte(a.Password))
	NewUser.Username = a.Username
	NewUser.UserBirthday = a.UserBirthday
	NewUser.UserEmail = a.UserEmail
	NewUser.UserPhone = a.UserPhone
	NewUser.UserGender = a.UserGender
	NewUser.UserType = a.UserType
	NewUser.UserPic = avatar
	NewUser.CreateId = user.ID
	err := global.NLY_DB.Create(&NewUser).Error
	return err

}

// 停用用户
func (s *DeleteUserInfo) DeleteUser() (err error) {
	var DeleteUser model.User
	err = global.NLY_DB.Model(&DeleteUser).Where("id = ?", s.ID).Update("is_delete", true).Error
	if err != nil {
		return errors.New("停用用户出错")
	}
	return nil
}

// 批量停用用户
func (s *DeleteUserListInfo) DeleteUserList() (err error) {
	var DeleteUser model.User
	deleteWork := global.NLY_DB.Begin()
	for _, item := range s.ID {
		err = deleteWork.Model(&DeleteUser).Where("id = ?", item).Update("is_delete", true).Error
		if err != nil {
			deleteWork.Rollback()
			return errors.New("停用用户出错")
		}
	}
	deleteWork.Commit()
	return nil
}

// 启用用户
func (s *DeleteUserInfo) RecoverUser() (err error) {
	var DeleteUser model.User
	err = global.NLY_DB.Model(&DeleteUser).Where("id = ?", s.ID).Update("is_delete", false).Error
	if err != nil {
		return errors.New("起用用户出错")
	}
	return nil
}
