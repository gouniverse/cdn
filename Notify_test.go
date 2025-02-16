package cdn

import (
	"strings"
	"testing"
)

func TestNotify_0_4_2(t *testing.T) {
	output := Notify_0_4_2()
	expected := "cdnjs.cloudflare.com/ajax/libs/notify/0.4.2/notify.min.js"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}
