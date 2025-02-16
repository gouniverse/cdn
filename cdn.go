package cdn

import "os"

func cdnBase(originalCDN string) string {
	customCDN := os.Getenv("CDN_URL_PREFIX")
	if customCDN != "" {
		return customCDN
	}
	return originalCDN
}
