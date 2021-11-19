package mongo

import "time"

// ClientName mongodb client名称枚举
type ClientName string

// 默认超时时间配置
const (
	DefaultConnectTimeout time.Duration = 5 // mongo connect 建立mongodb连接默认超时时间 5s
	DefaultQueryTimeout   time.Duration = 5 // mongo query 查询默认超时时间 5s
)
