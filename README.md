# MorningGo : Gin WebApp Project Skeleton

基于[Gin](https://github.com/gin-gonic/gin)的http项目开发框架。仅供参考。

更高的开发效率，<br>
更好的性能，<br>
更简单整洁的项目组织结构，<br>
更快的部署。

## 环境要求

- [GO >= 1.8](https://github.com/Unknwon/the-way-to-go_ZH_CN/blob/master/eBook/directory.md)

## 项目安装运行

使用安装器安装

### 安装项目

```
cd $GOPATH/src
wget https://github.com/chenhg5/morningo-installer/raw/master/morningo-installer
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


## 项目结构

```

.
├── Makefile
├── README.md
├── command                     命令工具
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
│   └── auth.go
├── frontend                    前端资源
│   ├── assets
│   │   ├── css
│   │   ├── images
│   │   └── js
│   ├── dist
│   └── templates
│       └── index.tpl
├── handle.go                   全局错误处理
├── main.go                     主函数
├── models                      模型
│   └── User.go
├── module                      项目模块
│   ├── cache
│   │   └── cache.go
│   └── session
│       └── session.go
│   └── schedule
│       └── schedule.go         定时任务模块
├── routers.go                  路由
├── routers_test.go             api测试
├── storage                     
│   ├── cache                   缓存文件
│   └── logs                    项目日志
│       ├── access.log          
│       └── error.log
└── vendor                      govendor 第三方包


```

## 项目依赖

- web框架：github.com/gin-gonic/gin
- ORM包：github.com/jinzhu/gorm
- Redis：github.com/go-redis/redis
- Mysql：github.com/go-sql-driver/mysql
- Wechat：github.com/silenceper/wechat
- 任务调度：github.com/robfig/cron

## TODO

- [X] 日志分析
- [X] 测试
- [ ] mysql读写分离
- [ ] redis集群
- [ ] cache/session多存储支持
- [ ] 队列任务
- [ ] 框架性能分析（对标laravel/swoole_php;beego_go）
- [ ] 脚手架命令行工具 sword

## ChangeLog

- 增加测试文件
- 增加平滑重启
- 增加定时任务
- 增加项目安装器
- 增加access.log与error.log
- 增加数据库事务
