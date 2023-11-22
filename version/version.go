package version

import (
	"strconv"
	"strings"
	"unicode"
)

type status uint8

const (
	majorSet status = 1 << iota
	minorSet
	patchSet

	allSet = majorSet | minorSet | patchSet
)

type Version struct {
	major  int
	minor  int
	patch  int
	status status
	latest bool
}

func (v *Version) String() string {
	if v.latest {
		return "latest"
	}
	if v.status&majorSet == 0 {
		return ""
	}
	builder := new(strings.Builder)
	builder.WriteString(strconv.Itoa(v.major))
	if v.status&minorSet != 0 {
		builder.WriteByte('.')
		builder.WriteString(strconv.Itoa(v.minor))
	}
	if v.status&patchSet != 0 {
		builder.WriteByte('.')
		builder.WriteString(strconv.Itoa(v.patch))
	}
	return builder.String()
}

func (v *Version) Version(prefix ...string) string {
	str := v.String()
	if str == "latest" {
		return str
	}
	if len(prefix) > 0 {
		return prefix[0] + str
	}
	return "v" + str
}

// Older returns true if this version v is older than the other.
func (v *Version) Older(other *Version) bool {
	if v.latest || other == nil { // Latest is always consider newer, even than future versions.
		return false
	}
	if other.latest {
		return true
	}
	if v.major != other.major {
		return v.major < other.major
	}

	if v.minor != other.minor {
		return v.minor < other.minor
	}

	return v.patch < other.patch
}

func (v *Version) Major() int {
	return v.major
}

func (v *Version) Minor() int {
	return v.minor
}

func (v *Version) Patch() int {
	return v.patch
}

func (v *Version) Latest() bool {
	return v.latest
}

func (v *Version) addNum(n int) {
	if v.status&majorSet == 0 {
		v.major = n
		v.status = majorSet
		return
	}

	if v.status&minorSet == 0 {
		v.minor = n
		v.status = majorSet | minorSet
		return
	}

	if v.status&patchSet == 0 {
		v.patch = n
		v.status = allSet
		return
	}
}

func (v *Version) add(str string) {
	n, err := strconv.Atoi(str)
	if err == nil {
		v.addNum(n)
	}
}

func (v *Version) complete() bool {
	return v.status == allSet
}

func ParseVersion(v string) *Version {
	ver := new(Version)
	if v == "" {
		return ver
	}
	if v == "latest" {
		ver.latest = true
		return ver
	}
	index := -1
	endIndex := len(v) - 1
	skip := false
	for i, r := range v {
		if ver.complete() {
			break
		}
		newNum := false
		if unicode.IsDigit(r) {
			if index < 0 && !skip {
				index = i
			}
			if i == endIndex && index >= 0 {
				ver.add(v[index:])
				break
			}
		} else {
			newNum = true
			if r == '.' {
				skip = false
			}
		}

		if newNum && index >= 0 {
			ver.add(v[index:i])
			index = -1
			if r != '.' {
				skip = true
			}
		}

	}

	return ver
}

// New 可以传负数忽略该版本数字，比如：New(10, 2, -1) string: 10.2
func New(num ...int) *Version {
	ver := new(Version)
	for _, n := range num {
		if ver.complete() {
			break
		}
		ver.addNum(n)
	}
	return ver
}

func NewLatest() *Version {
	return &Version{
		latest: true,
	}
}
