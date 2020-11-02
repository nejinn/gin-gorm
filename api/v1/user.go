package v1

import (
	"github.com/gin-gonic/gin"
	//"note/model"
	"note/response"
)

func GetUserList(c *gin.Context) {
	type User struct {
		User string
	}
	user := User{
		User: "ssss",
	}
	response.OkData(c, response.Ok, user)

}
