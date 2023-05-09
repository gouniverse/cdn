package cdn

import (
	"strings"
	"testing"
)

func TestJqTreeCss_1_7_0(t *testing.T) {
	output := JqTreeCss_1_7_0()
	expected := "jqtree/1.7.0/jqtree.min.css"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}

func TestJqTreeJs_1_7_0(t *testing.T) {
	output := JqTreeJs_1_7_0()
	expected := "jqtree/1.7.0/tree.jquery.min.js"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}
