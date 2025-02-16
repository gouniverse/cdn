package cdn

import (
	"strings"
	"testing"
)

func TestVueTrumbowyg_4_0_0(t *testing.T) {
	output := VueTrumbowyg_4_0_0()
	expected := "vue-trumbowyg@4.0.0/dist/vue-trumbowyg.min.js"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}
