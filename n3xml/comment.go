package n3xml

import (
	"regexp"
)

// RmComment :
func RmComment(xml string) string {
	r := regexp.MustCompile(`(?s)<!--.+-->`)
	pairs := r.FindAllStringIndex(xml, -1)
	return replByPosGrp(xml, pairs, []string{""})
}
