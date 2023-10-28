package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"runtime"
)

type pageElements struct {
	progressBar       *widget.ProgressBar // 进度条
	progressText      *widget.Label       // 进度条文字提示
	statusText        *widget.Label       // 状态提示
	dockerVersionInfo *widget.Label       // docker 版本信息
	dockerComposeInfo *widget.Label       // docker compose 版本信息
	buttonInstall     *widget.Button      // 安装按钮
}

func newPageElements() *pageElements {
	p := pageElements{
		progressBar:       widget.NewProgressBar(),
		progressText:      widget.NewLabel(""),
		statusText:        widget.NewLabel(""),
		dockerVersionInfo: widget.NewLabel("checking docker version...\n"),
		dockerComposeInfo: widget.NewLabel("checking docker compose version...\n"),
		buttonInstall:     widget.NewButton("Install", nil),
	}
	p.progressBar.Min = 0
	p.progressBar.Max = 1
	return &p
}

func (p *pageElements) layout() *fyne.Container {

	// 按钮水平布局
	horizontalLayout := container.NewCenter(p.buttonInstall)

	// 垂直布局
	content := container.NewVBox(
		widget.NewLabel(""),
		horizontalLayout,
		widget.NewLabel(""),
		p.progressBar,
		widget.NewLabel(""),
		container.NewCenter(p.progressText),
		container.NewCenter(p.statusText),
		widget.NewLabel(""),
		widget.NewLabel(""),
		container.NewCenter(widget.NewLabel("OS: "+runtime.GOOS)),
		container.NewCenter(widget.NewLabel("ARCH: "+runtime.GOARCH)),
		container.NewCenter(p.dockerVersionInfo),
	)
	return content
}

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
