package addr_test

import (
	"net"
	"testing"

	"github.com/biter777/countries"
	"github.com/stretchr/testify/assert"
	addr "github.com/thatmattlove/addr/pkg"
	goasn "github.com/thatmattlove/go-asn"
)

const (
	RES_VALID string = `AS      | IP               | BGP Prefix          | CC | Registry | Allocated  | AS Name
	13335   | 1.1.1.0          | 1.1.1.0/24          | US | ARIN     | 2010-07-14 | Cloudflare, Inc.`

	RES_WARNING string = `Warning: some warning
	AS      | IP               | BGP Prefix          | CC | Registry | Allocated  | AS Name
	13335   | 1.1.1.0          | 1.1.1.0/24          | US | ARIN     | 2010-07-14 | Cloudflare, Inc.`

	RES_EMPTY = ``

	RES_TOO_MANY_COLUMNS string = `AS      | IP               | BGP Prefix          | CC | Registry | Allocated  | AS Name | Some Column
	13335   | 1.1.1.0          | 1.1.1.0/24          | US | ARIN     | 2010-07-14 | Cloudflare, Inc. | Some Value`

	RES_INVALID_TIME string = `AS      | IP               | BGP Prefix          | CC | Registry | Allocated  | AS Name
	13335   | 1.1.1.0          | 1.1.1.0/24          | US | ARIN     | not a time | Cloudflare, Inc.`

	RES_INVALID_IP string = `AS      | IP               | BGP Prefix          | CC | Registry | Allocated  | AS Name
	13335   | invalid-ip          | 1.1.1.0/24          | US | ARIN     | 2010-07-14 | Cloudflare, Inc.`

	RES_INVALID_PREFIX = `AS      | IP               | BGP Prefix          | CC | Registry | Allocated  | AS Name
	13335   | 1.1.1.0          | invalid-prefix          | US | ARIN     | 2010-07-14 | Cloudflare, Inc.`

	RES_INVALID_ASN = `AS      | IP               | BGP Prefix          | CC | Registry | Allocated  | AS Name
	invalid-asn   | 1.1.1.0          | invalid-prefix          | US | ARIN     | 2010-07-14 | Cloudflare, Inc.`
)

func Test_QueryASN(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		t.Parallel()
		q := "as14525"
		asn, err := addr.QueryASN(q)
		assert.NoError(t, err, q)
		assert.True(t, asn.ASN.Equal(goasn.MustParse("14525")), q)
		assert.Equal(t, "ARIN", asn.Registry, q)
		assert.Equal(t, countries.USA, asn.Country, q)
		assert.Equal(t, "Stellar Technologies Inc.", asn.Name, q)
	})
	t.Run("invalid asn", func(t *testing.T) {
		t.Parallel()
		_, err := addr.QueryASN("this will fail")
		assert.Error(t, err)
	})
	t.Run("invalid whois client", func(t *testing.T) {
		original := addr.WHOIS_HOST
		addr.WHOIS_HOST = "fake"
		_, err := addr.QueryASN("as14525")
		assert.Error(t, err)
		addr.WHOIS_HOST = original
	})
	t.Run("invalid query", func(t *testing.T) {
		original := addr.WHOIS_HOST
		addr.WHOIS_HOST = "whois.arin.net"
		_, err := addr.QueryASN("14525")
		assert.Error(t, err)
		addr.WHOIS_HOST = original
	})
}

func Test_QueryIPPrefix(t *testing.T) {
	t.Run("ip", func(t *testing.T) {
		t.Parallel()
		i := net.ParseIP("1.1.1.1")
		data, err := addr.QueryIPPrefix(i.String())
		assert.NoError(t, err)
		assert.NotNil(t, data)
		assert.NotNil(t, data.IP)
		assert.NotNil(t, data.Prefix)
		assert.Equal(t, i.String(), data.IP.String())
		assert.Equal(t, "1.1.1.0/24", data.Prefix.String())
	})
	t.Run("cidr", func(t *testing.T) {
		t.Parallel()
		pfx := "1.1.1.0/24"
		data, err := addr.QueryIPPrefix(pfx)
		assert.NoError(t, err)
		assert.NotNil(t, data)
		assert.NotNil(t, data.IP)
		assert.NotNil(t, data.Prefix)
		assert.Equal(t, pfx, data.Prefix.String())
	})
	t.Run("invalid query", func(t *testing.T) {
		t.Parallel()
		_, err := addr.QueryIPPrefix("invalid ip")
		assert.Error(t, err)
	})

	t.Run("non-query result", func(t *testing.T) {
		t.Parallel()
		ip := "169.254.0.1"
		data, err := addr.QueryIPPrefix(ip)
		assert.NoError(t, err)
		assert.Equal(t, ip, data.IP.String())
	})
	t.Run("invalid whois client", func(t *testing.T) {
		original := addr.WHOIS_HOST
		addr.WHOIS_HOST = "fake"
		_, err := addr.QueryIPPrefix("1.1.1.0/24")
		assert.Error(t, err)
		addr.WHOIS_HOST = original
	})
	t.Run("invalid query", func(t *testing.T) {
		original := addr.WHOIS_HOST
		addr.WHOIS_HOST = "whois.arin.net"
		_, err := addr.QueryASN("1.1.1.1")
		assert.Error(t, err)
		addr.WHOIS_HOST = original
	})
}

func Test_ParseResponse(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		t.Parallel()
		_, err := addr.ParseResponse(RES_VALID)
		assert.NoError(t, err)
	})
	t.Run("with warning", func(t *testing.T) {
		t.Parallel()
		_, err := addr.ParseResponse(RES_WARNING)
		assert.NoError(t, err)
	})
	t.Run("empty", func(t *testing.T) {
		t.Parallel()
		_, err := addr.ParseResponse(RES_EMPTY)
		assert.Error(t, err)
	})
	t.Run("too many columns", func(t *testing.T) {
		t.Parallel()
		_, err := addr.ParseResponse(RES_TOO_MANY_COLUMNS)
		assert.Error(t, err)
	})
	t.Run("invalid time", func(t *testing.T) {
		_, err := addr.ParseResponse(RES_INVALID_TIME)
		assert.Error(t, err)
	})
	t.Run("invalid ip", func(t *testing.T) {
		_, err := addr.ParseResponse(RES_INVALID_IP)
		assert.Error(t, err)
	})
	t.Run("invalid prefix", func(t *testing.T) {
		_, err := addr.ParseResponse(RES_INVALID_PREFIX)
		assert.Error(t, err)
	})
	t.Run("invalid asn", func(t *testing.T) {
		_, err := addr.ParseResponse(RES_INVALID_ASN)
		assert.Error(t, err)
	})
}
