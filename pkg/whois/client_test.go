package whois_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thatmattlove/addr/addr/whois"
)

func Test_WhoisClient(t *testing.T) {
	w, err := whois.New("bgp.tools", 43)
	assert.NoError(t, err)
	res, err := w.Query("as14525")
	assert.NoError(t, err)
	assert.Contains(t, res, "Stellar")
}
