package osutils

import (
	"os"
	"time"
)

var UTC = "UTC"

func GetTimezone() string {
	if name, ok := os.LookupEnv("TZ"); ok {
		if name == "" {
			return UTC
		}
		_, err := time.LoadLocation(name)
		if err != nil {
			return UTC
		}
		return name
	}
	return LocalTimezone()
}
