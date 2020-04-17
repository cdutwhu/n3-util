package xml

import (
	"regexp"

	cmn "github.com/cdutwhu/json-util/common"
)

// RmComment :
func RmComment(xml string) string {
	r := regexp.MustCompile(`(?s)<!--.+-->`)
	pairs := r.FindAllStringIndex(xml, -1)
	return cmn.ReplByPosGrp(xml, pairs, []string{""})
}
