package app

import (
	"errors"
	"os"
	"os/exec"
	"runtime"
	"time"
)

// 检测操作系统是否未windows
func checkSystemIsWindows() bool {
	if runtime.GOOS != SupportSystem {
		return false
	}
	return true
}

// 检测 docker desktop.exe 启动文件
func checkDockerDesktopExeIsExist() bool {
	_, err := os.Stat(DefaultDockerDesktopExe)
	if err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	} else {
		return false
	}
}

// 获取docker版本
func getDockerVersion() (string, error) {
	out, err := exec.Command("docker", "version", "--format", "'{{.Server.Version}}'").Output()
	if err != nil {
		return "", err
	} else {
		return string(out), nil
	}
}

// 获取 docker compose 版本
func getDockerComposeVersion() (string, error) {
	out, err := exec.Command("docker", "compose", "version", "--short").Output()
	if err != nil {
		return "", err
	} else {
		return string(out), nil
	}
}

// 检查 docker 是否正在运行
func checkDockerIsRunning() bool {
	_, err := getDockerVersion()
	if err != nil {
		return false
	}
	return true
}

// 启动 docker 客户端
func startDockerClient() error {

	// 检测操作系统是否未windows
	if !checkSystemIsWindows() {
		return errors.New("only support windows system")
	}

	// 检测 docker desktop.exe 启动文件
	if !checkDockerDesktopExeIsExist() {
		return errors.New("fail to check Docker Desktop.exe file")
	}

	// 启动 Docker Desktop
	signal := make(chan error)
	quit := make(chan struct{})
	go func() {

		var err error

		// 进入文件夹所在文件夹
		err = os.Chdir(DefaultDockerDesktopDir)
		if err != nil {
			signal <- errors.New("fail to change direction " + DefaultDockerDesktopDir)
			return
		}

		// 执行启动文件
		cmd := "./" + DefaultExeFileName
		err = exec.Command(cmd).Start()
		if err != nil {
			signal <- err
			return
		}

		// 检测docker是否在运行当中
		for {
			select {
			case <-quit:
				return
			default:
				if checkDockerIsRunning() {
					signal <- nil
					return
				}
			}
		}
	}()

	// 检测是否成功启动 docker
	timer := time.NewTimer(DockerOpTimeOutSeconds * time.Second)
	select {
	case err := <-signal:
		if err != nil {
			return err
		} else {
			return nil
		}
	case <-timer.C:
		quit <- struct{}{}
		return errors.New("start docker time out")
	}
}

// 关闭 docker 客户端
func stopDockerClient() error {

	// 检测操作系统是否未windows
	if !checkSystemIsWindows() {
		return errors.New("only support windows system")
	}

	// 检测 docker desktop.exe 启动文件
	if !checkDockerDesktopExeIsExist() {
		return errors.New("fail to check Docker Desktop.exe file")
	}

	// 启动 Docker Desktop
	signal := make(chan error)
	quit := make(chan struct{})

	go func() {
		var err error

		// 执行杀进程命令
		err = exec.Command("taskkill", "/IM", DefaultExeFileName, "/F").Start()
		if err != nil {
			signal <- err
			return
		}

		// 检测docker是否在运行当中
		for {
			select {
			case <-quit:
				return
			default:
				if !checkDockerIsRunning() {
					signal <- nil
					return
				}
			}
		}
	}()

	// 检测是否成功关闭 docker
	timer := time.NewTimer(DockerOpTimeOutSeconds * time.Second)
	select {
	case err := <-signal:
		if err != nil {
			return err
		} else {
			return nil
		}
	case <-timer.C:
		quit <- struct{}{}
		return errors.New("stop docker time out")
	}
}
