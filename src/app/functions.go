package app

import (
	"fmt"
	"fyne.io/fyne/v2/widget"
	"os/exec"
	"time"
)

func stop() {
	time.Sleep(1 * time.Second)
}

func installFunc(s *widget.Label, p *widget.ProgressBar, t *widget.Label) func() {
	return func() {
		s.SetText("")
		stop()
		p.SetValue(0.2)
		t.SetText("checking docker version...")
		stop()
		p.SetValue(0.3)
		t.SetText("checking docker compose version...")
		stop()
		p.SetValue(0.5)
		t.SetText("loading images...")
		stop()
		p.SetValue(0.7)
		t.SetText("starting containers...")
		stop()
		p.SetValue(1.0)
		t.SetText("finish installation")
		stop()
		cmd := exec.Command("open", RunningUrl)
		err := cmd.Start()
		if err != nil {
			fmt.Println("running at", RunningUrl)
			return
		}
	}
}
