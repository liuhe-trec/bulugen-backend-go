# Bulugen Backend GO

## 初始化项目

```readme
// init go project
go mod init bulugen-backend-go

//设置代理
go env -w GOPROXY=https://goproxy.io,direct

//安装一些必要的工具
go get -u -v github.com/stamblerre/gocode@latest
go get -u -v github.com/uudashr/gopkgs/v2/cmd/gopkgs@latest
go get -u -v github.com/ramya-rao-a/go-outline@latest
go get -u -v github.com/acroca/go-symbols@latest
go get -u -v golang.org/x/tools/cmd/guru@latest
go get -u -v golang.org/x/tools/cmd/gorename@latest
go get -u -v github.com/cweill/gotests@latest
go get -u -v github.com/fatih/gomodifytags@latest
go get -u -v github.com/josharian/impl@latest
go get -u -v github.com/davidrjenni/reftools/cmd/fillstruct@latest
go get -u -v github.com/haya14busa/goplay/cmd/goplay@latest
go get -u -v github.com/godoctor/godoctor@latest
go get -u -v github.com/go-delve/delve/cmd/dlv@latest
go get -u -v github.com/stamblerre/gocode@latest
go get -u -v github.com/rogpeppe/godef@latest
go get -u -v github.com/sqs/goreturns@latest
go get -u -v golang.org/x/lint/golint@latest
go get -u -v golang.org/x/tools/gopls@latest
```

```readme
//测试执行
go test -v
```

## Directory structure

```readme
bulugen-backend-go
├── api         # 类似controller
├── cmd         # 通用的命令        
├── conf        # 跟系统相关的配置
├── dao         # 数据库相关的包
├── global      # 放置全局的一些内容
├    └── constants        # 全局变量
├── log         # 日志
├── middleware  # 中间件
├── model       # 映射数据库的实体类
├── router      # 后端路由导航
├── service     # api的逻辑
├    └── dto    # 定义数据组装
├── utils       # 通用的工具         
└── test        # 一些杂项功能代码
```
