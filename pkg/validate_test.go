package addr_test

import (
	"fmt"
	"log"
	"net"
	"testing"

	"github.com/biter777/countries"
	"github.com/stretchr/testify/assert"
	addr "github.com/thatmattlove/addr/pkg"
)

func TestIPValidator_NewIPValidator(t *testing.T) {
	t.Run("new validator from valid IP", func(t *testing.T) {
		t.Parallel()
		ipStr := "192.0.2.1"
		validator, err := addr.NewIPValidator(ipStr)
		assert.NoError(t, err)
		assert.NotNil(t, validator)
	})
	t.Run("new validator from valid CIDR", func(t *testing.T) {
		t.Parallel()
		ipStr := "192.0.2.0/24"
		validator, err := addr.NewIPValidator(ipStr)
		assert.NoError(t, err)
		assert.NotNil(t, validator)
	})
	t.Run("error from invalid input", func(t *testing.T) {
		t.Parallel()
		ipStr := "this is not an IP"
		validator, err := addr.NewIPValidator(ipStr)
		assert.Error(t, err)
		assert.Nil(t, validator)
	})
}

func Test_IsIPv6(t *testing.T) {
	ip4 := net.ParseIP("192.0.2.1")
	ip6 := net.ParseIP("2001:db8::1")
	t.Run("is ipv6", func(t *testing.T) {
		t.Parallel()
		result := addr.IsIPv6(ip6)
		assert.True(t, result)
	})
	t.Run("is not ipv6", func(t *testing.T) {
		t.Parallel()
		result := addr.IsIPv6(ip4)
		log.Println(len(ip4))
		assert.False(t, result)
	})
}

func Test_GetNonGlobalPrefix(t *testing.T) {
	type casesT struct {
		idx    int
		ip     string
		global bool
	}
	cases := []casesT{
		{0, "169.254.128.1", false},
		{1, "fe80::ff5:ff5:ff5:1", false},
		{2, "10.1.2.3", false},
		{3, "172.19.20.5", false},
		{4, "192.168.105.3", false},
		{5, "100.66.255.2", false},
		{6, "192.0.0.89", false},
		{7, "192.31.196.20", false},
		{8, "192.175.48.164", false},
		{9, "2001:4:112::abcd", false},
		{10, "2620:4f:8000::abcd", false},
		{11, "192.52.193.229", false},
		{12, "2001:3::abcd:0ff5", false},
		{13, "192.88.99.140", false},
		{14, "2002::0ff5:abcd", false},
		{15, "198.19.5.7", false},
		{16, "2001:2::0ff5:abcd", false},
		{17, "192.0.2.18", false},
		{18, "198.51.100.26", false},
		{19, "203.0.113.79", false},
		{20, "240.0.255.3", false},
		{21, "0.0.0.100", false},
		{22, "224.128.64.32", false},
		{23, "ff00::0ff5:abcd", false},
		{24, "127.0.0.5", false},
		{25, "::1", false},
		{26, "0.0.0.0", false},
		{27, "::", false},
		{28, "fc00::1234:5678:9abc:deff", false},
		{29, "100::abcd", false},
		{30, "2001::abcd", false},
		{31, "2001:10::abcd", false},
		{32, "2001:20::0ff5", false},
		{33, "2001:db8::abcd:ef9a", false},
		{34, "64:ff9b::abcd:ef9a", false},
		{35, "2001:30::abcd:ef9a:bcde", false},
		{36, "1.1.1.1", true},
		{37, "2606:4700:4700::1111", true},
		{38, "199.34.92.255", true},
		{39, "2604:c0c0:1000::0ff5", true},
	}
	for _, case_ := range cases {
		case_ := case_
		t.Run(fmt.Sprint(case_.idx), func(t *testing.T) {
			t.Parallel()
			net_, str := addr.GetNonGlobalPrefix(net.ParseIP(case_.ip))
			if case_.global {
				assert.Nil(t, net_, str)
			} else {
				assert.IsType(t, &net.IPNet{}, net_, str)
			}
		})
	}
}

func TestIPValidator_Validate(t *testing.T) {
	t.Run("ip4 non-global has fallback response", func(t *testing.T) {
		t.Parallel()
		ipStr := "169.254.100.2"
		ip := net.ParseIP(ipStr)
		v, _ := addr.NewIPValidator(ipStr)
		shouldQuery, response := v.Validate()
		assert.False(t, shouldQuery)
		assert.IsType(t, &addr.Response{}, response)
		assert.Equal(t, uint64(0), response.ASN)
		assert.Equal(t, addr.TXT_LINK_LOCAL, response.Name)
		assert.Equal(t, countries.USA, response.Country)
		assert.Equal(t, &ip, response.IP)
		assert.Equal(t, addr.LINK_LOCAL_v4, response.Prefix)
		assert.Equal(t, addr.REGISTRY_IANA, response.Registry)
		assert.Equal(t, addr.DEFAULT_ALLOCATED_DATE, response.Allocated)
	})
	t.Run("ip4 global should query", func(t *testing.T) {
		t.Parallel()
		v, _ := addr.NewIPValidator("1.1.1.1")
		shouldQuery, response := v.Validate()
		assert.True(t, shouldQuery)
		assert.Nil(t, response)
	})
	t.Run("ip6 non-global has fallback response", func(t *testing.T) {
		t.Parallel()
		ipStr := "2001:db8::1"
		ip := net.ParseIP(ipStr)
		v, _ := addr.NewIPValidator(ipStr)
		shouldQuery, response := v.Validate()
		assert.False(t, shouldQuery)
		assert.IsType(t, &addr.Response{}, response)
		assert.Equal(t, uint64(0), response.ASN)
		assert.Equal(t, addr.TXT_DOC, response.Name)
		assert.Equal(t, countries.USA, response.Country)
		assert.Equal(t, &ip, response.IP)
		assert.Equal(t, addr.DOC_v6, response.Prefix)
		assert.Equal(t, addr.REGISTRY_IANA, response.Registry)
		assert.Equal(t, addr.DEFAULT_ALLOCATED_DATE, response.Allocated)
	})
	t.Run("ip6 global should query", func(t *testing.T) {
		t.Parallel()
		v, _ := addr.NewIPValidator("2606:4700:4700::1111")
		shouldQuery, response := v.Validate()
		assert.True(t, shouldQuery)
		assert.Nil(t, response)
	})
}
