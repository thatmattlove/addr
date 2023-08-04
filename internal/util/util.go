package util

import (
	"errors"
	"net"
	"os"
	"regexp"
)

const (
	ASN_PATTERN string = `(AS|as)?([0-9]+)`
)

func PathExists(n string) bool {
	if _, err := os.Stat(n); errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}

func IsASN(v string) bool {
	ip := net.ParseIP(v)
	if ip != nil {
		return false
	}
	p := regexp.MustCompile(ASN_PATTERN)
	return p.MatchString(v)
}

func IsIP(v string) bool {
	_, _, err := net.ParseCIDR(v)
	if err == nil {
		return true
	}
	i := net.ParseIP(v)
	return i != nil
}
