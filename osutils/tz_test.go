package osutils

import (
	"testing"
)

func TestGetTimezone(t *testing.T) {
	t.Logf("timzeone: %s", GetTimezone())
}

func BenchmarkGetTimezone(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetTimezone()
	}
}
