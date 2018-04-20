# MorningGo : Gin WebApp Project Skeleton

[![Go Report Card](https://goreportcard.com/badge/github.com/chenhg5/morningo)](https://goreportcard.com/report/github.com/chenhg5/morningo)

基于[Gin](https://github.com/gin-gonic/gin)的web项目开发框架。仅供参考。

更高的开发效率，<br>
更好的性能，<br>
更简单整洁的项目组织结构，<br>
更快的部署。

适合于小型项目，大中型项目(pv高、需求复杂度高)移步 [kit](https://github.com/go-kit/kit)、[go-micro](https://github.com/micro/go-micro)、[kite](https://github.com/koding/kite)

## 环境要求

- [GO >= 1.8](https://github.com/Unknwon/the-way-to-go_ZH_CN/blob/master/eBook/directory.md)

## 项目安装运行

使用安装器安装 

### 安装项目

```
cd $GOPATH/src

wget https://gitee.com/cg33/morningo-installer/raw/master/morningo-installer       # mac
wget https://gitee.com/cg33/morningo-installer/raw/master/morningo-installer-linus # linus
wget https://gitee.com/cg33/morningo-installer/raw/master/morningo-installer.exe   # windows

chmod +x morningo-installer
./morningo-installer --project-name web
```

### 加载依赖

```
cd web
make deps
```

### 测试

```
make test
```

### 平滑重启

```
make restart
```

### 运行

```
make
```
浏览器访问 http://localhost:4000/api/index

## 项目部署

生成可执行文件
```
make build # linus用户
make cross # mac/windows用户
```
将```build```下文件上传到生产环境服务器，并设置好日志文件路径以及静态文件路径，然后直接运行即可。如端口不为80端口或有多个域名，可以配置nginx代理，或者采用反向代理中间件[gin-reverseproxy](https://github.com/chenhg5/gin-reverseproxy), 关于代理的使用，```routers.go```中有示例。运行的同时会在文件夹下生成```pid```文件，每次更新完文件后执行如下命令即可平滑热更。
```
kill -INT $(cat pid) && ./morningo # 即停止旧的进程，重启新的执行文件
```

## 项目结构

```

.
├── Makefile
├── README.md
├── command                     
│   └── sword.go
├── config                      全局配置
│   └── env.go
├── connections                 存储连接
│   ├── database
│   │   ├── mongodb
│   │   └── mysql
│   └── redis
│       └── redis.go
├── controllers                 控制器
│   └── MainController.go
├── filters                     中间件
│   ├── auth                    认证中间件
│   │   ├── drivers             认证引擎
│   │   └── auth.go   
│   └── filter.go              
├── frontend                    前端资源
│   ├── assets
│   │   ├── css
│   │   ├── images
│   │   └── js
│   ├── dist
│   └── templates
│       └── index.tpl
├── handle.go                   全局错误处理
├── main.go                     
├── models                      模型
│   └── User.go
├── module                      项目模块
│   │── schedule
│   │   └── schedule.go   
│   │── logger
│   │   └── logger.go 
│   └── server
│       └── server.go        定时任务模块
├── routers                     路由
│   └── api_routers.go          
├── routers.go                  路由初始化设置
├── routers_test.go             api测试
├── storage                     
│   ├── cache                   缓存文件
│   └── logs                    项目日志
│       ├── access.log          
│       └── error.log
└── vendor                      govendor 第三方包


```

## 箱子里有什么 what`s in the box

### HTTP 层(基于[Gin](https://github.com/gin-gonic/gin))
- 路由
- 中间件
- 控制器
- 请求
- 响应
- 视图
- Session

### 前端
- tpl模板

### 安全
- 用户认证
- 用户授权
- 加密解密
- 哈希

### 综合话题
- dancer 命令行
- 缓存系统
- 错误与日志
- 任务调度

### 数据库
- mysql
- mongodb
- redis

### ORM(基于[gorm](https://github.com/jinzhu/gorm))

### 测试
- api 测试

## 项目依赖

- web框架：github.com/gin-gonic/gin
- ORM包：github.com/jinzhu/gorm
- Redis：github.com/go-redis/redis
- Mysql：github.com/go-sql-driver/mysql
- Wechat：github.com/silenceper/wechat
- 任务调度：github.com/robfig/cron

## TODO

- [X] 日志
- [X] 测试
- [ ] 队列任务支持
- [ ] cache/session多存储支持
- [ ] mysql读写分离
- [ ] redis集群
- [ ] 框架性能分析（对标laravel/swoole_php;beego_go）
- [ ] 命令行工具
- [ ] 命令行交互环境

## ChangeLog

- 修改认证组件逻辑
- 增加反向代理示例
- 修复文件相对路径问题
- 增加session、cache、认证中间件
- 增加测试文件
- 增加平滑重启
- 增加定时任务
- 增加项目安装器
- 增加access.log与error.log
- 增加数据库事务
