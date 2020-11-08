package response

var Ok = OkRet{
	Code: 200,
	Msg:  "ok!",
}

var TokenError = OkRet{
	Code: 100002,
	Msg: "token错误",
}

var RegisterError = OkRet{
	Code: 200002,
	Msg:  "用户注册错误",
}