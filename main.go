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
		password := ""
		if p, ok := u.User.Password(); ok {
			password = p
		}

		parts := []string{
			u.Scheme,
			u.User.Username(),
			password,
			u.Host,
			u.Port(),
			u.Path,
			u.Fragment,
			u.RawQuery,
		}

		fmt.Println(strings.Join(parts, "\n"))
	}

}
