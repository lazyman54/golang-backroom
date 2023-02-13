# 查看go的env：
go env/go env GOPATH
```
GO111MODULE=""
GOARCH="amd64"
GOBIN=""
GOCACHE="/Users/eric.mao/Library/Caches/go-build"
GOENV="/Users/eric.mao/Library/Application Support/go/env"
GOEXE=""
GOEXPERIMENT=""
GOFLAGS=""
GOHOSTARCH="amd64"
GOHOSTOS="darwin"
GOINSECURE=""
GOMODCACHE="/Users/eric.mao/go/pkg/mod"
GONOPROXY="*.byted.org,*.everphoto.cn,git.smartisan.com"
GONOSUMDB="*.byted.org,*.everphoto.cn,git.smartisan.com"
GOOS="darwin"
GOPATH="/Users/eric.mao/go"
GOPRIVATE="*.byted.org,*.everphoto.cn,git.smartisan.com"
GOPROXY="https://go-mod-proxy.byted.org,https://goproxy.cn,https://proxy.golang.org,direct"
GOROOT="/usr/local/Cellar/go/1.19/libexec"
GOSUMDB="sum.golang.google.cn"
GOTMPDIR=""
GOTOOLDIR="/usr/local/Cellar/go/1.19/libexec/pkg/tool/darwin_amd64"
GOVCS=""
GOVERSION="go1.19"
GCCGO="gccgo"
GOAMD64="v1"
AR="ar"
CC="clang"
CXX="clang++"
CGO_ENABLED="1"
GOMOD="/dev/null"
GOWORK=""
CGO_CFLAGS="-g -O2"
CGO_CPPFLAGS=""
CGO_CXXFLAGS="-g -O2"
CGO_FFLAGS="-g -O2"
CGO_LDFLAGS="-g -O2"
PKG_CONFIG="pkg-config"
GOGCCFLAGS="-fPIC -arch x86_64 -m64 -pthread -fno-caret-diagnostics -Qunused-arguments -fmessage-length=0 -fdebug-prefix-map=/var/folders/vj/hw_f087d6q32njrss5y0hl5w0000gp/T/go-build1202119747=/tmp/go-build -gno-record-gcc-switches -fno-common"
```
修改env值：go env -w go111module=off
* GOROOT: GO的安装目录
* GOPATH：我们的开发区，用来存放我们项目的源代码，可以设置多个值，用";"号分割
1. src：放置源码的地方
2. bin:存放编译后可执行文件的地方 
3. pkg：存放编译后的包文件  
* GOBIN：用来存放我们项目代码编译后生成的二进制文件（可执行文件），当我们使用go install命令编译打包时，会把生成的二进制执行文件放入到GOBIN指定的目录下，默认是GOPATH/bin
* GOOS：为其编译代码的操作系统，如：linux、darwin、windows、netbsd、freebsd、openbsd、solaris（macOS操作系统对应的值是darwin）
* GOARCH：为其编译代码的CPU架构或处理器。比如amd64、386、arm等
* GOPROXY: go get 下载依赖时使用的代理地址列表，可以设置多个值，用都好分割
* GO111MODULE:go modules 是go语言的依赖解决方案，发布于go1.11，到go1.14后生产上推荐使用
1. go111module=off，无模块支持，当我们编译执行一个go源码文件时，如果有依赖其他包，则先查找GOROOT，再查找GOPATH
2. go111module=on，模块支持，go命令行会使用modules，而不会区查找gopath
3. go111module=auto，默认值，go命令行会根据当前目录来决定是否启用module功能