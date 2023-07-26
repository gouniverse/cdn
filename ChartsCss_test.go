package cdn

import (
	"strings"
	"testing"
)

func TestChartsCss_3_12_3(t *testing.T) {
	output := ChartsCss_0_9_3()
	expected := "charts.css@0.9.0/dist/charts.min.css"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}
