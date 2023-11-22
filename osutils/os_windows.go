//go:build windows

package osutils

import (
	"github.com/zdz1715/go-utils/version"
	"golang.org/x/sys/windows/registry"
	"strconv"
)

func getOSNameAndVersion() (name string, ver *version.Version) {
	name = "windows"
	nums := versionFromReg()
	if len(nums) > 0 {
		ver = version.New(nums...)
	}
	return
}

func versionFromReg() []int {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows NT\CurrentVersion`, registry.QUERY_VALUE)
	if err != nil {
		return nil
	}
	defer k.Close()
	num := make([]int, 0, 3)
	major, _, err := k.GetIntegerValue("CurrentMajorVersionNumber")
	if err != nil {
		return num
	}
	num = append(num, int(major))
	minor, _, err := k.GetIntegerValue("CurrentMinorVersionNumber")
	if err != nil {
		return num
	}
	num = append(num, int(minor))
	build, _, err := k.GetStringValue("CurrentBuildNumber")
	if err != nil {
		return num
	}
	buildNum, _ := strconv.Atoi(build)
	return append(num, buildNum)
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
