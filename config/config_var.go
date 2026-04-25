package config

import "os"

var VERIFY_CAPTCHA = false
var CLIENT_ID = "default"
var CLIENT_SECRET = "secret"

var SSO_ENABLED = false
var SSO_BASE_URL = "http://123.249.5.199:82"
var SSO_APP_KEY = "bf5a75d320c343f7a2536de79d8238a9"
var SSO_APP_SECRET = "e69a071cb9d54f23ac80eb4a92e912de"

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
	if val := os.Getenv("SSO_ENABLED"); val != "" {
		SSO_ENABLED = val == "true"
	}
	if val := os.Getenv("SSO_BASE_URL"); val != "" {
		SSO_BASE_URL = val
	}
	if val := os.Getenv("SSO_APP_KEY"); val != "" {
		SSO_APP_KEY = val
	}
	if val := os.Getenv("SSO_APP_SECRET"); val != "" {
		SSO_APP_SECRET = val
	}
}
