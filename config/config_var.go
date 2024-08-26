package config

import "os"

var VERIFY_CAPTCHA = false

func init() {
	if val := os.Getenv("VERIFY_CAPTCHA"); val != "" {
		VERIFY_CAPTCHA = val == "true"
	}
}
