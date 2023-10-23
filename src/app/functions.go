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

// 安装按钮执行的函数
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
		openBrowser(RunningUrl)
	}
}

// 获取docker版本
func dockerVersion(l *widget.Label) {
	out, err := exec.Command("docker", "version", "--format", "'{{.Server.Version}}'").Output()
	if err != nil {
		startDocker(l)
	} else {
		l.SetText("Docker Version: " + string(out))
	}
}

// 启动docker客户端
func startDocker(l *widget.Label) {
	l.SetText("starting docker... \n")
	err := exec.Command("open", "/Applications/Docker.app").Start()
	if err != nil {
		l.SetText("fail to start docker")
	}

	go func() {
		for {
			time.Sleep(1 * time.Second)
			out, err := exec.Command("docker", "version", "--format", "'{{.Server.Version}}'").Output()
			if err != nil {
				continue
			} else {
				l.SetText("Docker Version: " + string(out))
				return
			}
		}
	}()
}

// 打开浏览器
func openBrowser(url string) {
	cmd := exec.Command("open", url)
	err := cmd.Start()
	if err != nil {
		fmt.Println("running at", url)
		return
	}
}