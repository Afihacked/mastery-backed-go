package settings

import (
	"os"
)

type Settings struct {
	AppName string
	Debug   bool
	YtCookies string
}

func LoadSettings() Settings {
	debug := false
	if os.Getenv("DEBUG") == "true" {
		debug = true
	}
	return Settings{
		AppName: "MasterY Backend (Go)",
		Debug:   debug,
		YtCookies: os.Getenv("YT_COOKIES"),
	}
}
