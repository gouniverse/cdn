package cdn

import (
	"strings"
	"testing"
)

func TestBootstrapIconsCss_1_10_2(t *testing.T) {
	output := BootstrapIconsCss_1_10_2()
	expected := "bootstrap-icons@1.10.2/font/bootstrap-icons.css"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}

func TestBootstrapIconsCss_1_9_1(t *testing.T) {
	output := BootstrapIconsCss_1_9_1()
	expected := "bootstrap-icons@1.9.1/font/bootstrap-icons.css"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}
