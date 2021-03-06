package response

var Ok = OkRet{
	Code: 200,
	Msg:  "ok!",
}

var TokenError = OkRet{
	Code: 100002,
	Msg:  "token不能为空",
}

var RegisterError = OkRet{
	Code: 200002,
	Msg:  "用户注册错误",
}

var CreateAvatarError = OkRet{
	Code: 200003,
	Msg:  "创建默认头像失败",
}

var LoginError = OkRet{
	Code: 200004,
	Msg:  "登录信息错误",
}

var CreateNavError = OkRet{
	Code: 200005,
	Msg:  "创建导航错误",
}

var GetNavListError = OkRet{
	Code: 200006,
	Msg:  "获取导航列表错误",
}

var GetUserListError = OkRet{
	Code: 200007,
	Msg:  "获取用户列表错误",
}

var AddUserError = OkRet{
	Code: 200008,
	Msg:  "新增用户错误",
}

var DeleteUserError = OkRet{
	Code: 200009,
	Msg:  "删除用户错误",
}

var RecoverUserError = OkRet{
	Code: 200010,
	Msg:  "启动用户错误",
}

var InfoBoxError = OkRet{
	Code: 200011,
	Msg:  "获取infox信息错误",
}

//var JwtError = OkRet{
//	Code: 200006,
//	Msg:  "token不能为空",
//}
