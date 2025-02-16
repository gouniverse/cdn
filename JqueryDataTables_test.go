package cdn

import (
	"strings"
	"testing"
)

func TestJqueryDataTablesCss_1_13_4(t *testing.T) {
	output := JqueryDataTablesCss_1_13_4()
	expected := "cdn.datatables.net/1.13.4/css/jquery.dataTables.css"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}

func TestJqueryDataTablesJs_1_13_4(t *testing.T) {
	output := JqueryDataTablesJs_1_13_4()
	expected := "cdn.datatables.net/1.13.4/js/jquery.dataTables.js"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}
