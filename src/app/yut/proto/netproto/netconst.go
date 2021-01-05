package netproto

const (
	NET_STATUS_UNKNOWN = -1   // 未知错误
	NET_STATUS_INVALID_PARAM = 1 // 参数错误
	NET_STATUS_NOT_LOGIN = 101 // 未登录
	NET_STATUS_NO_AUTH   = 102 // 没有访问权限
)