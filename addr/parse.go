package addr

import (
	"fmt"
	"regexp"
	"strconv"
)

func ParseASN(in string) (uint64, error) {
	p := regexp.MustCompile(`([0-9]+)`)
	matches := p.FindStringSubmatch(in)
	if len(matches) == 0 {
		err := fmt.Errorf("failed to parse '%s' as ASN", in)
		return 0, err
	}
	asn, err := strconv.ParseUint(matches[0], 10, 64)
	if err != nil {
		return 0, err
	}
	return asn, nil
}
