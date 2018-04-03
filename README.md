# gin web项目模板

gin web项目模板


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
│   ├── alipay
│   │   └── alipay.go
│   ├── cache
│   │   └── cache.go
│   └── session
│       └── session.go
├── routers.go                  路由
└── test                        测试


```

## 项目依赖

- web框架：github.com/gin-gonic/gin
- ORM包：github.com/jinzhu/gorm
- Redis：github.com/go-redis/redis
- Mysql：github.com/go-sql-driver/mysql
- Wechat：github.com/silenceper/wechat