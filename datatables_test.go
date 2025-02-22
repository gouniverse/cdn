package cdn

import (
	"strings"
	"testing"
)

func TestJqueryDataTablesCss_2_2_2(t *testing.T) {
	output := JqueryDataTablesCss_2_2_2()
	expected := "https://cdn.datatables.net/2.2.2/css/dataTables.dataTables.min.css"
	if !strings.Contains(output, expected) {
		t.Errorf("Does not contain '%s', Output: %s", expected, output)
	}
}

func TestJqueryDataTablesJs_2_2_2(t *testing.T) {
	output := JqueryDataTablesJs_2_2_2()
	expected := "https://cdn.datatables.net/2.2.2/js/dataTables.min.js"
	if !strings.Contains(output, expected) {
		t.Errorf("Does not contain '%s', Output: %s", expected, output)
	}
}
