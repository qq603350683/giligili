package constbase

const (
	YES = 1
	NO = 0

	PROP_TYPE_GOLD = "gold"
	PROP_TYPE_DIAMOND = "diamond"
	PROP_TYPE_BULLET_ENHANCER = "bullet_enhancer"
	PROP_TYPE_BULLET_SPEED_ENHANCER = "bullet_speed_enhancer"
	PROP_TYPE_SKILL_ENHANCER = "skill_enhancer"
	PROP_TYPE_SKILL_SPEED_ENHANCER = "skill_speed_enhancer"

	ENHANCER_SUCCESS = "success"
	ENHANCER_FAIL = "fail"

	// 530 ~ 540 是 websocket 专用码
	// 服务端要求客户端关闭连接
	WEBSOCKET_CLOSE = 530
	// 被踢下线
	WEBSOCKET_OFFLINE = 531

	// websocket 接收信息类型
	WEBSOCKET_MESSAGE_TYPE_TEXT = 1
	WEBSOCKET_MESSAGE_TYPE_BINARY = 2
	WEBSOCKET_MESSAGE_TYPE_CLOSE = 8
	WEBSOCKET_MESSAGE_TYPE_PING = 9
	WEBSOCKET_MESSAGE_TYPE_PONG = 10

	// 登录用户信息详情
	LOGIN_USER_INFO = 229
	// 强化成功
	ENHANCER_RESULT = 230
	// 刷副本奖励
	LEVEL_PASS_PRIZE = 231
	// 关卡详情
	LEVEL_INFO = 232

	CHAT_MESSAGE = 250
)
