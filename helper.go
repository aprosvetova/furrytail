package furrytail

import "time"

var errorTranslations = map[string]string{
	"未查询到设备绑定关系": "device not found",
	" 未知异常 ":     "unknown exception",
	"设备网络异常": "device network exception",
	"设备不在线无法操作": "device is offline",
	"称重操作失败": "weighing error",
}

var ctz, _ = time.LoadLocation("Asia/Shanghai")

func convertToTime(chinese int, loc time.Location) (hour, min int) {
	chinese = chinese/60
	converted := time.Date(2000, time.January, 1, chinese/60, chinese % 60, 0, 0, ctz).In(&loc)
	return converted.Hour(), converted.Minute()
}

func convertToSeconds(hour, min int, loc time.Location) (chinese int) {
	converted := time.Date(2000, time.January, 1, hour, min, 0, 0, &loc).In(ctz)
	return (converted.Hour()*60+converted.Minute())*60
}

func hasWeekday(needle time.Weekday, days []time.Weekday) bool {
	for _, day := range days {
		if needle == day {
			return true
		}
	}
	return false
}

func translateError(chinese string) string {
	english, ok := errorTranslations[chinese]
	if !ok {
		return chinese
	}
	return english
}