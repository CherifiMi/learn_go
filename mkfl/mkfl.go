package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {

	app := cli.NewApp()
	app.Name = "mkfl"
	app.Usage = "Create a new file in the current directory"
	app.Action = func(c *cli.Context) error {
		filename := c.Args().First()
		if filename == "" {
			fmt.Println("Usage: mkfl <file.ext>")
			return nil
		}

		// Create or open the file
		file, err := os.Create(filename)
		if err != nil {
			fmt.Println("Error creating the file:", err)
			return err
		}
		defer file.Close()

		fmt.Printf("%s", filename)
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
