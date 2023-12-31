package src

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"runtime"
)

type pageElements struct {
	progressBar   *widget.ProgressBar // 进度条
	progressText  *widget.Label       // 进度条文字提示
	statusText    *widget.Label       // 状态提示
	buttonRecover *widget.Button      // 恢复按钮
}

func newPageElements() *pageElements {
	p := pageElements{
		progressBar:   widget.NewProgressBar(),
		progressText:  widget.NewLabel(""),
		statusText:    widget.NewLabel(""),
		buttonRecover: nil,
	}

	p.buttonRecover = widget.NewButton("Install", recoverFunc(p.statusText, p.progressBar, p.progressText))

	p.progressBar.Min = 0
	p.progressBar.Max = 1
	return &p
}

func (p *pageElements) createLayout() *fyne.Container {

	// 按钮水平布局
	horizontalLayout := container.NewCenter(p.buttonRecover)

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
		container.NewCenter(widget.NewLabel("OS: "+runtime.GOOS)),
		container.NewCenter(widget.NewLabel("ARCH: "+runtime.GOARCH)),
	)
	return content
}

func RunApp() {

	// 创建和启动应用
	myApp := app.New()
	myWindow := myApp.NewWindow(WindowName)
	myWindow.Resize(fyne.NewSize(WindowWidth, WindowHeight))

	// 创建所有元素
	elements := newPageElements()

	// 布局
	content := elements.createLayout()

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
