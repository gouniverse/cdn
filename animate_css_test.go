package cdn

import (
	"strings"
	"testing"
)

func TestAnimatedCSS_4_1_1(t *testing.T) {
	output := AnimatedCSS_4_1_1()
	expected := "animate.css/4.1.1/animate.min.css"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}
