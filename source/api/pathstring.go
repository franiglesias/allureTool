package api

import "strings"

type PathString string

func (p PathString) toString() string {
	return string(p)
}

func (p PathString) WithBackslash() PathString {
	if !strings.HasSuffix(p.toString(), "/") {
		p += "/"
	}

	return p
}

func (p PathString) WithoutBackslash() PathString {
	if strings.HasSuffix(p.toString(), "/") {
		return PathString(strings.TrimSuffix(p.toString(), "/"))
	}

	return p
}

func (p PathString) WithSchema() PathString {
	if !strings.HasPrefix(p.toString(), "https://") && !strings.HasPrefix(p.toString(), "http://") {
		return "https://" + p
	}

	return p
}
