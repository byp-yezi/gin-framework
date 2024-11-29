package e

// 全局
const (
	SUCCESS       int = 200
	InvalidParams int = 400
	Unauthorized  int = 401
	ERROR         int = 500
)

// 用户
const (
	ErrorExistUser int = 10001 + iota
	ErrorNotExistUser
)

var CodeMsg = map[int]string{
	SUCCESS:       "请求成功",
	InvalidParams: "请求参数错误",
	Unauthorized:  "未认证",
	ERROR:         "请求失败",

	ErrorExistUser:    "用户已存在",
	ErrorNotExistUser: "用户不存在",
}

func GetMsg(code int) string {
	if msg, ok := CodeMsg[code]; ok {
		return msg
	}
	return CodeMsg[ERROR]
}
