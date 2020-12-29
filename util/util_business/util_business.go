package util_business

import "time"

//取得預設修復時間
func GetRepairTime(nowtime time.Time) time.Time {
	// 預設修復時間為當日晚上6點(utc + 8)
	repairTime := time.Date(nowtime.Year(), nowtime.Month(), nowtime.Day(), 18, 0, 0, 0, nowtime.Location()) //注意這裡是根據nowtime的時區

	// 如果現在時間超過當日晚上6點則進位一天
	if nowtime.After(repairTime) {
		repairTime = repairTime.Add(time.Hour * 24)
	}
	return repairTime
}
