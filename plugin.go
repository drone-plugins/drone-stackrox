package main

import (
	stdlog "log"
	"os"
	"os/exec"
)

type (
	Plugin struct {
		Action     string
		Url        string
		Image      string
		Token      string
		Output     string
		OutputFile string
	}
)

func (p *Plugin) Exec() error {
	os.Setenv("ROX_API_TOKEN", p.Token)
	args := []string{"image", p.Action, "-e", p.Url, "--image", p.Image, "-o", p.Output}
	cmd := exec.Command("roxctl", args...)
	out, err := cmd.Output()
	if err != nil {
		stdlog.Println(err)
		return err
	}
	if p.OutputFile != "" {
		err = os.WriteFile(p.OutputFile, out, 0644)
		if err != nil {
			return err
		}
	}
	stdlog.Println(string(out))
	return nil
}
