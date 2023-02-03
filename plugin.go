package main

import (
	stdlog "log"
	"os"
	"os/exec"
)

type (
	Plugin struct {
		Action string
		Url    string
		Image  string
		Token  string
		Output string
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
	stdlog.Println(string(out))
	return nil
}
