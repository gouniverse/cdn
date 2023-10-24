package cdn

import (
	"strings"
	"testing"
)

func TestTailwindCss_3_3_3(t *testing.T) {
	output := TailwindCss_3_3_3()
	expected := "cdn.tailwindcss.com/3.3.3"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}
