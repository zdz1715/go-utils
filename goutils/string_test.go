package goutils

import (
	"testing"
)

func TestTrimSpaceAndQuote(t *testing.T) {
	tests := []struct {
		value string
		want  string
	}{
		{value: " \"Alpine Linux\" ", want: "Alpine Linux"},
		{value: " Alpine Linux ", want: "Alpine Linux"},
		{value: " 'Alpine Linux' ", want: "Alpine Linux"},
	}

	for _, tt := range tests {
		target := TrimSpaceAndQuote(tt.value)
		if target != tt.want {
			t.Errorf("value=%s target=%s want=%s", tt.value, target, tt.want)
		}
	}
}

func TestHidePassword(t *testing.T) {
	t.Log(HidePassword("s=123456", "123456"))
	t.Log(HidePassword("s1=123 s2=123456", "123", "123456"))
}

func BenchmarkHidePassword(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HidePassword("password=123456", "123456")
	}
}

func BenchmarkTrimSpaceAndQuote(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TrimSpaceAndQuote(" \"Alpine Linux\" ")
	}
}
