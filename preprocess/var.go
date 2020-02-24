package preprocess

import (
	"fmt"
	"strings"
)

var (
	fPf = fmt.Printf
	fEf = fmt.Errorf

	sReplaceAll = strings.ReplaceAll
	sHasPrefix  = strings.HasPrefix
	sHasSuffix  = strings.HasSuffix
)
