package main

import (
	"encoding/json"
	"fmt"
	"github.com/frolovo22/tag"
	"github.com/urfave/cli"
	"log"
	"os"
	"path/filepath"
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
					Name:  "input, in",
					Usage: "path to input file",
				},
				cli.StringFlag{
					Name:  "output, out",
					Usage: "path to output file",
				},
			},

			Action: func(c *cli.Context) error {
				input := c.String("input")
				inputFile, err := os.Open(input)
				if err != nil {
					return err
				}
				defer inputFile.Close()

				stat, err := inputFile.Stat()
				if err != nil {
					return err
				}

				metadata, err := tag.Read(inputFile)
				if err != nil {
					return err
				}
				tags := tag.GetMap(metadata)

				output := c.String("output")
				extension := filepath.Ext(output)
				switch extension {
				case ".json":
					data, err := json.Marshal(&struct {
						Version  string                 `json:"version"`
						FileName string                 `json:"file name"`
						FileSize int64                  `json:"file size (bytes)"`
						Tags     map[string]interface{} `json:"tags"`
					}{
						Version:  metadata.GetVersion().String(),
						FileName: stat.Name(),
						FileSize: stat.Size(),
						Tags:     tags,
					})
					if err != nil {
						return err
					}

					// save to file
					file, err := os.Create(output)
					if err != nil {
						return err
					}
					defer file.Close()

					_, err = file.Write(data)
					if err != nil {
						return err
					}

				default:
					fmt.Println(fmt.Sprintf("%s: %v", "version", metadata.GetVersion()))
					for key, val := range tags {
						fmt.Println(fmt.Sprintf("%-20s: %v", key, val))
					}
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
