# dss-data

#### 介绍
交易决策支持系统-基础数据服务

#### 软件架构
Gin 框架构建后台，数据采用mysql存储，数据来源主要依赖tushare（https://www.tushare.pro/）以及爬虫获取

#### 使用说明

项目初始化 go.mod, 在当前工程目录下依次执行以下两个命令
1. go mod init
2. go mod tidy
3. 运行项目 go run main.go
