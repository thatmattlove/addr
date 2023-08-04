package style

import (
	"fmt"
	"strings"

	addr "github.com/thatmattlove/addr/pkg"
)

func IPBox(r *addr.Response, ptrs []string) string {
	netPrefix := "from "
	var asn string
	if r.FromQuery {
		netPrefix = "advertised as "
		asn = Plain("AS") + Highlight2(fmt.Sprint(r.ASN))
	} else {
		asn = Subtle("Never Advertised")
	}
	net := Subtle(netPrefix) + Highlight1(r.Prefix.String())
	org := Country(r)
	reg := Subtle("Registry: ") + Plain(r.Registry)
	body := strings.Join([]string{net, asn, org, reg}, "\n")
	title := Title(r.IP.String())
	box := Box.WithTitle(title)
	if len(ptrs) > 0 {
		for _, p := range ptrs {
			body = fmt.Sprintf("%s\n\n%s", Subtle(strings.Trim(p, ".")), body)
			box = box.WithTopPadding(0)
		}
	}
	return Wrapper.Sprint(box.Sprint(body))
}

func ASNBox(r *addr.Response) string {
	asn := Plain("AS") + Title(fmt.Sprint(r.ASN))
	org := Country(r)
	return Wrapper.Sprint(
		Box.WithTitle(asn).Sprint(org),
	)
}
