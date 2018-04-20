# MorningGo : Gin WebApp Project Skeleton

[![Go Report Card](https://goreportcard.com/badge/github.com/chenhg5/morningo)](https://goreportcard.com/report/github.com/chenhg5/morningo)

[中文文档](./README_CN.md)

A Web develop project skeleton base on [Gin](https://github.com/gin-gonic/gin) which just for reference.

More efficiency,<br>
Faster and clear,<br>
Easier to deploy

Suitable for simple project. [kit](https://github.com/go-kit/kit)、[go-micro](https://github.com/micro/go-micro)、[kite](https://github.com/koding/kite) are better choice for the middle and large project.

## Environment Requirements

- [GO >= 1.8](https://github.com/Unknwon/the-way-to-go_ZH_CN/blob/master/eBook/directory.md)

## Installation And Run

Via moroingo installer

### Install

```
cd $GOPATH/src

wget https://gitee.com/cg33/morningo-installer/raw/master/morningo-installer       # mac
wget https://gitee.com/cg33/morningo-installer/raw/master/morningo-installer-linus # linus
wget https://gitee.com/cg33/morningo-installer/raw/master/morningo-installer.exe   # windows

chmod +x morningo-installer
./morningo-installer --project-name web
```

### Load Dependency

```
cd web
make deps
```

### Test

```
make test
```

### Smooth Restart

```
make restart
```

### Run It

```
make
```
visit by browser: http://localhost:4000/api/index

## Deploy

First build the executable file
```
make build # for linus
make cross # for mac/windows
```
Then put files of the ```build``` in your server and set the path of log and static file(html/css/js),and run the executable file.If 80 port is not allowed to use,consider the nginx proxy,or use the gin middleware [gin-reverseproxy](https://github.com/chenhg5/gin-reverseproxy) instead, which has some example in ```routers.go```. When the project start running, it will generate the ```pid```file in the root path of the project. Excute the following command to update your project. 
```
kill -INT $(cat pid) && ./morningo # smooth stop the process and restart
```

## Project Structure

```

.
├── Makefile
├── README.md
├── command                     
│   └── sword.go
├── config                      global config
│   └── env.go
├── connections                 store connection
│   ├── database
│   │   ├── mongodb
│   │   └── mysql
│   └── redis
│       └── redis.go
├── controllers                 controller
│   └── MainController.go
├── filters                     middleware
│   └── auth                    auth middleware
│       ├── drivers             auth engine
│       └── auth.go             
├── frontend                    frontend resource
│   ├── assets
│   │   ├── css
│   │   ├── images
│   │   └── js
│   ├── dist
│   └── templates
│       └── index.tpl
├── handle.go                   global error handler
├── main.go                     
├── models                      model
│   └── User.go
├── module                      module of project
│   └── schedule
│       └── schedule.go         
├── routers.go                  router
├── routers_test.go             test for api
├── storage                     
│   ├── cache                   cache file
│   └── logs                    log file
│       ├── access.log          
│       └── error.log
└── vendor                      govendor vendor


```

## What`s in the box

### HTTP (based on [Gin](https://github.com/gin-gonic/gin))
- Router
- Middleware
- Controller
- Request
- Response
- View
- Session

### Frontend
- Go template

### Security
- Authentication
- Authorization
- Encryption
- Hash

### Digging Deeper
- Dancer Command
- Cache System
- Error and Log
- Schedule

### Database
- Mysql
- Mongodb
- Redis

### ORM(based on [gorm](https://github.com/jinzhu/gorm))

### test
- Api test

## Project Dependency

- web framework：github.com/gin-gonic/gin
- orm：github.com/jinzhu/gorm
- redis：github.com/go-redis/redis
- mysql：github.com/go-sql-driver/mysql
- wechat：github.com/silenceper/wechat
- schedule：github.com/robfig/cron

## TODO

- [X] Logger
- [X] Test
- [ ] Queue of task
- [ ] Cache/Session
- [ ] Read & Write Connections
- [ ] Redis cluster
- [ ] Profiling(Laravel/Swoole;beego）
- [ ] Command tool
- [ ] Interaction command env

## ChangeLog

- Add Reverse Proxy
- Fixed the path
- Add session/cache and Auth middleware
- Add test
- Add smooth restart
- Add schedule module
- Add installer of project
- Add access.log and error.log
- Add database transcation
