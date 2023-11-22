//go:build unix && !ios && !android

package osutils

import (
	"os"
	"strings"
)

var (
	localtime    = "/etc/localtime"
	zoneinfoPath = "zoneinfo/"
)

func LocalTimezone() string {
	// 不是软连接则获取不到，直接返回utc
	p, err := os.Readlink(localtime)
	if err != nil || p == "" {
		return UTC
	}
	i := strings.LastIndex(p, zoneinfoPath)
	if i < 0 {
		return UTC
	}
	return p[i+len(zoneinfoPath):]
}
