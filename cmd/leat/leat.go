package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/nodar-chkuaselidze/go-leat/leats"
	"gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Name = "go-leat"
	app.Usage = "Golang Length Extention Attacks"
	app.Version = "0.1.0"
	app.HideVersion = true
	app.Commands = commands()

	app.Run(os.Args)
}

func commands() cli.Commands {
	commands := make(cli.Commands, len(leats.LeatList))

	for i, leat := range leats.LeatList {
		commands[i] = cli.Command{
			Action: func(ctx *cli.Context) error {
				args := ctx.Args()
				hash := args.Get(1)
				extend := args.Get(2)
				fprint := args.Get(3)
				length, err := strconv.ParseInt(args.Get(0), 10, 32)

				if err != nil {
					return errors.New("Couldn't parse length")
				}

				hash, extendWith, err := leat.Fn(int(length), hash, extend, fprint)

				if err != nil {
					return err
				}

				fmt.Println(hash)
				fmt.Println(extendWith)

				return nil
			},
			Name:        leat.Name,
			Usage:       "length oldhash extendwith fprintByte",
			Description: "Extend " + leat.Name + " hash",
		}
	}

	return commands
}
