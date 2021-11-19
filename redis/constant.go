package redis

// ClientName redis连接名称
type ClientName string

// ClientName redis连接名称
const (
	MKSession      ClientName = "session"
	MKCache        ClientName = "cache"
	MKCacheService ClientName = "service"
)

// Prefix 缓存前缀配置（一级命名空间管理，放在底层package中，方式各自维护各自的导致命名空间冲突）
const (
	SessionPrefix              = "wpmanage."    // SessionPrefix 工作手机团队session前缀（所有团队要统一）
	LRUCachePrefix             = "lru."         // LRUCachePrefix lru缓存前缀
	FIFOCachePrefix            = "fifo."        // FIFOCachePrefix fifo缓存前缀
	TopicStdRedisPrefix        = "topic.std."   // 标准发布/订阅
	TopicListRedisPrefix       = "topic.list."  // list 模拟发布/订阅
	ListPrefix                 = "list."        // redis list队列
	HashPrefix                 = "hash."        // redis list队列
	ZSetPrefix                 = "zset."        // redis zset队列
	SessionForMobilePrefix     = "wpapp."       // SessionForMobilePrefix 工作手机团队session前缀（所有团队要统一）
	SessionForQWPrefix         = "qw."          // SessionForQWPrefix 企业微信session前缀（所有团队要统一）
	SessionForTWPrefix         = "tw."          // SessionForTWPrefix 企业微信第三方应用安装自建应用临时session 单独中间件
	SessionForThirdPartyPrefix = "tp."          // SessionForThirdPartyPrefix 企业微信第三方session前缀（所有团队要统一）
	SessionForAsPrefix         = "as."          // SessionForAsPrefix 辅助APP登录前缀（所有团队要统一）
	SessionForWpBillPrefix     = "wpsession."   // SessionForWpBillPrefix 计费团队session前缀（所有团队要统一）
	SessionForMallMemberPrefix = "ml."          // SessionForMallMemberPrefix 商城会员session前缀（所有团队要统一）
	SessionForAgencyPlatform   = "ap."          // 代理商平台session
	WorkphoneSmsCodePrefix     = "wp_sms_code." // 工作手机短信验证码redis前缀
)

// time 缓存时间配置
const (
	MINUTE   = 60          // 分钟
	HOUR     = MINUTE * 60 // 小时
	HALFHOUR = MINUTE * 30 // 半小时
	DAY      = HOUR * 24   // 天

	LRUDuration      = MINUTE * 10 // lru持续时间
	LRUDurationByDay = DAY * 10    // lru持续时间 day

	FIFODuration = MINUTE * 10 // fifo缓存持续时间
)
