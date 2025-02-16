package cdn

import (
	"strings"
	"testing"
)

func TestVueElementPlusJs_2_3_8(t *testing.T) {
	output := VueElementPlusJs_2_3_8()
	expected := "element-plus@2.3.8/dist/index.full.min.js"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}
