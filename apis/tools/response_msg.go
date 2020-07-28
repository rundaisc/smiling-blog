package tools

type ResponseCode int

const (
	OK                 ResponseCode = 0
	ParamsError        ResponseCode = 1000
	UnKnowError        ResponseCode = 1001
	ValidateError      ResponseCode = 1002
	AdminLoginFailed   ResponseCode = 1003
	FlinkNotFound      ResponseCode = 1004
	FlinkSaveFailed    ResponseCode = 1005
	NavNotFound        ResponseCode = 1006
	NavSaveFailed      ResponseCode = 1007
	CategoryNotFound   ResponseCode = 1008
	CategorySaveFailed ResponseCode = 1009
)

var CodeMsg = map[ResponseCode]string{
	OK:                 "success",
	UnKnowError:        "服务器异常，请重试",
	ParamsError:        "参数类型异常",
	ValidateError:      "参数验证错误，请检查参数格式",
	AdminLoginFailed:   "账号或密码错误",
	FlinkNotFound:      "未找到匹配的友情链接",
	FlinkSaveFailed:    "保存友情链接失败",
	NavNotFound:        "未找到匹配的导航",
	NavSaveFailed:      "保存导航失败",
	CategoryNotFound:   "未找到匹配的分类",
	CategorySaveFailed: "保存分类失败",
}

func GetCodeAndMsg(code ResponseCode) string {
	value, ok := CodeMsg[code]
	if ok {
		return value
	}
	return ""
}
