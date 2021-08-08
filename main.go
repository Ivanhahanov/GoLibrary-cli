package main

import (
	"fmt"
	"github.com/Ivanhahanov/GoLibrary-cli/client"
	"log"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
)

func searchRequest(c *cli.Context) error {
	var searchString string
	if c.NArg() > 0 {
		searchString = strings.Join(c.Args().Slice(), " ")
	}

	if c.String("lang") == "ru" {
		fmt.Println("Language: Russian")
	} else {
		fmt.Println("Language: English")
	}

	if c.String("title") != "" {
		fmt.Println("Title:", c.String("title"))
	}

	if c.String("tags") != "" {
		tags := strings.Split(c.String("tags"), ",")
		fmt.Println("Tags:", tags)
	}

	if searchString != "" {
		fmt.Println("In content:", searchString)
	}
	return nil
}

var searchFlags = []cli.Flag{
	&cli.StringFlag{
		Name:  "lang",
		Value: "english",
		Usage: "language for result",
	},
	&cli.StringFlag{
		Name:    "title",
		Aliases: []string{"T"},
		Usage:   "Search by Title",
	},
	&cli.StringFlag{
		Name:    "tags",
		Aliases: []string{"t"},
		Usage:   "Search by Tags",
	},
}

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:  "search",
				Usage: "options for task templates",
				Flags: searchFlags,
				Action: func(c *cli.Context) error {
					return searchRequest(c)
				},
				Subcommands: []*cli.Command{
					{
						Name:  "books",
						Usage: "search books",
						Flags: searchFlags,
						Action: func(c *cli.Context) error {
							client.GetBooks()
							return nil
						},
					},
					{
						Name:  "articles",
						Usage: "search articles",
						Flags: searchFlags,
						Action: func(c *cli.Context) error {
							// TODO: search articles
							searchRequest(c)
							return nil
						},
					},
					{
						Name:  "links",
						Usage: "search links",
						Flags: searchFlags,
						Action: func(c *cli.Context) error {
							// TODO: search links
							searchRequest(c)
							return nil
						},
					},
				},
			},
			{
				Name:  "login",
				Usage: "options for login",
				Action: func(c *cli.Context) error {
					url := c.Args().Get(0)
					if url == "" {
						fmt.Println("Url doesn't specify")
					} else {
						fmt.Printf("Log in to %q\n", url)
					}
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
