package g

import (
	"errors"
	"regexp"
	"strings"
)

type Parser interface {
	Parse(string) (*Address, error)
	Match(string) bool
}

var (
	ErrInvalidFormat = errors.New("invalid format")
)

var defaultParsers = []Parser{SSHParser, HTTPSParser}

func Parse(v string) (*Address, error) {
	for _, p := range defaultParsers {
		if p.Match(v) {
			return p.Parse(v)
		}
	}

	return nil, ErrInvalidFormat
}

var SSHParser = _SSHParser{}

type _SSHParser struct {
}

func (p _SSHParser) Parse(v string) (*Address, error) {
	if !p.Match(v) {
		return nil, ErrInvalidFormat
	}

	v = strings.Split(v, "@")[1]
	v = strings.TrimSuffix(v, ".git")
	vSplit := strings.Split(v, ":")

	server := vSplit[0]

	vSplit = strings.Split(vSplit[1], "/")
	ns, repo := vSplit[0], vSplit[1]

	return &Address{
		Server:     server,
		Namespace:  ns,
		Repository: repo,
	}, nil
}

func (_SSHParser) Match(v string) bool {
	f, e := regexp.Compile("git@[^/]+:[^/]+/[^/]+\\.git")
	if e != nil {
		return false
	}

	return f.Match([]byte(v))
}

var HTTPSParser = _HTTPSParser{}

type _HTTPSParser struct {
}

func (p _HTTPSParser) Match(s string) bool {
	f, e := regexp.Compile("http[s]?://[^/]+/[^/]+/[^/]+.git")
	if e != nil {
		return false
	}

	return f.Match([]byte(s))
}

func (p _HTTPSParser) Parse(v string) (*Address, error) {
	if !p.Match(v) {
		return nil, ErrInvalidFormat
	}

	v = strings.TrimPrefix(v, "https://")
	v = strings.TrimPrefix(v, "http://")
	v = strings.TrimSuffix(v, ".git")

	vSplit := strings.Split(v, "/")

	return &Address{
		Server:     vSplit[0],
		Namespace:  vSplit[1],
		Repository: vSplit[2],
	}, nil
}
