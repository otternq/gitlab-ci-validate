package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/urfave/cli"
)

func run(cliContext *cli.Context) error {
	var (
		err    error
		logger = log.New(ioutil.Discard, "", log.Lshortfile|log.Ldate|log.Ltime)

		filePath  = cliContext.String("filepath")
		gitlabURL = cliContext.String("gitlab-url")

		fileContents   []byte
		gitlabClient   = NewGitlabClient(gitlabURL)
		ciLintResponse CILintResponse
	)

	if cliContext.Bool("verbose") {
		logger.SetOutput(os.Stderr)
	}

	gitlabClient.Logger = logger

	if fileContents, err = getFileContents(filePath); err != nil {
		return err
	}

	if ciLintResponse, err = gitlabClient.CILint(fileContents); err != nil {
		return err
	}

	if ciLintResponse.Status != "valid" {
		fmt.Println("Invalid contents:")

		for _, errStr := range ciLintResponse.Errors {
			fmt.Println("Error:", errStr)
		}

		return invalidYamlError
	}

	fmt.Println("Valid gitlab-ci.yml file")
	return nil
}

func getFileContents(filePath string) ([]byte, error) {
	var (
		err  error
		file *os.File
	)

	if file, err = os.Open(filePath); err != nil {
		return []byte{}, err
	}

	return ioutil.ReadAll(file)
}
