package cdn

import (
	"os"
	"testing"
)

func TestCdnBase(t *testing.T) {
	originalCDN := "https://unpkg.com/"
	customCDN := "https://my-custom-cdn.com/"

	// Test case 1: CDN_URL_PREFIX is not set
	os.Unsetenv("CDN_URL_PREFIX")
	result := cdnBase(originalCDN)
	if result != originalCDN {
		t.Errorf("Expected %s, got %s", originalCDN, result)
	}

	// Test case 2: CDN_URL_PREFIX is set
	os.Setenv("CDN_URL_PREFIX", customCDN)
	defer os.Unsetenv("CDN_URL_PREFIX")
	result = cdnBase(originalCDN)
	if result != customCDN {
		t.Errorf("Expected %s, got %s", customCDN, result)
	}
}
