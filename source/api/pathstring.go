package api

import "strings"

type PathString string

func (p PathString) toString() string {
	return string(p)
}

func (p PathString) WithTrailingSlash() PathString {
	if !strings.HasSuffix(p.toString(), "/") {
		p += "/"
	}

	return p
}

func (p PathString) WithoutTrailingSlash() PathString {
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

func (p PathString) WithoutSlashes() PathString {
	s := p.WithoutTrailingSlash()

	if strings.HasPrefix(s.toString(), "/") {
		return PathString(strings.TrimPrefix(s.toString(), "/"))
	}

	return s
}
