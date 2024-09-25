package config

import "os"

var VERIFY_CAPTCHA = false
var CLIENT_ID = "default"
var CLIENT_SECRET = "secret"

func init() {
	if val := os.Getenv("VERIFY_CAPTCHA"); val != "" {
		VERIFY_CAPTCHA = val == "true"
	}
	if val := os.Getenv("CLIENT_ID"); val != "" {
		CLIENT_ID = val
	}
	if val := os.Getenv("CLIENT_SECRET"); val != "" {
		CLIENT_SECRET = val
	}
}
