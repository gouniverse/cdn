package cdn

import (
	"strings"
	"testing"
)

func TestHtmx_1_9_2(t *testing.T) {
	output := Htmx_1_9_2()
	expected := "htmx.org@1.9.2"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}
