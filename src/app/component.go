package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"runtime"
)

func RunApp() {

	// 创建和启动应用
	myApp := app.New()
	myWindow := myApp.NewWindow(WindowName)
	myWindow.Resize(fyne.NewSize(WindowWidth, WindowHeight))

	// 创建进度条
	progressBar := widget.NewProgressBar()
	progressBar.Min = 0
	progressBar.Max = 1

	// 进度条文字提示
	progressText := widget.NewLabel("")

	// 状态提示
	statusText := widget.NewLabel("")

	// docker 信息
	dockerText := widget.NewLabel("checking docker version...\n")
	dockerVersion(dockerText)

	// 安装按钮
	buttonInstall := widget.NewButton("Install",
		installFunc(statusText, progressBar, progressText),
	)

	// 布局
	content := layout(buttonInstall, progressBar, progressText, statusText, dockerText)

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}

func layout(buttonInstall *widget.Button, progressBar *widget.ProgressBar, progressText *widget.Label, statusText *widget.Label, dockerText *widget.Label) *fyne.Container {

	// 按钮水平布局
	horizontalLayout := container.NewCenter(buttonInstall)

	// 垂直布局
	content := container.NewVBox(
		widget.NewLabel(""),
		horizontalLayout,
		widget.NewLabel(""),
		progressBar,
		widget.NewLabel(""),
		container.NewCenter(progressText),
		container.NewCenter(statusText),
		widget.NewLabel(""),
		widget.NewLabel(""),
		container.NewCenter(widget.NewLabel("OS: "+runtime.GOOS)),
		container.NewCenter(widget.NewLabel("ARCH: "+runtime.GOARCH)),
		container.NewCenter(dockerText),
	)
	return content
}
