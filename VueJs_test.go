package cdn

import (
	"strings"
	"testing"
)

func TestVueJs_3(t *testing.T) {
	output := VueJs_3()
	expected := "vue@3/dist/vue.global.js"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}
