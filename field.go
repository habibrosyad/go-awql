package awql

import (
	"strings"
	"unicode"
)

// Return map and list format of awql fields
func getFields(q string) []string {
	// Find start and end of field definitions in the query.
	qlow := strings.ToLower(q)
	start := strings.Index(qlow, "select")
	end := -1
	for _, sep := range []string{"from", "where", "during", "order by", "limit"} {
		end = strings.Index(qlow, sep)
		if end != -1 {
			break
		}
	}

	// Clean whitespaces.
	tmp := strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, q[start+6:end])

	return strings.Split(tmp, ",")
}
