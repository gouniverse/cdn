package cdn

import (
	"strings"
	"testing"
)

// Keep in reverse alphabetical order (latest version on top)

func TestSweetalert2_11(t *testing.T) {
	output := Sweetalert2_11()
	expected := "sweetalert2@11"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}

func TestSweetalert2_10(t *testing.T) {
	output := Sweetalert2_10()
	expected := "sweetalert2@10"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}
