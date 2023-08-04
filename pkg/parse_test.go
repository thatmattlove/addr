package addr_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	addr "github.com/thatmattlove/addr/pkg"
)

func Test_ParseASN(t *testing.T) {
	t.Run("valid asn", func(t *testing.T) {
		t.Parallel()
		asn, err := addr.ParseASN("as14525")
		assert.NoError(t, err)
		assert.Equal(t, uint64(14525), asn)
	})
	t.Run("invalid asn 1", func(t *testing.T) {
		t.Parallel()
		_, err := addr.ParseASN("not an asn")
		assert.Error(t, err)
	})
	t.Run("invalid asn 2", func(t *testing.T) {
		t.Parallel()
		_, err := addr.ParseASN("18446744073709551616")
		assert.Error(t, err)
	})
}
