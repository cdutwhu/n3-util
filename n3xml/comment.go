package n3xml

// RmComment :
func RmComment(xml string) string {
	rx := rxMustCompile(`<!--[\s\S]*-->`)
	return RmMultiBlank(rx.ReplaceAllString(xml, ""))
}
