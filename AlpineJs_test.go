package cdn

import (
	"strings"
	"testing"
)

func TestAlpine_3_12_3(t *testing.T) {
	output := AlpineJs_3_12_3()
	expected := "alpinejs@3.12.3/dist/cdn.min.js"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}
