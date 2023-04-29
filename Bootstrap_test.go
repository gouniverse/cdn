package cdn

import (
	"strings"
	"testing"
)

func TestBootstrapCss_5_2_3(t *testing.T) {
	output := BootstrapCss_5_2_3()
	expected := "bootstrap@5.2.3/dist/css/bootstrap.min.css"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}

func TestBootstrapJs_5_2_3(t *testing.T) {
	output := BootstrapJs_5_2_3()
	expected := "bootstrap@5.2.3/dist/js/bootstrap.bundle.min.js"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}
