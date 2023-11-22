//go:build unix && !ios && !android

package osutils

import (
	"github.com/zdz1715/go-utils/fileutils"
	"github.com/zdz1715/go-utils/goutils"
	"github.com/zdz1715/go-utils/version"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func getLinuxNameAndVersion() (name string, ver *version.Version) {
	var verStr string
	_ = fileutils.ReadLineFunc("/etc/os-release", func(num int, line string) bool {
		if s, found := strings.CutPrefix(line, "ID="); found && s != "" {
			name = goutils.TrimSpaceAndQuote(s)
		}
		if s, found := strings.CutPrefix(line, "VERSION_ID="); found && s != "" {
			verStr = goutils.TrimSpaceAndQuote(s)
		}
		return true
	})
	switch name {
	case "centos":
		if vb, err := os.ReadFile("/etc/redhat-release"); err == nil && len(vb) > 0 {
			verStr = string(vb)
		}
	case "debian":
		if vb, err := os.ReadFile("/etc/debian_version"); err == nil && len(vb) > 0 {
			verStr = string(vb)
		}
	}
	ver = version.ParseVersion(verStr)
	return
}

func getMacosNameAndVersion() (name string, ver *version.Version) {
	name = "macOS"
	if out, err := exec.Command("sw_vers", "--productVersion").Output(); err == nil &&
		len(out) > 0 {
		ver = version.ParseVersion(string(out))
	}
	return
}

func getOSNameAndVersion() (name string, ver *version.Version) {
	switch runtime.GOOS {
	case "linux":
		return getLinuxNameAndVersion()
	case "darwin":
		return getMacosNameAndVersion()
	}
	return
}

func initInfo() {
	name, ver := getOSNameAndVersion()
	if ver == nil {
		info.version = new(version.Version)
	} else {
		info.version = ver
	}
	info.name = name
}
