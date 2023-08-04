package addr

import (
	"fmt"
	"net"

	"github.com/miekg/dns"
)

var DNS_SERVER string = "1.1.1.1:53"

type ErrLookupFailure error
type ErrLookupAssertionFailure error

func NewErrLookupFailure(target string, code int) ErrLookupFailure {
	return fmt.Errorf("failed to query '%s', code %d", target, code)
}

func NewErrLookupAssertionFailure(rr dns.RR) ErrLookupAssertionFailure {
	return fmt.Errorf("failed to assert response for target '%s' as %s record", rr.Header().Name, dns.Type(rr.Header().Rrtype).String())
}

func DNSLookup[T dns.RR](target string, lookupType uint16) ([]T, error) {
	client := new(dns.Client)
	msg := new(dns.Msg)
	if target[len(target)-1] != '.' {
		target += "."
	}
	msg.SetQuestion(target, lookupType)
	msg.RecursionDesired = true
	res, _, err := client.Exchange(msg, DNS_SERVER)
	if err != nil {
		return nil, err
	}
	if res.Rcode != dns.RcodeSuccess {
		return nil, NewErrLookupFailure(target, res.Rcode)
	}
	answers := []T{}
	for _, a := range res.Answer {
		rec, ok := a.(T)
		if !ok {
			return nil, NewErrLookupAssertionFailure(a)
		}
		answers = append(answers, rec)
	}
	return answers, nil
}

func DNSForwardLookup(host string) ([]net.IP, []net.IP, error) {
	host = dns.Fqdn(host)
	answersA, err := DNSLookup[*dns.A](host, dns.TypeA)
	if err != nil {
		return nil, nil, err
	}
	a := make([]net.IP, 0, len(answersA))
	for _, rec := range answersA {
		if rec != nil {
			a = append(a, rec.A)
		}
	}
	answersAAAA, err := DNSLookup[*dns.AAAA](host, dns.TypeAAAA)
	if err != nil {
		return nil, nil, err
	}
	aaaa := make([]net.IP, 0, len(answersAAAA))
	for _, rec := range answersAAAA {
		if rec != nil {
			aaaa = append(aaaa, rec.AAAA)
		}
	}
	return a, aaaa, nil
}

func DNSReverseLookup(ip *net.IP) ([]string, error) {
	arpa, err := dns.ReverseAddr(ip.String())
	if err != nil {
		return nil, err
	}
	answers, err := DNSLookup[*dns.PTR](arpa, dns.TypePTR)
	if err != nil {
		return nil, err
	}
	results := make([]string, len(answers)-1)
	for _, a := range answers {
		results = append(results, a.Ptr)
	}
	return results, nil
}
