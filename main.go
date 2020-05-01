package main

import (
	"bytes"
	"errors"
	"flag"
	"log"
	"strings"

	"github.com/hashicorp/terraform/command"
	"github.com/mitchellh/cli"
)

func main() {
	// terraform init
	if err := initialize(); err != nil {
		log.Fatal("Could not apply/destroy changes", err)
	}

	// terraform apply, destroy
	var destroy bool
	flag.BoolVar(&destroy, "destroy", false, "destroy")
	flag.Parse()

	if err := applyDestroy(destroy); err != nil {
		log.Fatal("Could not apply/destroy changes", err)
	}

	// finished
	if destroy {
		log.Println("Changes destroyed successfully")
	} else {
		log.Println("Changes applied successfully")
	}
}

func initialize() error {
	var r strings.Reader
	var w bytes.Buffer
	var ew bytes.Buffer

	ui := cli.BasicUi{
		Reader:      &r,
		Writer:      &w,
		ErrorWriter: &ew,
	}

	meta := command.Meta{
		Ui:                  &ui,
		RunningInAutomation: true,
	}

	init := command.InitCommand{
		Meta: meta,
	}

	if exitcode := init.Run([]string{"-input=false"}); exitcode == 1 {
		if ew.Len() > 0 {
			return errors.New(string(ew.Bytes()))
		}

		// not expecting this to happen
		return errors.New("Unknown error")
	}
	return nil
}

func applyDestroy(destroy bool) error {
	var r strings.Reader
	var w bytes.Buffer
	var ew bytes.Buffer

	ui := cli.BasicUi{
		Reader:      &r,
		Writer:      &w,
		ErrorWriter: &ew,
	}

	meta := command.Meta{
		Ui:                  &ui,
		RunningInAutomation: true,
	}

	apply := command.ApplyCommand{
		Meta:    meta,
		Destroy: destroy,
	}
	if exitcode := apply.Run([]string{"-input=false", "-auto-approve"}); exitcode == 1 {
		if ew.Len() > 0 {
			return errors.New(string(ew.Bytes()))
		}

		// not expecting this to happen
		return errors.New("Unknown error")
	}
	return nil
}
