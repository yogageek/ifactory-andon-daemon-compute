package setenv

import (
	"fmt"
	"os"
)

func SetEnv() {
	// 外部env的值會先吃,才會進來這,所以這裡的setEnv會把外部env覆蓋掉
	// 部屬時候也可以DEBUG參數不要設 直接給ENV
	if os.Getenv("DEBUG") == "true" {
		fmt.Println("debug=true")
		os.Setenv("MONGODB_URL", "52.187.110.12:27017")
		os.Setenv("MONGODB_DATABASE", "6dba1e66-a658-445a-b4cc-cb9602bb2d3e")
		os.Setenv("MONGODB_USERNAME", "5845a1b5-c3e2-438e-9157-a14a559675ad")
		os.Setenv("MONGODB_PASSWORD", "RS15YM9WUibOj79kAF86SfaR")
		os.Setenv("NOTIFI_URL", "https://ifps-andon-api-ifpsdev-eks005.sa.wise-paas.com/andon/api/v1.0/notification")
	} else if os.Getenv("DEBUG") == "false" {
		fmt.Println("debug=false")
		os.Setenv("MONGODB_URL", "52.187.110.12:27017")
		os.Setenv("MONGODB_DATABASE", "87e1dc58-4c20-4e65-ad81-507270f6bdac")
		os.Setenv("MONGODB_USERNAME", "19e0ce80-af51-404c-8d55-9edefcbd4bdf")
		os.Setenv("MONGODB_PASSWORD", "TYyvTeVemAlJzzuq4w3sBr2D")
		os.Setenv("NOTIFI_URL", "https://ifps-andon-api-ifpsdemo-eks005.sa.wise-paas.com/andon/api/v1.0/notification")
	} else if os.Getenv("DEBUG") == "tg_release" {
		fmt.Println("debug=tg_release")
		os.Setenv("MONGODB_URL", "40.65.167.39:27017")
		os.Setenv("MONGODB_DATABASE", "99569d89-366f-4a92-aef3-e17fc521d370")
		os.Setenv("MONGODB_USERNAME", "54e9d215-fbf5-4591-87d0-7ad0be51227f")
		os.Setenv("MONGODB_PASSWORD", "1nsEGzF6gmF70E2BlA1PP25T")
		os.Setenv("NOTIFI_URL", "https://ifps-andon-api-tienkang-eks002.sa.wise-paas.com/andon/api/v1.0/notification")
	}
}
