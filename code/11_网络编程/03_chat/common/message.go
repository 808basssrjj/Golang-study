package message

const (
	LoginMesType = "LoginMse"
	LoginResType = "LoginRes"
)

// Message 消息结构体
type Message struct {
	Type string `json:"type"` //消息类型
	Data string `json:"data"` //消息内容
}

// LoginMse 登录消息
type LoginMse struct {
	UserId   int    `json:"user_id"`
	UserPwd  string `json:"user_pwd"`
	UserName string `json:"user_name"`
}

// LoginRes 登录返回消息
type LoginRes struct {
	Code  int    `json:"code"`  //返回状态码 500未注册 200登录成功
	Error string `json:"error"` //错误信息
}
