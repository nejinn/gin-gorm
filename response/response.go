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
func ErrorWithCustomMsg(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, MsgResponse{
		Ret: OkRet{
			Code: code,
			Msg:  msg,
		},
	})
}

// 返回 response_msg 中定义好的常量错误码
func ErrorMsgWithResponseMsg(c *gin.Context, errJson OkRet) {
	c.JSON(http.StatusOK, MsgResponse{
		Ret: OkRet{
			Code: errJson.Code,
			Msg:  errJson.Msg,
		},
	})
}
