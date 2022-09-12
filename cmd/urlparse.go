package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gosimple/slug"
	"github.com/matti/urlparse"
)

func main() {
	flag.Parse()

	subcommand := flag.Arg(0)

	switch subcommand {
	case "slug":
		if u, err := urlparse.Parse(flag.Arg(1)); err != nil {
			panic(err)
		} else {
			fmt.Print(
				slug.Make(u.String()),
			)
		}
	case "parse":
		if u, err := urlparse.Parse(flag.Arg(1)); err != nil {
			panic(err)
		} else {
			fmt.Println(u.String())
		}
	default:
		println(
			fmt.Sprintf("unknown subcommand '%s'", subcommand),
		)
		os.Exit(1)
	}

}
