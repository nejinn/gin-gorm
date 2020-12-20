package view

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"note/global"
	"note/model"
	"note/utils"
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

type Feilds struct {
	Key          string `json:"key"`
	Label        string `json:"label"`
	Class        string `json:"class"`
	StickyColumn bool   `json:"stickyColumn"`
	Fixed        string `json:"fixed"`
	Variant      string `json:"variant"`
	TdClass      string `json:"tdClass"`
}

type UserListInfo struct {
	ID                int       `json:"id"`
	Username          string    `json:"username"`
	UserPic           string    `json:"user_pic"`
	UserType          bool      `json:"user_type"`
	UserPhone         string    `json:"user_phone"`
	UserEmail         string    `json:"user_email"`
	UserExp           int       `json:"user_exp"`
	UserLevel         int       `json:"user_level"`
	UserBirthday      string    `json:"user_birthday"`
	UserGender        int       `json:"user_gender"` // 1 男 2 女
	UserIp            string    `json:"user_ip" gorm:"column:user_ip;default:null"`
	UserLastLoginIp   string    `json:"user_last_login_ip" gorm:"column:user_last_login_ip;default:null"`
	UserLastLoginTime time.Time `json:"user_last_login_time" gorm:"column:user_last_login_time;default:null"`
	CreatedAt         time.Time `json:"create_date" gorm:"column:create_date"`
	UpdatedAt         time.Time `json:"update_date" gorm:"column:update_date"`
	DeletedTime       time.Time `json:"deleted_date" gorm:"column:deleted_date;default:null"`
	IsDelete          bool      `json:"is_delete" gorm:"column:is_delete;default:false"`
	CreateId          int       `json:"create_id" gorm:"column:create_id;default:null"`
	UpdateId          int       `json:"update_id" gorm:"column:update_id;default:null"`
}

type UserListRes struct {
	Items  []model.User `json:"items"`
	Fields []Feilds     `json:"fields"`
	Page   utils.Page   `json:"page"`
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
	return nil, &loginUser
}

// 用户列表
func UserList(pageInfo *utils.PageInfo, c *gin.Context) (err error, res interface{}) {
	offset := pageInfo.Size * (pageInfo.Page - 1)
	limit := pageInfo.Size
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
	if userType, isExist := c.GetQuery("user_email"); isExist == true {
		userDb = userDb.Where("user_type LIKE ?", fmt.Sprintf("%s%s%s", "%", userType, "%"))
	}
	if isDelete, isExist := c.GetQuery("user_email"); isExist == true {
		userDb = userDb.Where("is_delete LIKE ?", fmt.Sprintf("%s%s%s", "%", isDelete, "%"))
	}
	err = userDb.Find(&model.User{}).Count(&count).Error
	if err != nil {
		return err, nil
	}
	if limit != 0 {
		err = userDb.Limit(limit).Offset(offset).Find(&items).Error
	} else {
		err = userDb.Limit(utils.DefaultSize).Offset(offset).Find(&items).Error
	}
	if err != nil {
		return err, nil
	}
	data.Items = items
	for index, item := range data.Items {
		data.Items[index].UserPic = fmt.Sprintf("http://%s/%s", c.Request.Host, item.UserPic)
		//update_time_str,_ :=fmt.Println(item.UpdatedAt.String())
		//updateTimeStr := item.UpdatedAt.Format("2006-01-02 15:04:05")
		tt,_:=time.ParseInLocation("2006-01-02 15:04:05", item.UpdatedAt.Format("2006-01-02 15:04:05"), time.Local)
		fmt.Println(tt.Format("2006-01-02 15:04:05"))
		data.Items[index].UpdatedAt= tt
	}
	data.Page.Page = pageInfo.Page
	data.Page.Count = count
	data.Page.Size = pageInfo.Size
	data.Fields = GenFields(data.Fields)
	return nil, data
}

func GenFields(f []Feilds) (res []Feilds) {
	f = append(f, Feilds{
		Key:          "user_pic",
		Label:        "头像",
		Class:        "text-center",
		StickyColumn: true,
		Fixed:        "left",
		TdClass:      "custom-tbody-td-background-color-info",
	})
	f = append(f, Feilds{
		Key:          "username",
		Label:        "用户名称",
		Class:        "text-center",
		StickyColumn: true,
		Fixed:        "left",
		TdClass:      "custom-tbody-td-background-color-info",
	})
	f = append(f, Feilds{
		Key:          "id",
		Label:        "用户id",
		Class:        "text-center",
		StickyColumn: true,
		Fixed:        "left",
		TdClass:      "custom-tbody-td-background-color-info",
	})
	f = append(f, Feilds{
		Key:          "isSelected",
		Label:        "选中",
		Class:        "text-center text-info",
		StickyColumn: true,
		Fixed:        "left",
		TdClass:      "custom-tbody-td-background-color-info",
	})

	f = append(f, Feilds{
		Key:   "user_type",
		Label: "类型",
		Class: "text-center",
	})
	f = append(f, Feilds{
		Key:   "user_phone",
		Label: "电话",
		Class: "text-center",
	})
	f = append(f, Feilds{
		Key:   "user_email",
		Label: "邮箱",
		Class: "text-center",
	})
	f = append(f, Feilds{
		Key:   "user_exp",
		Label: "经验",
		Class: "text-center",
	})
	f = append(f, Feilds{
		Key:   "user_level",
		Label: "等级",
		Class: "text-center",
	})
	f = append(f, Feilds{
		Key:   "user_last_login_time",
		Label: "最后登录时间",
		Class: "text-center",
	})
	f = append(f, Feilds{
		Key:   "user_birthday",
		Label: "生日",
		Class: "text-center",
	})
	f = append(f, Feilds{
		Key:   "user_gender",
		Label: "性别",
		Class: "text-center",
	})
	f = append(f, Feilds{
		Key:   "user_ip",
		Label: "最后登录ip",
		Class: "text-center",
	})
	f = append(f, Feilds{
		Key:   "create_date",
		Label: "注册日期",
		Class: "text-center",
	})
	//f = append(f, Feilds{
	//	Key:   "update_date",
	//	Label: "更新日期",
	//	Class: "text-center",
	//})
	f = append(f, Feilds{
		Key:   "is_delete",
		Label: "已删除",
		Class: "text-center",
	})
	//f = append(f, Feilds{
	//	Key:   "create_user",
	//	Label: "创建人",
	//	Class: "text-center",
	//})
	//f = append(f, Feilds{
	//	Key:   "update_user",
	//	Label: "更新人",
	//	Class: "text-center",
	//})
	f = append(f, Feilds{
		Key:          "action",
		Label:        "操作",
		Class:        "text-center",
		StickyColumn: true,
		Fixed:        "right",
		TdClass:      "custom-tbody-td-background-color-info",
	})

	return f
}
