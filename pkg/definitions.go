package addr

import (
	"net"
	"time"
)

const (
	TXT_PRIVATE    string = "Private Use"
	TXT_MULTICAST  string = "Multicast"
	TXT_LINK_LOCAL string = "Link Local"
	TXT_ULA        string = "Unique Local"
	TXT_LOOPBACK   string = "Loopback"
	TXT_DEFAULT    string = "Unspecified/Default"
	TXT_EMBEDDED   string = "Embedded IPv4-IPv6 Translation (RFC6052)"
	TXT_ORCHIDv1   string = "ORCHIDv1 (Deprecated)"
	TXT_ORCHIDv2   string = "ORCHIDv2 (RFC7343)"
	TXT_DOC        string = "Documentation"
	TXT_RESERVED   string = "Reserved"
	TXT_THIS       string = "This Network"
	TXT_AS112      string = "AS112"
	TXT_6to4       string = "6to4 Relay"
	TXT_DETS       string = "Drone Remote ID Protocol Entity Tags"
	TXT_TEREDO     string = "TEREDO (RFC4380, RFC8190)"
	TXT_DISCARD    string = "Discard-Only Block"
	TXT_BENCHMARK  string = "Benchmarking (RFC5180)"
	TXT_AMT        string = "Automatic Multicast Tunneling (RFC7450)"
	TXT_CGNAT      string = "Shared Address Space/Carrier-Grade NAT"
)

const (
	REGISTRY_IANA string = "IANA"
)

var (
	IPv4Bits                         = net.IPv4len * 8
	IPv6Bits                         = net.IPv6len * 8
	DEFAULT_ALLOCATED_DATE time.Time = time.Date(1981, 9, 1, 0, 0, 0, 0, time.UTC)
)

