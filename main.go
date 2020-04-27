package main

import (
	"bytes"
	"flag"
	"log"
	"strings"

	"github.com/hashicorp/terraform/command"
	"github.com/mitchellh/cli"
)

func main() {
	// flag for apply vs. destroy
	var destroy bool
	flag.BoolVar(&destroy, "destroy", false, "destroy")
	flag.Parse()

	// set up command meta used by all commands
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

	// init
	init := command.InitCommand{
		Meta: meta,
	}
	if exitcode := init.Run([]string{}); exitcode == 1 {
		log.Fatal("Could not initialize")
	}

	// apply/destroy
	apply := command.ApplyCommand{
		Meta:    meta,
		Destroy: destroy,
	}
	if exitcode := apply.Run([]string{"-auto-approve=true"}); exitcode == 1 {
		log.Fatal("Could not apply changes")
	}

	// finished
	if destroy {
		log.Println("Changes destroyed successfully")
	} else {
		log.Println("Changes applied successfully")
	}
}
