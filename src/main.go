package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"os/exec"
	"time"
)

func stop() {
	time.Sleep(1 * time.Second)
}

func main() {

	// 创建和启动应用
	myApp := app.New()
	myWindow := myApp.NewWindow("FundAnalysis")
	myWindow.Resize(fyne.NewSize(300, 350))

	// 创建进度条
	progressBar := widget.NewProgressBar()
	progressBar.Min = 0
	progressBar.Max = 1

	// 进度条文字提示
	progressText := widget.NewLabel("")

	// 状态提示
	statusText := widget.NewLabel("")

	// 安装按钮
	buttonInstall := widget.NewButton("Install", func() {
		statusText.SetText("")
		stop()
		progressBar.SetValue(0.2)
		progressText.SetText("checking docker version...")
		stop()
		progressBar.SetValue(0.3)
		progressText.SetText("checking docker compose version...")
		stop()
		progressBar.SetValue(0.5)
		progressText.SetText("loading images...")
		stop()
		progressBar.SetValue(0.7)
		progressText.SetText("starting containers...")
		stop()
		progressBar.SetValue(1.0)
		progressText.SetText("finish installation")
		stop()
		cmd := exec.Command("open", "http://127.0.0.1:13011")
		err := cmd.Start()
		if err != nil {
			fmt.Println("running at 127.0.0.1:13011")
			return
		}
	})

	// 启动按钮
	buttonStart := widget.NewButton("Start", func() {
		statusText.SetText("")
		stop()
		progressBar.SetValue(0.2)
		progressText.SetText("checking docker version...")
		stop()
		progressBar.SetValue(0.3)
		progressText.SetText("checking docker compose version...")
		stop()
		progressBar.SetValue(0.7)
		progressText.SetText("starting containers...")
		stop()
		progressBar.SetValue(1.0)
		progressText.SetText("finish starting")
		statusText.SetText("running at 127.0.0.1:13011")

	})

	// 更新按钮
	buttonUpdate := widget.NewButton("Update", func() {
		statusText.SetText("")
		stop()
		progressBar.SetValue(0.2)
		progressText.SetText("checking docker version...")
		stop()
		progressBar.SetValue(0.3)
		progressText.SetText("checking docker compose version...")
		stop()
		progressBar.SetValue(0.4)
		progressText.SetText("checking update folder...")
		stop()
		progressBar.SetValue(0.6)
		progressText.SetText("updating images...")
		stop()
		progressBar.SetValue(1.0)
		progressText.SetText("finish updating")
		statusText.SetText("running at 127.0.0.1:13011")

	})

	// 补丁按钮
	buttonPatch := widget.NewButton("Patch", func() {
		statusText.SetText("")
		stop()
		progressBar.SetValue(0.2)
		progressText.SetText("checking docker version...")
		stop()
		progressBar.SetValue(0.3)
		progressText.SetText("checking docker compose version...")
		stop()
		progressBar.SetValue(0.4)
		progressText.SetText("checking patch folder...")
		stop()
		progressBar.SetValue(0.6)
		progressText.SetText("patching...")
		stop()
		progressBar.SetValue(1.0)
		progressText.SetText("finish patching")
		statusText.SetText("running at 127.0.0.1:13011")
	})

	// 水平布局
	horizontalLayout := container.NewHBox(
		widget.NewLabel(""),
		buttonInstall,
		widget.NewLabel(""),
		buttonStart,
		widget.NewLabel(""),
		buttonUpdate,
		widget.NewLabel(""),
		buttonPatch,
		widget.NewLabel(""),
	)
	// 垂直布局
	content := container.NewVBox(
		container.NewCenter(widget.NewLabel("Select your action")),
		horizontalLayout,
		progressBar,
		widget.NewLabel(""),
		container.NewCenter(progressText),
		container.NewCenter(statusText),
	)

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
