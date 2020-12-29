package util

import (
	"fmt"
	"time"

	"github.com/golang/glog"
)

const (
	layout1 = "2006-01-02T15:04:05.000Z"
	layout2 = "2006-01-02 15:04:05"
)

var (
	//時區
	Location, _ = time.LoadLocation("Asia/Taipei")
	// Location, _ = time.LoadLocation("UTC")
	//現在時間
	// nowT = GetNow()
)

func GetNow() time.Time {
	return time.Now().In(Location)
}

//取得月份起點時間
func GetMonthTime(t time.Time) time.Time {
	year, month, _ := t.Date()
	return time.Date(year, month, 0, 0, 0, 0, 0, t.Location())
}

//取得日期起點時間
func GetDayTime(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, t.Location())
}

//取得今日起點時間
func GetTodayTime() time.Time {
	// timeStr2 := "2020-09-18 00:00:00"
	// t, _ := time.Parse(layout2, timeStr2)
	t := GetNow()
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, GetNow().Location())
}

func Tomorrow() time.Time {
	t := GetNow()
	//加一天
	tomoro := t.AddDate(0, 0, 1)
	return tomoro
}

func ParseTimeTest(day time.Time) time.Time {
	timeStr1 := "2014-12-03T01:11:18.965Z"
	t, err := time.Parse(layout1, timeStr1)
	if err != nil {
		glog.Error(err)
	}

	// timeStr2 := "2020-09-17 00:00:00"
	// t, err := time.Parse(layout2, timeStr2)
	// if err != nil {
	// 	glog.Error(err)
	// }

	fmt.Println(t)
	return t
}

//day可放任意日期(要有dash) ex:2020-01-12
func ParseYourTime(day string) time.Time {
	// timeStr1 := "2014-12-03T01:11:18.965Z"
	// t, err := time.Parse(layout1, timeStr1)
	// if err != nil {
	// 	glog.Error(err)
	// }

	timeStr2 := day + " 00:00:00"
	fmt.Println(timeStr2)
	t, err := time.Parse(layout2, timeStr2)
	if err != nil {
		glog.Error(err)
	}

	fmt.Println(t)
	return t
}

//time to specific format, return string
func FmtTime(t time.Time) time.Time {
	ft := t.Format(layout2) //string
	// fmt.Println(ft)
	tft := ParseTime(ft) //string to time
	fmt.Println("formatted time", tft)
	return tft
}

// string to specific format time, return time
func ParseTime(ts string) time.Time {
	t, err := time.Parse(layout2, ts)
	if err != nil {
		glog.Error(err)
	}
	// fmt.Println(t)
	return t
}

func FmtDuration(d time.Duration) string {
	// fmt.Println("duration:", d)
	d = d.Round(time.Second)
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute
	d -= m * time.Minute
	s := d / time.Second

	str := fmt.Sprintf("%02d:%02d:%02d", h, m, s)
	// fmt.Println(str)
	return str
}

func DurationToSecs(d time.Duration) float64 {
	// fmt.Println("duration:", d)
	d = d.Round(time.Second)

	// duration轉string
	// s := d / time.Second
	// str := fmt.Sprintf("%02d", s)
	// fmt.Println(str)

	//duration轉number
	seconds := float64(d / time.Second)

	return seconds
}

// t := m.Ts.Format("2006-01-02 15:04:05")
// dayTime := m.Ts.Truncate(24 * time.Hour)
// fmt.Println(dayTime)
// fmt.Println(m.Ts.Unix())

// func peterTime(){
// 	TimeString := fmt.Sprintf("%02d-%02d-%02dT%02d:%02d:%02d+08:00", 2020, 9, 17, 10, 00, 00)
// 	formatT, _ := time.Parse(time.RFC3339, TimeString)
// }

/*
	//格式化持續時間方法失敗
		//time calculate sample
		t1 := time.Now()
		t2 := t1.Add(time.Second * 341)
		fmt.Println(t1)
		fmt.Println(t2)
		diff := t2.Sub(t1)
		fmt.Println(diff)
*/

// //sample
// func T() {
// 	const base_format = "2006-01-02 15:04:05"
// 	// const base_format = "2006/01/02 15:04:05"
// 	//今日工時給10

// 	//現在時間
// 	nt := time.Now()
// 	//--------轉換為時間格式字串
// 	fnt := nt.Format(base_format)
// 	fmt.Printf("現在時間:%v\n", fnt)

// 	//--------時間字串轉換為日期格式
// 	parse_str_time, _ := time.Parse(base_format, fnt)
// 	fmt.Printf("string to datetime :%v\n", parse_str_time)

// 	//上班時間
// 	//获取今天08:00:00秒的时间戳
// 	startWorkT := time.Date(nt.Year(), nt.Month(), nt.Day(), 8, 0, 0, 0, nt.Location())
// 	fmt.Printf("上班時間:%v\n", startWorkT)
// 	fmt.Println(startWorkT.Format(base_format))

// 	//目前為止工時(duration)
// 	workHour := nt.Sub(startWorkT)
// 	fmt.Println("目前持續工時", workHour.String())

// 	//現在完成率
// 	// count/60000

// 	//預計完成率

// }
