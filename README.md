### gin框架的一个demo

### 环境简介
```.env
系统环境：1.16sdk + window
```

### 功能实现
```.env
①请求参数绑定
②路由转发
③配置文件读取
④日期方法（练习）
⑤日志打印、切割
⑥全局异常捕捉
⑦swagger
```

### 目录介绍
```.env
│  .env 自定义配置文件
│  .env.develop 自定义配置文件
│  go.mod   
│  go.sum
│  main.go  主方法
│  README.md    
│
├─.idea
│      .gitignore
│      go_gin.iml
│      misc.xml
│      modules.xml
│      vcs.xml
│      workspace.xml
│
├─docs  swagger文件（自动生成）
│      docs.go
│      swagger.json
│      swagger.yaml
│
├─internal  操作层
│  ├─config 配置层
│  │      config.go 配置信息结构
│  │
│  ├─exception  异常层
│  │      error_handler.go  异常捕捉
│  │      response.go   返回结构
│  │
│  ├─routers    路由层
│  │      relatice_path.go  路由转发
│  │
│  ├─utils  常用方法
│  │      datax.go  常用日期类型-练手版本
│  │
│  ├─web    控制层
│  │      web.go    控制+业务处理（内含一个swagger示例）
│  │
│  └─xlogger    日志层
│          log.go   日志结构
│          logCore.go   日志业务
│
└─support   配置文件读取
    └─genv
            dotfile.go
            loader.go
            unmarshal.go
```

### 个人心得
```.env
该版本有待改进
```