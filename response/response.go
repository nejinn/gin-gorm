package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type OkRet struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type MsgResponse struct {
	Ret OkRet `json:"ret"`
}

type DataResponse struct {
	Ret  OkRet       `json:"ret"`
	Data interface{} `json:"data"`
}

// 请求成功返回 json
func OkData(c *gin.Context, ret OkRet, data interface{}) {
	//res := response{
	//	Ret: ret,
	//	Data: map[string]interface{}{},
	//}
	if data == nil {
		c.JSON(http.StatusOK, MsgResponse{
			Ret: ret,
		})
		return
	} else {
		c.JSON(http.StatusOK, DataResponse{
			Ret:  ret,
			Data: data,
		})
		return
	}
}

// 自定义msg code 返回json
func ErrorMsg(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, MsgResponse{
		Ret: OkRet{
			Code: code,
			Msg:  msg,
		},
	})
}
