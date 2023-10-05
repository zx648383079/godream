# Go-ZoDream
Go语言学习

运行

```
go run main.go
```

打包运行
```
go build
.\zodream
```

更新
```
go get -u
```


安装环境

```
$env:GOPROXY="https://goproxy.io"
```

```
$GOPATH/src/golang.org/x/tools
git clone https://github.com/golang/tools.git ./
go install golang.org/x/tools/cmd/guru
go install golang.org/x/tools/cmd/gorename
go install golang.org/x/tools/cmd/fiximports
go install golang.org/x/tools/cmd/gopls
go install golang.org/x/tools/cmd/godex

```

### live 模块说明

代码来源 [livego](https://github.com/gwuhaolin/livego) 

同时与本项目进行逻辑结构适配