package addr_test

import (
	"fmt"
	"net"
	"testing"

	"github.com/miekg/dns"
	"github.com/stretchr/testify/assert"
	addr "github.com/thatmattlove/addr/pkg"
)

func Test_DNSLookup(t *testing.T) {
	t.Run("trailing period", func(t *testing.T) {
		t.Parallel()
		_, err := addr.DNSLookup[*dns.A]("example.com.", dns.TypeA)
		assert.NoError(t, err)
	})
	t.Run("no trailing period", func(t *testing.T) {
		t.Parallel()
		_, err := addr.DNSLookup[*dns.A]("example.com", dns.TypeA)
		assert.NoError(t, err)
	})
	t.Run("exchange error", func(t *testing.T) {
		original := addr.DNS_SERVER
		addr.DNS_SERVER = "not a server:53"
		_, err := addr.DNSLookup[*dns.A]("1", dns.TypeA)
		addr.DNS_SERVER = original
		assert.Error(t, err)
	})
	t.Run("query error", func(t *testing.T) {
		t.Parallel()
		d := "thisaintnodomain.example.com."
		_, err := addr.DNSLookup[*dns.A](d, dns.TypeA)
		e := addr.NewErrLookupFailure(d, dns.RcodeNameError)
		assert.EqualError(t, err, e.Error())
	})
	t.Run("type assertion error", func(t *testing.T) {
		t.Parallel()
		_, err := addr.DNSLookup[*dns.SOA]("www.google.com", dns.TypeA)
		e := fmt.Errorf("failed to assert response for target 'www.google.com.' as A record")
		assert.EqualError(t, err, e.Error())
	})
}

func Test_DNSReverseLookup(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		t.Parallel()
		ptr := "one.one.one.one."
		ipStr := "1.1.1.1"
		ip := net.ParseIP(ipStr)
		assert.NotNil(t, ip)
		reverse, err := addr.DNSReverseLookup(&ip)
		assert.NoError(t, err)
		assert.Contains(t, reverse, ptr)
	})
	t.Run("invalid ip", func(t *testing.T) {
		t.Parallel()
		_, err := addr.DNSReverseLookup(&net.IP{})
		assert.Error(t, err)
	})
	t.Run("errors with wrong dns server", func(t *testing.T) {
		original := addr.DNS_SERVER
		addr.DNS_SERVER = "not a server:53"
		ip := net.ParseIP("1.1.1.1")
		_, err := addr.DNSReverseLookup(&ip)
		assert.Error(t, err)
		addr.DNS_SERVER = original
	})
}

func Test_DNSForwardLookup(t *testing.T) {
	t.Run("dual", func(t *testing.T) {
		t.Parallel()
		ip4, ip6, err := addr.DNSForwardLookup("one.one.one.one")
		assert.NoError(t, err)
		ip4s := []string{}
		for _, i := range ip4 {
			ip4s = append(ip4s, i.String())
		}
		ip6s := []string{}
		for _, i := range ip6 {
			ip6s = append(ip6s, i.String())
		}
		assert.Contains(t, ip4s, "1.1.1.1")
		assert.Contains(t, ip4s, "1.0.0.1")
		assert.Contains(t, ip6s, "2606:4700:4700::1111")
		assert.Contains(t, ip6s, "2606:4700:4700::1001")
	})
	t.Run("lookup errors", func(t *testing.T) {
		_, _, err := addr.DNSForwardLookup("this will fail")
		assert.Error(t, err)
	})
}
