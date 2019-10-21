package main

import (
	"fmt"
	"github.com/frolovo22/tag"
	"github.com/urfave/cli"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "tag"
	app.Usage = "mp3, mp4, flac tags reader/writer"

	app.Commands = cli.Commands{
		cli.Command{
			Name:      "read",
			ShortName: "r",
			Usage:     "read metadata from file",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:      "path",
					Usage:     "path to file",
					EnvVar:    "",
					FilePath:  "",
					Required:  false,
					Hidden:    false,
					TakesFile: true,
					Value:     "",
				},
			},
			Action: func(c *cli.Context) error {
				path := c.Args().First()
				metadata, err := tag.ReadFile(path)
				if err != nil {
					return err
				}
				tags := tag.GetMap(metadata)
				for key, val := range tags {
					fmt.Println(fmt.Sprintf("%-20s: %v", key, val))
				}
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
