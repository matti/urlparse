package main

import (
	"flag"
	"fmt"
	"net/url"
	"strings"
)

func main() {
	flag.Parse()

	urlString := flag.Arg(0)
	if u, err := url.Parse(urlString); err != nil {
		panic(err)
	} else {
		if u.Scheme == "" {
			u.Scheme = "http"
			if u, err = url.Parse(u.String()); err != nil {
				panic(err)
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

		parts := []string{
			"scheme=" + u.Scheme,
			"user=" + u.User.Username(),
			"password=" + password,
			"host=" + u.Host,
			"port=" + port,
			"path=" + u.Path,
			"fragment=" + u.Fragment,
			"query=" + u.RawQuery,
		}

		fmt.Println(strings.Join(parts, "\n"))
	}

}
