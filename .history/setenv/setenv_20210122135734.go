package setenv

import (
	"os"
)

func SetEnv() {
	// 外部env的值會先吃,才會進來這,所以這裡的setEnv會把外部env覆蓋掉
	if os.Getenv("DEBUG") == "true" {
		os.Setenv("MONGODB_URL", "52.187.110.12:27017")
		os.Setenv("MONGODB_DATABASE", "ifp-data-hub-dev")
		os.Setenv("MONGODB_USERNAME", "e270673c-ce08-4c35-93e2-333ed103736f")
		os.Setenv("MONGODB_PASSWORD", "VUSkt9bbTKSTzb7ZArp36jLk")
	} else if os.Getenv("DEBUG") == "tg" {
		os.Setenv("MONGODB_URL", "52.187.110.12:27017")
		os.Setenv("MONGODB_DATABASE", "tienkang")
		os.Setenv("MONGODB_USERNAME", "50cc2508-87cb-441e-9d61-9bcd943686f6")
		os.Setenv("MONGODB_PASSWORD", "yYRiiFW0Kh0ejyLoCfEC3bNB")

	} else {
		os.Setenv("MONGODB_URL", "52.187.110.12:27017")
		os.Setenv("MONGODB_DATABASE", "ifp-data-hub")
		os.Setenv("MONGODB_USERNAME", "8676b401-a6ce-417b-a0f0-8dee6dee0a67")
		os.Setenv("MONGODB_PASSWORD", "9TuSZ7CD3ah0aQmdHbGqNjrr")
	}
}
