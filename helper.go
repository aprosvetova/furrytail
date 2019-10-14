package furrytail

var errorTranslations = map[string]string{
	"未查询到设备绑定关系": "device not found",
	" 未知异常 ":     "unknown exception",
}

func convertToTime(chinese int) (hour, min int) {
	chinese += 68400
	if chinese > 86400 {
		chinese -= 86400
	}
	chinese /= 60
	hour = chinese / 60
	min = chinese - hour*60
	return
}

func convertToChinese(hour, min int) (chinese int) {
	chinese = (hour*60+min)*60 - 68400
	if chinese < 0 {
		chinese = 86400 + chinese
	}
	return
}

func translateError(chinese string) string {
	english, ok := errorTranslations[chinese]
	if !ok {
		return chinese
	}
	return english
}
