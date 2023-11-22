package osutils

import "testing"

func TestNewOS(t *testing.T) {
	t.Logf("os: %+v Major: %d timezone: %s\n",
		Info,
		Info.Version().Major(),
		GetTimezone())
}

func BenchmarkNewOS(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Info.get()
	}
}
