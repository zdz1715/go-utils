//go:build windows

package osutils

import (
	"os/exec"
	"strings"

	"golang.org/x/sys/windows/registry"
)

func LocalTimezone() string {
	name, err := localTZFormUtil()
	if err == nil && name != "" {
		return name
	}
	name, err = localTZFromReg()
	if err == nil && name != "" {
		return name
	}
	return ""
}

func localTZFormUtil() (string, error) {
	cmd := exec.Command("tzutil", "/g")
	data, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(data)), nil
}

func localTZFromReg() (string, error) {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SYSTEM\CurrentControlSet\Control\TimeZoneInformation`, registry.QUERY_VALUE)
	if err != nil {
		return "", err
	}
	defer k.Close()

	name, _, err := k.GetStringValue("TimeZoneKeyName")
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(name), nil
}
