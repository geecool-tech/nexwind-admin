package resp

import "nexwind.net/admin/server/constant"

var msg = map[int]string{
	constant.CODE_TOKEN_NOT_EXIST: "token not exists.",
	constant.CODE_FAIL:            "request fail.",
	constant.CODE_SUCCESS:         "request success.",
	constant.CODE_UN_AUTH:         "un auth.",
	constant.CODE_INVALID_PARAM:   "invalid param",
	constant.CODE_ILLEGAL_TOKEN:   "illegal token",
}

// GetMsg get response message.
func GetMsg(code int) string {
	return msg[code]
}
