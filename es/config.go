package es

//配置结构
type Config struct {
	Addr string //连接地址
	Auth *Auth
}

//鉴权
type Auth struct {
	UserName string //用户名
	Password string //密码
}
