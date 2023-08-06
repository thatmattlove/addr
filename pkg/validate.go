package addr

import (
	"net"

	"github.com/biter777/countries"
	goasn "github.com/thatmattlove/go-asn"
)

type IPValidator struct {
	InitialValue string
	IP           net.IP
	Net          *net.IPNet
}

func IsIPv6(ip net.IP) bool {
	return ip.To4() == nil
}

func GetNonGlobalPrefix(ip net.IP) (*net.IPNet, string) {
	if IsIPv6(ip) {
		switch true {
		case MULTICAST_v6.Contains(ip):
			return MULTICAST_v6, TXT_MULTICAST
		case LINK_LOCAL_v6.Contains(ip):
			return LINK_LOCAL_v6, TXT_LINK_LOCAL
		case UNIQUE_LOCAL.Contains(ip):
			return UNIQUE_LOCAL, TXT_ULA
		case LOOPBACK_v6.IP.Equal(ip):
			return LOOPBACK_v6, TXT_LOOPBACK
		case DEFAULT_v6.IP.Equal(ip):
			return DEFAULT_v6, TXT_DEFAULT
		case DOC_v6.Contains(ip):
			return DOC_v6, TXT_DOC
		case EMBEDDED.Contains(ip):
			return EMBEDDED, TXT_EMBEDDED
		case ORCHIDv1.Contains(ip):
			return ORCHIDv1, TXT_ORCHIDv1
		case ORCHIDv2.Contains(ip):
			return ORCHIDv2, TXT_ORCHIDv2
		case DETS.Contains(ip):
			return DETS, TXT_DETS
		case TEREDO.Contains(ip):
			return TEREDO, TXT_TEREDO
		case DISCARD.Contains(ip):
			return DISCARD, TXT_DISCARD
		case BENCHMARK_v6.Contains(ip):
			return BENCHMARK_v6, TXT_BENCHMARK
		case SIX_TO_FOUR_v6.Contains(ip):
			return SIX_TO_FOUR_v6, TXT_6to4
		case AS112_v6.Contains(ip):
			return AS112_v6, TXT_AS112
		case AS112_v6_Direct.Contains(ip):
			return AS112_v6_Direct, TXT_AS112
		case AMT_v6.Contains(ip):
			return AMT_v6, TXT_AMT
		default:
			return nil, ""
		}
	}
	switch true {
	case LINK_LOCAL_v4.Contains(ip):
		return LINK_LOCAL_v4, TXT_LINK_LOCAL
	case RFC1918_10.Contains(ip):
		return RFC1918_10, TXT_PRIVATE
	case RFC1918_172.Contains(ip):
		return RFC1918_172, TXT_PRIVATE
	case RFC1918_192.Contains(ip):
		return RFC1918_192, TXT_PRIVATE
	case CGNAT.Contains(ip):
		return CGNAT, TXT_CGNAT
	case RFC6890_192.Contains(ip):
		return RFC6890_192, TXT_PRIVATE
	case AS112_v4.Contains(ip):
		return AS112_v4, TXT_AS112
	case AS112_v4_Direct.Contains(ip):
		return AS112_v4_Direct, TXT_AS112
	case AMT_v4.Contains(ip):
		return AMT_v4, TXT_AMT
	case SIX_TO_FOUR_v4.Contains(ip):
		return SIX_TO_FOUR_v4, TXT_6to4
	case BENCHMARK_v4.Contains(ip):
		return BENCHMARK_v4, TXT_BENCHMARK
	case DOC_1.Contains(ip):
		return DOC_1, TXT_DOC
	case DOC_2.Contains(ip):
		return DOC_2, TXT_DOC
	case DOC_3.Contains(ip):
		return DOC_3, TXT_DOC
	case RESERVED.Contains(ip):
		return RESERVED, TXT_RESERVED
	case THIS_NETWORK.Contains(ip):
		return THIS_NETWORK, TXT_THIS
	case MULTICAST_v4.Contains(ip):
		return MULTICAST_v4, TXT_MULTICAST
	case LOOPBACK_v4.Contains(ip):
		return LOOPBACK_v4, TXT_LOOPBACK
	case DEFAULT_v4.IP.Equal(ip):
		return DEFAULT_v4, TXT_DEFAULT
	default:
		return nil, ""
	}
}

func (ipv *IPValidator) Validate() (bool, *Response) {
	pfx, txt := GetNonGlobalPrefix(ipv.IP)
	if pfx == nil {
		return true, nil
	}
	response := &Response{
		ASN:       goasn.ASN{0, 0, 0, 0},
		IP:        &ipv.IP,
		Prefix:    pfx,
		Name:      txt,
		Country:   countries.USA,
		Allocated: DEFAULT_ALLOCATED_DATE,
		Registry:  REGISTRY_IANA,
	}
	return false, response
}

func NewIPValidator(in string) (ipv *IPValidator, err error) {
	var prefix *net.IPNet
	ip := net.ParseIP(in)
	if ip == nil {
		ip, prefix, err = net.ParseCIDR(in)
		if err != nil {
			return nil, err
		}
	}
	ipv = &IPValidator{
		InitialValue: in,
		IP:           ip,
		Net:          prefix,
	}
	return ipv, nil
}
