/*
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
	"math/rand"
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
		cli.BoolFlag{
			Name:  "random",
			Usage: "RANDOMZ!",
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

		millipede := millipede.NewWithSize(size)

		if c.Bool("random") {
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			millipede.Reverse = r.Intn(2) != 0
			millipede.Opposite = r.Intn(2) != 0
			millipede.Chameleon = r.Intn(2) != 0
			millipede.Rainbow = r.Intn(2) != 0
			millipede.Zalgo = r.Intn(5) == 0
			millipede.Center = r.Intn(2) != 0
			millipede.Size = uint64(r.Intn(50) + 15)
			millipede.Curve = uint64(r.Intn(7) + 2)
			millipede.Width = uint64(r.Intn(10) + 3)
		} else {
			millipede.Reverse = c.Bool("reverse")
			millipede.Opposite = c.Bool("opposite")
			millipede.Chameleon = c.Bool("chameleon")
			millipede.Rainbow = c.Bool("rainbow")
			millipede.Zalgo = c.Bool("zalgo")
			millipede.Center = c.Bool("center")
			millipede.Skin = c.String("skin")
			millipede.Width = uint64(c.Int("width"))
			millipede.Curve = uint64(c.Int("curve"))
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
