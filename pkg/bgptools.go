package addr

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/biter777/countries"
	"github.com/thatmattlove/addr/pkg/whois"
	goasn "github.com/thatmattlove/go-asn"
)

type Response struct {
	ASN       goasn.ASN
	IP        *net.IP
	Prefix    *net.IPNet
	Country   countries.CountryCode
	Registry  string
	Allocated time.Time
	Name      string
	FromQuery bool
}

var (
	WHOIS_HOST string = "bgp.tools"
	WHOIS_PORT uint   = 43
)

var ErrEmptyResponse = errors.New("empty response")

func ParseResponse(res string) (*Response, error) {
	scanner := bufio.NewScanner(strings.NewReader(res))
	scanner.Scan()

	// Get first line, which will either be a warning or the headers.
	line := scanner.Text()
	if line == "" {
		return nil, ErrEmptyResponse
	}
	if strings.Contains(line, "Warning") {
		// Skip ahead to the header row if the first row is a warning.
		scanner.Scan()
	}

	values := []string{}
	parts := []string{}

	var next = func() error {
		scanner.Scan()
		line = scanner.Text()
		parts = strings.Split(line, "|")
		for _, p := range parts {
			values = append(values, strings.TrimSpace(p))
		}
		if line == "" {
			return ErrEmptyResponse
		}
		return nil
	}

	// Move from the header row to the first data row, which is the only row we care about.
	err := next()
	if err != nil {
		return nil, err
	}

	if len(values) != 7 {
		var err error
		err = next()
		if err != nil {
			return nil, err
		}
		if len(values) != 7 {
			err = fmt.Errorf("expected 7 columns, got %d", len(values))
			return nil, err
		}
	}

	asnStr := values[0]       // 'AS' column.
	ipStr := values[1]        // 'IP' column.
	pfxStr := values[2]       // 'BGP Prefix' column.
	countryStr := values[3]   // 'CC' column.
	registry := values[4]     // 'Registry' column.
	allocatedStr := values[5] // 'Allocated' column.
	name := values[6]         // 'AS Name' column.

	asn, err := goasn.Parse(asnStr)
	if err != nil {
		return nil, err
	}

	allocated, err := time.Parse(time.DateOnly, allocatedStr)
	if err != nil {
		return nil, err
	}

	country := countries.ByName(countryStr)

	response := &Response{
		ASN:       asn,
		IP:        nil,
		Prefix:    nil,
		Country:   country,
		Registry:  registry,
		Allocated: allocated,
		Name:      name,
		FromQuery: true,
	}

	// Parse IP & BGP Prefix if it was returned.
	if ipStr != "" && pfxStr != "" {
		ip := net.ParseIP(ipStr)
		if ip != nil {
			response.IP = &ip
		} else {
			err = fmt.Errorf("failed to parse IP '%s'", ipStr)
			return nil, err
		}
		_, pfx, err := net.ParseCIDR(pfxStr)
		if err != nil {
			return nil, err
		}
		response.Prefix = pfx
	}
	return response, nil
}

func QueryASN(asnStr string) (*Response, error) {
	asn, err := goasn.Parse(asnStr)
	if err != nil {
		return nil, err
	}
	w, err := whois.New(WHOIS_HOST, WHOIS_PORT)
	if err != nil {
		return nil, err
	}
	result, err := w.Query(fmt.Sprintf("as%s", asn.ASPlain()))
	if err != nil {
		return nil, err
	}
	res, err := ParseResponse(result)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func QueryIPPrefix(q string) (*Response, error) {
	validator, err := NewIPValidator(q)
	if err != nil {
		return nil, err
	}
	shouldQuery, res := validator.Validate()
	if !shouldQuery && res != nil {
		return res, nil
	}
	w, err := whois.New(WHOIS_HOST, WHOIS_PORT)
	if err != nil {
		return nil, err
	}
	result, err := w.Query(q)
	if err != nil {
		return nil, err
	}
	res, err = ParseResponse(result)
	if err != nil {
		return nil, err
	}
	return res, nil
}
