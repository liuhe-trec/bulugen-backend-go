# Bulugen Backend GO

## Use Frameworks

* [**Viper**](https://pkg.go.dev/github.com/spf13/viper#section-readme) :A complete configuration solution for Go applications
* [**Gin**](https://pkg.go.dev/github.com/gin-gonic/gin) :A web framework written in Go
* [**GinSwagger**](https://pkg.go.dev/github.com/swaggo/gin-swagger#section-readme) :Gin middleware to automatically generate RESTful API documentation with Swagger
* [**Zap**](https://pkg.go.dev/go.uber.org/zap) :Blazing fast, structured, leveled logging in Go
* [**lumberjack**](https://github.com/natefinch/lumberjack) :A Go package for writing logs to rolling files
* [**Gorm**](https://gorm.io) :A popular Object Relational Mapping (ORM) library for the Go programming language
* [**go-redis**](https://pkg.go.dev/github.com/go-redis/redis/v8) :Redis client for Go
* [**jwt-go**](https://pkg.go.dev/github.com/golang-jwt/jwt/v5) :A go (or 'golang' for search engine friendliness) implementation of JSON Web Tokens
* [**BCrypt**](https://pkg.go.dev/golang.org/x/crypto/bcrypt) : Encryption related

## Middleware

* [**CORS gin's middleware**](https://pkg.go.dev/github.com/gin-contrib/cors) :Gin middleware/handler to enable CORS support.

## Develop Tools

* [**delve**](https://github.com/go-delve/delve) : for debug

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

## Swagger documentation

swagger: /swagger/index.html
