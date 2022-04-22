# gover

> 一个 gin + mysql + redis 的 多用户博客 的后端示例
> 
> 之前给大一同学上课的时候写的例子，现在开源出来啦～
> 
> 代码结构什么的就按照自己喜好来啦！

## 运行

```bash
# 使用 docker 运行
docker-compose up -d
# 使用 vscode 进入容器内部开发

# 本地运行
# 1. 安装 mariadb
# 2. 安装 redis
# 3. 配置好 config.toml
# 4. 启动
go run main.go
```

## 代码结构

```bash
├── api # 接口预处理
│   ├── article_edit.go # 文章编辑
│   ├── article_info.go # 文章信息
│   ├── category_edit.go # 分类信息
│   ├── common.go # 通用
│   ├── enter.go # 入口
│   ├── user_account.go # 用户账号
│   └── user_info.go # 用户信息
├── config # 配置
│   ├── config.go # 配置加载
│   └── config.toml.example # 配置文件, 需删去 .example 后缀
├── global # 全局变量
│   └── global.go
├── initialize # 初始化
│   ├── config.go
│   ├── mysql.go
│   └── redis.go
├── middleware # 中间件
│   ├── auth.go
│   └── cors.go
├── model # 数据库 及 API接口 模型
│   ├── article
│   │   ├── model.go
│   │   ├── request.go
│   │   └── response.go
│   ├── category
│   │   ├── model.go
│   │   ├── request.go
│   │   └── response.go
│   ├── response
│   │   └── response.go
│   ├── tag
│   │   └── model.go
│   └── user
│       ├── model.go
│       ├── request.go
│       └── response.go
├── router # 路由
│   └── router.go
├── service # 服务
│   ├── article_edit.go
│   ├── article_info.go
│   ├── category_edit.go
│   ├── enter.go
│   ├── user_account.go
│   ├── user_edit.go
│   └── user_info.go
├── utils # 工具
│   ├── email.go
│   ├── jwt.go
│   └── rand.go
├── docker-compose.yaml # docker-compose 容器化开发
├── go.mod
├── go.sum
└── main.go # 入口

```