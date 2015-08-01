/* millipede-go binary - Print a beautiful millipede from a command line interface

$ millipede-go -h
NAME:
   millipede-go - Print a beautiful millipede

USAGE:
   millipede-go [global options] command [command options] [arguments...]

VERSION:
   1.3.0-dev (HEAD)

AUTHOR(S):
   Millipede crew <https://github.com/getmillipede/millipede-go>

COMMANDS:
   help, hShows a list of commands or help for one command

GLOBAL OPTIONS:
   --reverse, -rreverse the millipede
   --opposite, -ogo the opposite direction
   --chameleonthe millipede use its environment color
   --rainbowthe millipede live with care bears
   --zalgoinvoke the hive-mind representing chaos
   --animatehe is alive !
   --center, -Cmillipede in the midle
   --skin, --template, -s, -t "default"millipede skin (default, frozen, love, corporate, musician, bocal, ascii, inception, humancentipede, finger)
   --width, -w "3"millipede width
   --curve, -c "4"millipede curve size
   --steps "0"amount of steps done by the millipede
   --help, -hshow help
   --version, -vprint the version
*/
package main

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"time"

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
		cli.BoolFlag{
			Name:  "center, C",
			Usage: "millipede in the midle",
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
			Center:    c.Bool("center"),
		}
		if c.Bool("animate") {
			millipede.PadRight = true
			fmt.Println("\033[2J")
			for steps := 0; ; steps++ {
				frame := fmt.Sprint(millipede)
				fmt.Println("\033[0;0H")
				fmt.Println(frame)
				millipede.Steps++
				time.Sleep(100 * time.Millisecond)
			}
		} else {
			fmt.Println(millipede)
		}
	}
	app.Run(os.Args)
}
