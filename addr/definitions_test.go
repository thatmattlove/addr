package addr_test

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thatmattlove/addr/addr"
)

var cases = map[*net.IPNet]string{
	addr.LINK_LOCAL_v4:   "169.254.0.0/16",
	addr.LINK_LOCAL_v6:   "fe80::/10",
	addr.RFC1918_10:      "10.0.0.0/8",
	addr.RFC1918_172:     "172.16.0.0/12",
	addr.RFC1918_192:     "192.168.0.0/16",
	addr.CGNAT:           "100.64.0.0/10",
	addr.RFC6890_192:     "192.0.0.0/24",
	addr.AS112_v4:        "192.31.196.0/24",
	addr.AS112_v4_Direct: "192.175.48.0/24",
	addr.AS112_v6:        "2001:4:112::/48",
	addr.AS112_v6_Direct: "2620:4f:8000::/48",
	addr.AMT_v4:          "192.52.193.0/24",
	addr.AMT_v6:          "2001:3::/32",
	addr.SIX_TO_FOUR_v4:  "192.88.99.0/24",
	addr.SIX_TO_FOUR_v6:  "2002::/16",
	addr.BENCHMARK_v4:    "198.18.0.0/15",
	addr.BENCHMARK_v6:    "2001:2::/48",
	addr.DOC_1:           "192.0.2.0/24",
	addr.DOC_2:           "198.51.100.0/24",
	addr.DOC_3:           "203.0.113.0/24",
	addr.RESERVED:        "240.0.0.0/4",
	addr.THIS_NETWORK:    "0.0.0.0/8",
	addr.MULTICAST_v4:    "224.0.0.0/4",
	addr.MULTICAST_v6:    "ff00::/8",
	addr.LOOPBACK_v4:     "127.0.0.0/8",
	addr.LOOPBACK_v6:     "::1/128",
	addr.DEFAULT_v4:      "0.0.0.0/0",
	addr.DEFAULT_v6:      "::/0",
	addr.UNIQUE_LOCAL:    "fc00::/7",
	addr.DISCARD:         "100::/64",
	addr.TEREDO:          "2001::/32",
	addr.ORCHIDv1:        "2001:10::/28",
	addr.ORCHIDv2:        "2001:20::/28",
	addr.DOC_v6:          "2001:db8::/32",
	addr.EMBEDDED:        "64:ff9b::/96",
	addr.DETS:            "2001:30::/28",
}

func Test_Definitions(t *testing.T) {
	for pfx, expected := range cases {
		pfx := pfx
		expected := expected
		t.Run(expected, func(t *testing.T) {
			t.Parallel()
			result := pfx.String()
			_, exp, _ := net.ParseCIDR(expected)
			e := bytes.Clone(exp.IP)
			p := bytes.Clone(pfx.IP)
			es := hex.EncodeToString(e)
			ps := hex.EncodeToString(p)
			assert.Equal(t, expected, result, "expected bytes: %s, got %s", es, ps)
		})
	}
	for pfx, expected := range cases {
		pfx := pfx
		expected := expected
		t.Run(fmt.Sprintf("%s should fail", expected), func(t *testing.T) {
			t.Parallel()
			result := pfx.String()
			assert.NotEqual(t, "1.1.1.0/24", result)
		})
	}
}
