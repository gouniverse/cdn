package cdn

import "net/url"

func GoogleFont(family string, weight string) string {
	familyEscaped := url.QueryEscape(family)
	weightEscaped := url.QueryEscape(weight)
	return `https://fonts.googleapis.com/css2?family=` + familyEscaped + `:wght@` + weightEscaped + `&display=swap`
}
