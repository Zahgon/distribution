package v2

import (
	"regexp"
)

var (
	reToken            = regexp.MustCompile(`^[^"(),/:;<=>?@[\]{}[:space:][:cntrl:]]+`)
	reQuotedValue      = regexp.MustCompile(`^[^\\"]+`)
	reEscapedCharacter = regexp.MustCompile(`^[[:blank:][:graph:]]`)
)

func parseForwardedHeader(forwarded string) (map[string]string, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}
