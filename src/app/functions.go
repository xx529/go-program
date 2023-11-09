package app

import (
	"fmt"
	"fyne.io/fyne/v2/widget"
	"os/exec"
)

// 恢复按钮执行的函数
func recoverFunc(st *widget.Label, pb *widget.ProgressBar, pt *widget.Label) func() {
	return func() {

		st.SetText("")
		pb.SetValue(0.0)

		pt.SetText("checking system...")
		if !checkSystemIsWindows() {
			st.SetText("only support windows")
			pt.SetText("finished")
			pb.SetValue(1.0)
			return
		}

		pt.SetText("checking docker is running or not...")
		pb.SetValue(0.3)
		if !checkDockerIsRunning() {
			st.SetText("docker is not running")

			pt.SetText("starting docker...")
			pb.SetValue(0.6)
			err := startDockerClient()
			if err != nil {
				st.SetText("fail to start docker")
				pb.SetValue(1.0)
			} else {
				st.SetText("docker is running now")
				pb.SetValue(1.0)
			}
		} else {
			st.SetText("docker is running")

			st.SetText("stop docker...")
			pb.SetValue(0.5)
			err := stopDockerClient()
			if err != nil {
				st.SetText("fail to stop docker")
				pb.SetValue(1.0)
			} else {
				st.SetText("docker is stopped")
				pb.SetValue(0.7)

				pt.SetText("starting docker...")
				pb.SetValue(0.8)
				err := startDockerClient()
				if err != nil {
					st.SetText("fail to start docker")
					pb.SetValue(1.0)
				} else {
					st.SetText("docker is running now")
					pb.SetValue(1.0)
				}
			}

		}
		pt.SetText("finished")
		openBrowser(RunningUrl)
	}
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
