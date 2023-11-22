package osutils

import (
	"fmt"
	"github.com/zdz1715/go-utils/version"
	"runtime"
	"strings"
	"sync"
)

var Info *OS = &info

var info OS
var infoOnce sync.Once

func (o *OS) get() *OS {
	if o == &info {
		infoOnce.Do(initInfo)
	}
	return o
}

type OS struct {
	name    string
	version *version.Version
}

func (o *OS) String() string {
	return o.get().string()
}

func (o *OS) string() string {
	builder := new(strings.Builder)
	if o.name != "" {
		builder.WriteString(o.name)
		builder.WriteString(" ")
	}
	if o.version != nil {
		builder.WriteString(o.version.String())
		builder.WriteString(" ")
	}
	builder.WriteString(o.OSArch())
	return builder.String()
}

func (o *OS) OSArch() string {
	return fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH)
}

func (o *OS) Name() string {
	return o.get().name
}

func (o *OS) Version() *version.Version {
	return o.get().version
}

func (o *OS) Timezone() string {
	return GetTimezone()
}
