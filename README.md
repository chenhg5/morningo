# MorningGo : Gin WebApp Project Skeleton

[![Go Report Card](https://goreportcard.com/badge/github.com/chenhg5/morningo)](https://goreportcard.com/report/github.com/chenhg5/morningo) [![Build Status](https://api.travis-ci.org/chenhg5/morningo.svg?branch=master)](https://api.travis-ci.org/chenhg5/morningo) [![sonarcloud](https://sonarcloud.io/api/project_badges/measure?project=morningo&metric=alert_status)](https://sonarcloud.io/dashboard?id=morningo)

![doggy](https://ws2.sinaimg.cn/large/006tKfTcgy1fr2um9hwduj303u03w0sv.jpg)

[中文文档](./README_CN.md)

A Web develop project skeleton base on [Gin](https://github.com/gin-gonic/gin) which just for reference.

More efficiency,<br>
Faster and clear,<br>
Easier to deploy

Suitable for simple project. [ko](https://github.com/chenhg5/ko)、[kit](https://github.com/go-kit/kit)、[go-micro](https://github.com/micro/go-micro)、[kite](https://github.com/koding/kite) are better choice for the middle and large project.

## Environment Requirements 

- [GO >= 1.8](https://github.com/Unknwon/the-way-to-go_ZH_CN/blob/master/eBook/directory.md)

## Installation And Run

Via moroingo installer or use docker

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

### Graceful Restart

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
kill -INT $(cat pid) && ./morningo # graceful stop the process and restart
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
│   ├── auth                    auth middleware
│   │   ├── drivers             auth engine
│   │   └── auth.go   
│   └── filter.go               middleware initer                  
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
│   │── schedule
│   │   └── schedule.go   
│   │── logger
│   │   └── logger.go 
│   └── server
│       └── server.go 
├── routers                     routers
│   └── api_routers.go       
├── routers.go                  router initer
├── routers_test.go             unit test for api
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

## Benchmark

```
wrk -c100 -d30s -t4 http://localhost:4000/api/index

Running 30s test @ http://localhost:4000/api/index
  4 threads and 100 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     4.03ms    4.17ms  72.72ms   89.95%
    Req/Sec     7.40k     1.67k   11.20k    67.75%
  884816 requests in 30.03s, 107.17MB read
  Non-2xx or 3xx responses: 884816
Requests/sec:  29462.96
Transfer/sec:  3.57MB
```

## TODO

- [X] Logger
- [X] Test
- [X] Cache/Session
- [ ] Queue of task
- [ ] Read & Write Connections
- [ ] Redis cluster
- [ ] Profiling(Laravel/Swoole;beego）
- [ ] Command tool
- [ ] Interaction command env
- [ ] Fast CRUD Generator

## ChangeLog

- Add Reverse Proxy
- Fixed the path
- Add session/cache and Auth middleware
- Add test
- Add graceful restart
- Add schedule module
- Add installer of project
- Add access.log and error.log
- Add database transcation
