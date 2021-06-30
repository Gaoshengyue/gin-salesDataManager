package e

var MsgFlags = map[int]string{
	SUCCESS:                           "ok",
	ERROR:                             "fail",
	INVALID_PARAMS:                    "请求参数错误",
	ERROR_EXIST_CALCULATION_TSR_GRADE: "计算坐席评分错误",
	ERROR_AUTH_CHECK_TOKEN_FAIL:       "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT:    "Token已超时",
	ERROR_AUTH_TOKEN:                  "Token生成失败",
	ERROR_AUTH:                        "Token错误",
	ERROR_EXIST_CALCULATION_TSR_TOP:   "计算坐席排名错误",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
