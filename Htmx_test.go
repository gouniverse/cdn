package cdn

import (
	"strings"
	"testing"
)

// Keep in reverse alphabetical order (latest version on top)

func TestHtmx_1_9_6(t *testing.T) {
	output := Htmx_1_9_6()
	expected := "htmx.org@1.9.6"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}

func TestHtmx_1_9_5(t *testing.T) {
	output := Htmx_1_9_5()
	expected := "htmx.org@1.9.5"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}

func TestHtmx_1_9_4(t *testing.T) {
	output := Htmx_1_9_4()
	expected := "htmx.org@1.9.4"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}

func TestHtmx_1_9_2(t *testing.T) {
	output := Htmx_1_9_2()
	expected := "htmx.org@1.9.2"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}
