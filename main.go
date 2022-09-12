package urlparse

import (
	"errors"
	"net/url"
	"strings"
)

type URL struct {
	Scheme   string
	Password string
	Username string
	Host     string
	Port     string
	Path     string
	Query    string
	Fragment string
}

func (u2 *URL) String() string {
	var parts []string

	parts = append(parts, u2.Scheme)
	parts = append(parts, "://")

	if u2.Username != "" {
		parts = append(parts, u2.Username)
		if u2.Password != "" {
			parts = append(parts, ":"+u2.Password)
		}
		parts = append(parts, "@")
	}

	parts = append(parts, u2.Host)

	switch u2.Scheme {
	case "http":
		if u2.Port != "80" {
			parts = append(parts, ":"+u2.Port)
		}
	case "https":
		if u2.Port != "443" {
			parts = append(parts, ":"+u2.Port)
		}
	}

	if u2.Path != "" {
		parts = append(parts, u2.Path)
	}

	if u2.Query != "" {
		parts = append(parts, "?"+u2.Query)
	}

	if u2.Fragment != "" {
		parts = append(parts, "#"+u2.Fragment)
	}

	return strings.Join(parts, "")
}

func Parse(urlString string) (*URL, error) {
	urlString = strings.TrimSpace(urlString)
	if urlString == "" {
		return nil, errors.New("empty URL")
	}

	if parts := strings.Split(urlString, "@"); len(parts) > 1 {
		beforeAt := parts[0]
		switch len(strings.Split(beforeAt, ":")) {
		case 1:
			return nil, errors.New("URL contains user but no password")
		case 2:
			urlString = "http://" + urlString
		}
	}

	var u *url.URL
	var err error
	if u, err = url.Parse(urlString); err != nil {
		return nil, err
	}

	if u.Scheme == "" {
		u.Scheme = "http"
		if u, err = url.Parse(u.String()); err != nil {
			return nil, err
		}
	}

	password := ""
	if p, ok := u.User.Password(); ok {
		password = p
	}

	port := u.Port()
	if port == "" {
		switch u.Scheme {
		case "http":
			port = "80"
		case "https":
			port = "443"
		}
	}

	u2 := &URL{
		Scheme:   u.Scheme,
		Password: password,
		Username: u.User.Username(),
		Host:     u.Hostname(),
		Port:     port,
		Path:     u.Path,
		Query:    u.RawQuery,
		Fragment: u.Fragment,
	}

	return u2, err
}
