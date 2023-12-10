package errmessage

const (
	SUCCESS = 200
	ERROR   = 500

	//code=1000..表示用户模块的错误
	ERROR_USERNAME_USED    = 1001
	ERROR_PASSWORD_WRONG   = 1002
	ERROR_USER_NOT_EXIST   = 1003
	ERROR_TOKEN_NOT_EXIST  = 1004
	ERROR_TOKEN_RUNTIME    = 1005
	ERROR_TOKEN_WRONG      = 1006
	ERROR_TOKEN_TYPE_WRONG = 1007
	ERROR_PHONE_USED       = 1008
)

var codemsg = map[int]string{
	SUCCESS:                "OK",
	ERROR:                  "FAIL",
	ERROR_USERNAME_USED:    "用户名已被使用",
	ERROR_PASSWORD_WRONG:   "用户密码错误",
	ERROR_USER_NOT_EXIST:   "用户不存在",
	ERROR_TOKEN_NOT_EXIST:  "Token不存在",
	ERROR_TOKEN_RUNTIME:    "Token已过期",
	ERROR_TOKEN_WRONG:      "Token不正确",
	ERROR_TOKEN_TYPE_WRONG: "Token格式错误",
	ERROR_PHONE_USED:       "手机号已被使用",
}

func Geterrmessage(code int) string {
	return codemsg[code]
}
