package osutils

import (
	"fmt"
	"testing"
)

func TestLocalTimezone(t *testing.T) {
	fmt.Println(LocalTimezone())
}
