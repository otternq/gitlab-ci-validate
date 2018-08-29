// Command gitlab-ci-validate validates the supplies yaml file against Gitlab.
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

var (
	invalidYamlError = fmt.Errorf("invalid file contents")
)

func main() {
	var (
		err error
		app = cli.NewApp()
	)

	app.Name = "gitlab-ci-validate"
	app.Description = "Validate a .gitlab-ci.yml file against Gitlab"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "filepath",
			EnvVar: "FILEPATH",
			Usage:  "The location of the gitlab-ci yaml file",
			Value:  ".gitlab-ci.yml",
		},
		cli.StringFlag{
			Name:   "gitlab-url",
			EnvVar: "GITLAB_URL",
			Usage:  "The Gitlab API to validate with",
			Value:  "https://gitlab.com",
		},
		cli.BoolFlag{
			Name:  "verbose",
			Usage: "Enable for logging",
		},
	}
	app.Action = run

	if err = app.Run(os.Args); err != nil {
		if err.Error() != invalidYamlError.Error() {
			log.Println("execution failed:", err.Error())
		}

		os.Exit(1)
	}
}
