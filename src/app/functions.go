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

		t.SetText("checking docker version...")
		dockerVersion(t)
		p.SetValue(0.3)

		t.SetText("checking docker compose version...")
		p.SetValue(0.5)
		dockerComposeVersion(t)

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

// 获取 docker compose 版本
func dockerComposeVersion(l *widget.Label) {
	out, err := exec.Command("docker-compose", "version", "--short").Output()
	if err != nil {
		l.SetText("Docker Compose Version: not installed")
	} else {
		l.SetText("Compose Version: " + string(out))
	}
}

// 启动docker客户端
func startDocker(l *widget.Label) {
	l.SetText("starting docker...")
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
		fmt.Println("open with your browser", url)
		return
	}
}
