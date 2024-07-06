package cdn

import (
	"strings"
	"testing"
)

// Keep in reverse alphabetical order (latest version on top)

func TestJQuery3_7_1(t *testing.T) {
	output := Jquery_3_7_1()
	expected := "jquery-3.7.1.min.js"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}

func TestJQuery3_6_4(t *testing.T) {
	output := Jquery_3_6_4()
	expected := "jquery-3.6.4.min.js"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}