var (
	// 169.254.0.0/16
	LINK_LOCAL_v4 = &net.IPNet{
		IP:   net.IPv4(169, 254, 0, 0),
		Mask: net.CIDRMask(16, IPv4Bits),
	}
	// fe80::/10
	LINK_LOCAL_v6 = &net.IPNet{
		IP:   net.IP{0xfe, 0x80, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		Mask: net.CIDRMask(10, net.IPv6len*8),
	}
	// 10.0.0.0/8
	RFC1918_10 = &net.IPNet{
		IP:   net.IPv4(10, 0, 0, 0),
		Mask: net.CIDRMask(8, IPv4Bits),
	}
	// 172.16.0.0/12
	RFC1918_172 = &net.IPNet{
		IP:   net.IPv4(172, 16, 0, 0),
		Mask: net.CIDRMask(12, IPv4Bits),
	}
	// 192.168.0.0/16
	RFC1918_192 = &net.IPNet{
		IP:   net.IPv4(192, 168, 0, 0),
		Mask: net.CIDRMask(16, IPv4Bits),
	}
	// 100.64.0.0/10
	CGNAT = &net.IPNet{
		IP:   net.IPv4(100, 64, 0, 0),
		Mask: net.CIDRMask(10, IPv4Bits),
	}
	// 192.0.0.0/24
	RFC6890_192 = &net.IPNet{
		IP:   net.IPv4(192, 0, 0, 0),
		Mask: net.CIDRMask(24, IPv4Bits),
	}
	// 192.31.196.0/24
	AS112_v4 = &net.IPNet{
		IP:   net.IPv4(192, 31, 196, 0),
		Mask: net.CIDRMask(24, IPv4Bits),
	}
	// 192.175.48.0/24
	AS112_v4_Direct = &net.IPNet{
		IP:   net.IPv4(192, 175, 48, 0),
		Mask: net.CIDRMask(24, IPv4Bits),
	}
	// 2001:4:112::/48
	AS112_v6 = &net.IPNet{
		IP:   net.IP{0x20, 0x01, 0x00, 0x04, 0x01, 0x12, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		Mask: net.CIDRMask(48, IPv6Bits),
	}
	// 2620:4f:8000::/48
	AS112_v6_Direct = &net.IPNet{
		IP:   net.IP{0x26, 0x20, 0x00, 0x4f, 0x80, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		Mask: net.CIDRMask(48, IPv6Bits),
	}
	// 192.52.193.0/24
	AMT_v4 = &net.IPNet{
		IP:   net.IPv4(192, 52, 193, 0),
		Mask: net.CIDRMask(24, IPv4Bits),
	}
	// 2001:3::/32
	AMT_v6 = &net.IPNet{
		IP:   net.IP{0x20, 0x01, 0x00, 0x03, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		Mask: net.CIDRMask(32, IPv6Bits),
	}
	// 192.88.99.0/24
	SIX_TO_FOUR_v4 = &net.IPNet{
		IP:   net.IPv4(192, 88, 99, 0),
		Mask: net.CIDRMask(24, IPv4Bits),
	}
	// 2002::/16
	SIX_TO_FOUR_v6 = &net.IPNet{
		IP:   net.IP{0x20, 0x02, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		Mask: net.CIDRMask(16, IPv6Bits),
	}
	// 198.18.0.0/15
	BENCHMARK_v4 = &net.IPNet{
		IP:   net.IPv4(198, 18, 0, 0),
		Mask: net.CIDRMask(15, IPv4Bits),
	}
	// 2001:2::/48
	BENCHMARK_v6 = &net.IPNet{
		IP:   net.IP{0x20, 0x01, 0x00, 0x02, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		Mask: net.CIDRMask(48, IPv6Bits),
	}
	// 192.0.2.0/24
	DOC_1 = &net.IPNet{
		IP:   net.IPv4(192, 0, 2, 0),
		Mask: net.CIDRMask(24, IPv4Bits),
	}
	// 198.51.100.0/24
	DOC_2 = &net.IPNet{
		IP:   net.IPv4(198, 51, 100, 0),
		Mask: net.CIDRMask(24, IPv4Bits),
	}
	// 203.0.113.0/24
	DOC_3 = &net.IPNet{
		IP:   net.IPv4(203, 0, 113, 0),
		Mask: net.CIDRMask(24, IPv4Bits),
	}
	// 240.0.0.0/4
	RESERVED = &net.IPNet{
		IP:   net.IPv4(240, 0, 0, 0),
		Mask: net.CIDRMask(4, IPv4Bits),
	}
	// 0.0.0.0/8
	THIS_NETWORK = &net.IPNet{
		IP:   net.IPv4(0, 0, 0, 0),
		Mask: net.CIDRMask(8, IPv4Bits),
	}
	// 224.0.0.0/4
	MULTICAST_v4 = &net.IPNet{
		IP:   net.IPv4(224, 0, 0, 0),
		Mask: net.CIDRMask(4, IPv4Bits),
	}
	// ff00::/8
	MULTICAST_v6 = &net.IPNet{
		IP:   net.IP{0xff, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		Mask: net.CIDRMask(8, IPv6Bits),
	}
	// 127.0.0.0/8
	LOOPBACK_v4 = &net.IPNet{
		IP:   net.IPv4(127, 0, 0, 0),
		Mask: net.CIDRMask(8, IPv4Bits),
	}
	// ::1/128
	LOOPBACK_v6 = &net.IPNet{
		IP:   net.IP{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		Mask: net.CIDRMask(IPv6Bits, IPv6Bits),
	}
	// 0.0.0.0/0
	DEFAULT_v4 = &net.IPNet{
		IP:   net.IPv4zero,
		Mask: net.CIDRMask(0, IPv4Bits),
	}
	// ::/0
	DEFAULT_v6 = &net.IPNet{
		IP:   net.IPv6zero,
		Mask: net.CIDRMask(0, IPv6Bits),
	}
	// fc00::/7
	UNIQUE_LOCAL = &net.IPNet{
		IP:   net.IP{0xfc, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		Mask: net.CIDRMask(7, IPv6Bits),
	}
	// 100::/64
	DISCARD = &net.IPNet{
		IP:   net.IP{0x01, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		Mask: net.CIDRMask(64, IPv6Bits),
	}
	// 2001::/32
	TEREDO = &net.IPNet{
		IP:   net.IP{0x20, 0x01, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		Mask: net.CIDRMask(32, IPv6Bits),
	}
	// 2001:10::/28
	ORCHIDv1 = &net.IPNet{
		IP:   net.IP{0x20, 0x01, 0x00, 0x10, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		Mask: net.CIDRMask(28, IPv6Bits),
	}
	// 2001:20::/28
	ORCHIDv2 = &net.IPNet{
		IP:   net.IP{0x20, 0x01, 0x00, 0x20, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		Mask: net.CIDRMask(28, IPv6Bits),
	}
	// 2001:30::/28
	DETS = &net.IPNet{
		IP:   net.IP{0x20, 0x01, 0x00, 0x30, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		Mask: net.CIDRMask(28, IPv6Bits),
	}
	// 2001:db8::/32
	DOC_v6 = &net.IPNet{
		IP:   net.IP{0x20, 0x01, 0x0d, 0xb8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		Mask: net.CIDRMask(32, IPv6Bits),
	}
	// 64:ff9b::/96
	EMBEDDED = &net.IPNet{
		IP:   net.IP{0x00, 0x64, 0xff, 0x9b, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		Mask: net.CIDRMask(96, IPv6Bits),
	}
)
