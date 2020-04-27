package main

import (
	"bytes"
	"log"
	"strings"

	"github.com/hashicorp/terraform/command"
	"github.com/mitchellh/cli"
)

func main() {
	var r strings.Reader
	var w bytes.Buffer
	var ew bytes.Buffer
	ui := cli.BasicUi{
		Reader:      &r,
		Writer:      &w,
		ErrorWriter: &ew,
	}

	meta := command.Meta{
		Ui: &ui,
	}

	init := command.InitCommand{
		Meta: meta,
	}
	if exitcode := init.Run([]string{}); exitcode == 1 {
		log.Fatal("Could not initialize")
	}

	apply := command.ApplyCommand{
		Meta: meta,
	}
	if exitcode := apply.Run([]string{"-auto-approve=true"}); exitcode == 1 {
		log.Fatal("Could not apply changes")
	}

	log.Println("Changes applied successfully")
}
