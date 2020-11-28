package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	pcu "github.com/xjh22222228/python-check-updates"
	"github.com/xjh22222228/python-check-updates/constants"
	"log"
	"os"
	"time"
)

func main()  {
	app := &cli.App{
		Name: "python-check-updates",
		Usage: "Find the latest version of your requirements.txt current dependency package",
		Version: constants.Version,

		Flags: []cli.Flag {
			&cli.BoolFlag{
				Name: "update",
				Value: true,
				Aliases: []string{"u"},
				Usage: "update pkg",
			},
			&cli.StringFlag{
				Name: "file",
				Value: "requirements.txt",
				Aliases: []string{"f"},
				Usage: "Specify the file name of the check dependency package",
			},
		},

		Action: func(c *cli.Context) error {
			depFileName := c.String("f")

			beginTimestamp := time.Now().Unix()
			pkgList := pcu.ReadRequirements(depFileName)
			pkgList = pcu.GetNewVersion(pkgList)

			// Print
			pcu.EchoPrintVersion(pkgList)

			useTotalTime := time.Now().Unix() - beginTimestamp
			fmt.Println("\nDone in ", useTotalTime, "s.")
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
