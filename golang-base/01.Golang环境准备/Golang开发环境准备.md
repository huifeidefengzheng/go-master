# Golang开发环境准备

# 0.1 Golang 环境配置

Golang 官网下载页面：https://go.dev/dl/，下载需要科学上网。

国内 Golang 官网下载页面：[https://golang.google.cn/dl/](https://golang.google.cn/dl/)，国内可以直接下载。

按照自己的系统下载对应版本：

windows 系统：

32 位下载 [go1.23.0.windows-386.zip](https://golang.google.cn/dl/go1.23.0.windows-386.zip)；

64 位下载<u>go1.23.0.windows-amd64.zip</u>。

linux 系统：

传统 32 位下载 [go1.23.0.linux-386.tar.gz](https://golang.google.cn/dl/go1.23.0.linux-386.tar.gz)；

传统 64 位下载<u>go1.23.0.linux-amd64.tar.gz</u>；

arm 架构下载<u>go1.23.0.linux-arm64.tar.gz</u>。

macOS 系统：

intel 平台下载 [go1.23.0.darwin-amd64.tar.gz](https://golang.google.cn/dl/go1.23.0.darwin-amd64.tar.gz)；

arm 平台下载 [go1.23.0.darwin-arm64.tar.gz](https://golang.google.cn/dl/go1.23.0.darwin-arm64.tar.gz)。

注：尽量下载压缩包版本，这样可以自己决定 go 环境的路径。

下载好 golang 的压缩包之后解压:

```shell
mkdir $HOME/go1.22.6 && tar -xf go1.22.6.darwin-arm64.tar.gz --strip 1 -C $HOME/go1.22.6
```

设置 Golang 所需的环境变量:

```shell
export GOROOT=$HOME/go1.22.6
# go的工作环境，go语言项目中声明的依赖项目的代码会被下载到这个路径下
export GOPATH=$HOME/.go_path
export PATH=$GOROOT/bin:$GOPATH/bin:$PATH
```

检查 go 语言环境:

```shell
go version
```

设置 GOPROXY:

```shell
go env -w GOPROXY="https://goproxy.cn,direct"
```

# 0.2 Vscode 环境准备

一般推荐使用 goland IDE 开发 golang，但是 goland 收费，并且在某些情况下破解版并不稳定，这里会介绍一下 Vscode 开发 Golang 环境准备事项。

## 0.2.1 插件安装

为了避免插件安装错，这里提供了插件网页地址，可以使用这个地址把插件安装到本地 Vscode 环境中。

### 0.2.1.1 Go 插件

[https://marketplace.visualstudio.com/items?itemName=golang.Go](https://marketplace.visualstudio.com/items?itemName=golang.Go)

### 0.2.1.2 Debugger 插件

[https://marketplace.visualstudio.com/items?itemName=wowbox.code-debuger](https://marketplace.visualstudio.com/items?itemName=wowbox.code-debuger)

注：虽然安装完 Go 插件之后，就可以 debug go 代码，但是如果启动 go 的 main 函数需要使用 args 传递参数时，实测使用这个插件，兼容性更好。

### 0.2.1.3 AI 辅助编码插件

[https://marketplace.visualstudio.com/items?itemName=sourcegraph.cody-ai](https://marketplace.visualstudio.com/items?itemName=sourcegraph.cody-ai)

# 0.3 初始化一个项目

首先创建一个目录并进入该目录中。

```go
mkdir init_project && cd init_project
```

初始化 go.mod 文件。

```go
go mod init github.com/test/init_project
```

用 IDE 工具打开这个项目目录并创建文件 main.go，复制粘贴以下内容。

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, world!")
}
```

运行程序。

```shell
go run main.go
```
