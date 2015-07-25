package main

import (
	"fmt"
	"os"
	"path"
	"strconv"

	"github.com/codegangsta/cli"

	"github.com/getmillipede/millipede-go/millipede"
	"github.com/getmillipede/millipede-go/version"
)

func main() {
	app := cli.NewApp()
	app.Name = path.Base(os.Args[0])
	app.Author = "Millipede crew"
	app.Email = "https://github.com/getmillipede/millipede-go"
	app.Version = version.VERSION + " (" + version.GITCOMMIT + ")"
	app.Usage = "Print a beautiful millipede"

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "reverse, r",
			Usage: "reverse the millipede",
		},
		cli.BoolFlag{
			Name:  "opposite, o",
			Usage: "go the opposite direction",
		},
		cli.StringFlag{
			Name:  "skin, template, s, t",
			Usage: "millipede skin (default)",
			Value: "default",
		},
	}

	app.Action = func(c *cli.Context) {
		var size uint64
		var err error
		size = 20
		if len(c.Args()) > 0 {
			size, err = strconv.ParseUint(c.Args()[0], 10, 0)
			if err != nil {
				panic(err)
			}
		}

		millipede := &millipede.Millipede{
			Size:     size,
			Reverse:  c.Bool("reverse"),
			Skin:     c.String("skin"),
			Opposite: c.Bool("opposite"),
		}
		fmt.Println(millipede)
	}
	app.Run(os.Args)
}
