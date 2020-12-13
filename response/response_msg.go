package response

var Ok = OkRet{
	Code: 200,
	Msg:  "ok!",
}

var TokenError = OkRet{
	Code: 100002,
	Msg:  "token错误",
}

var RegisterError = OkRet{
	Code: 200002,
	Msg:  "用户注册错误",
}

var LoginError = OkRet{
	Code: 200004,
	Msg:  "登录信息错误",
}

var CreateAvatarError = OkRet{
	Code: 200003,
	Msg:  "创建默认头像失败",
}
