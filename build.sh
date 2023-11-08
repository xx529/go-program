# 编译 windows 版本
CGO_ENABLED=1 GOOS=windows GOARCH=amd64 go build -o ./dist/main.exe ./src/main.go

#go build -o ./dist/main ./src/main.go