// Package swot finds academic domains and emails
//go:generate broccoli -src domains -o domains
package swot

import (
	"errors"
	"io/ioutil"
	"net/url"
	"path/filepath"
	"strings"

	valid "github.com/asaskevich/govalidator"
)

var (
	// ErrDomainNotFound happens when a domain is not found
	ErrDomainNotFound = errors.New("Domain name not found")

	// ErrSchoolNotFound happens whan a school name is not found
	ErrSchoolNotFound = errors.New("School name not found")
)

func isBlacklisted(domain string) bool {
	for _, blacklisted := range blacklist {
		if strings.HasSuffix(domain, blacklisted) {
			return true
		}
	}
	return false
}

func isAcademicTLD(domain string) bool {
	for _, tld := range tlds {
		if strings.HasSuffix(domain, tld) {
			return true
		}
	}
	return false
}

func parseDomain(address string) (string, error) {
	address = strings.ToLower(strings.TrimSpace(address))
	switch {
	case valid.IsEmail(address):
		return strings.Split(address, "@")[1], nil
	case valid.IsURL(address):
		if valid.IsRequestURL(address) {
			url, err := url.Parse(address)
			if err != nil {
				return "", err
			}
			return strings.Split(url.Host, ":")[0], nil
		}

		return address, nil
	}
	return "", ErrDomainNotFound
}

func fileExists(path string) bool {
	if _, err := br.Stat(path); err == nil {
		return true
	}
	return false
}

func getInstitutionName(address string) (string, error) {
	domain, err := parseDomain(address)
	if err != nil {
		return "", err
	}

	domainParts := splitdomain(domain)
	path := "domains"
	for i := len(domainParts) - 1; i >= 0; i-- {
		path = filepath.Join(path, domainParts[i])
		if fileExists(path + ".txt") {
			f, err := br.Open(path + ".txt")
			if err != nil {
				return "", err
			}

			b, err := ioutil.ReadAll(f)
			if err != nil {
				return "", err
			}

			return string(b), nil
		}
	}

	return "", ErrSchoolNotFound
}

func splitdomain(domain string) []string {
	return strings.Split(domain, ".")
}

// IsAcademic returns true if the email address or URL belongs
// to an academic institution.
func IsAcademic(address string) bool {
	domain, err := parseDomain(address)
	if err != nil {
		return false
	}

	if isBlacklisted(domain) {
		return false
	} else if isAcademicTLD(domain) {
		return true
	}

	_, err = getInstitutionName(domain)
	return err == nil
}

// GetSchoolName returns the name of the academic institution or
// an empty string if the name of the institution is not found.
func GetSchoolName(address string) string {
	s, err := getInstitutionName(address)
	if err != nil {
		return ""
	}

	return strings.TrimSpace(s)
}
