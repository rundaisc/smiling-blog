package tools

type ResponseCode int

const (
	OK               ResponseCode = 0
	ParamsError      ResponseCode = 1000
	UnKnowError      ResponseCode = 1001
	ValidateError    ResponseCode = 1002
	AdminLoginFailed ResponseCode = 1003
)

var CodeMsg = map[ResponseCode]string{
	OK:               "success",
	UnKnowError:      "服务器异常，请重试",
	ParamsError:      "参数类型异常",
	ValidateError:    "参数验证错误，请检查参数格式",
	AdminLoginFailed: "账号或密码错误",
}

func GetCodeAndMsg(code ResponseCode) string {
	value, ok := CodeMsg[code]
	if ok {
		return value
	}
	return ""
}
