package main

import (
	"fmt"
	"os"
	"path"
	"strconv"

	"github.com/getmillipede/millipede-go/vendor/github.com/codegangsta/cli"

	"github.com/getmillipede/millipede-go"
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
		cli.BoolFlag{
			Name:  "chameleon",
			Usage: "the millipede use its environment color",
		},
		cli.BoolFlag{
			Name:  "rainbow",
			Usage: "the millipede live with care bears",
		},
		cli.BoolFlag{
			Name:  "zalgo",
			Usage: "invoke the hive-mind representing chaos",
		},
		cli.BoolFlag{
			Name:  "animate",
			Usage: "he is alive !",
		},
		cli.StringFlag{
			Name:  "skin, template, s, t",
			Usage: "millipede skin (default, frozen, love, corporate, musician, bocal, ascii, inception, humancentipede, finger)",
			Value: "default",
		},
		cli.IntFlag{
			Name:  "width, w",
			Usage: "millipede width",
			Value: 3,
		},
		cli.IntFlag{
			Name:  "curve, c",
			Usage: "millipede curve size",
			Value: 4,
		},
		cli.IntFlag{
			Name:  "steps",
			Usage: "amount of steps done by the millipede",
			Value: 0,
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
			Size:      size,
			Reverse:   c.Bool("reverse"),
			Skin:      c.String("skin"),
			Opposite:  c.Bool("opposite"),
			Width:     uint64(c.Int("width")),
			Curve:     uint64(c.Int("curve")),
			Chameleon: c.Bool("chameleon"),
			Rainbow:   c.Bool("rainbow"),
			Zalgo:     c.Bool("zalgo"),
			Steps:     uint64(c.Int("steps")),
		}
		if c.Bool("animate") {
			millipede.Animate()
		} else {
			fmt.Println(millipede)
		}
	}
	app.Run(os.Args)
}
