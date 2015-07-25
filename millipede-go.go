package main

import (
	"os"
	"path"
	"strconv"

	"github.com/codegangsta/cli"

	"github.com/tehmoon/millipede-go/millipede"
	"github.com/tehmoon/millipede-go/version"
)

func main() {
	app := cli.NewApp()
	app.Name = path.Base(os.Args[0])
	app.Author = "Millipede crew"
	app.Email = "https://github.com/getmillipede/millipede-go"
	app.Version = version.VERSION + " (" + version.GITCOMMIT + ")"
	app.Usage = "Print a beautiful millipede"

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

		m := millipede.New(size)

		// or fmt.Println(m.Print)
		// or m.Reverse().Println()
		m.Reverse().Reverse().Println()
	}
	app.Run(os.Args)
}
