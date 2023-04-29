package cdn

import (
	"strings"
	"testing"
)

func TestFontAwesomeCss_6_1_2(t *testing.T) {
	output := FontAwesomeCss_6_1_2()
	expected := "@fortawesome/fontawesome-free@6.1.2/css/fontawesome.min.cs"
	if !strings.Contains(output, expected) {
		t.Error("Does not contain '" + expected + "', Output:" + output)
	}
}
