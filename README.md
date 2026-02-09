gohttptemplate

# 快速开始

## 使用 gonew 创建新项目

本项目提供了 gonew 模板支持，可以快速创建新的 HTTP 服务项目。

```bash
# 1. 安装
go install golang.org/x/tools/cmd/gonew@latest

# 2. 在要创建项目的父目录执行（不要进到已有非空目录里）
gonew github.com/wangzhione/gohttptemplate@latest github.com/user/myproject

cd myproject

# 3. 执行关键词替换
make gonew
```

# 测试

## postman

**1. heath 存活测试**

```bash
curl --location 'localhost:8089/health'
```

# 部署指南

## docker

**1. 构建镜像**

`cd {project}` -> `docker build -t {name}:{tag} .` 到当前目录, 基于当前目录所有文件, 开始构建镜像

```bash
docker build -t gohttptemplate:v0.0.0 .
```

**2. 查看镜像**

```bash
docker images
```

**3. 运行镜像[可选]**

```bash
docker run -d -p 8089:8089 gohttptemplate:v0.0.0
```

**4. 推送镜像到远端 [可选]**

```bash
docker login

docker push gohttptemplate:v0.0.0
```

这会将镜像上传到 Login 的 Docker 服务器。上传后，其他人就可以通过 docker pull 拉取镜像来使用。

对于 login 的服务器地址, 和组内沟通, 包括运维协助, 获取远端服务配置, 特殊会添加白名单.

## 传统平台

linux or mac 之间交叉编译

```
# 1️⃣ 编译 Linux 版本（64-bit）
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o gohttptemplate main.go 

# 2️⃣ 编译 macOS 版本（64-bit）
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o gohttptemplate main.go
```

windows 编译 linux 版本

```
$env:CGO_ENABLED="0"; $env:GOOS="linux"; $env:GOARCH="amd64"; go build -trimpath -buildvcs=true -o gohttptemplate main.go
```

## Systemd 服务部署（物理机器）

**部署命令（按顺序执行）：**

```bash
cd ... && make

# 2. 复制服务文件
sudo cp gohttptemplate.service /etc/systemd/system/

# 3. 重新加载 systemd
sudo systemctl daemon-reload

# 4. 启用开机自启
sudo systemctl enable gohttptemplate

# 5. 启动服务
sudo systemctl start gohttptemplate

# 6. 查看状态
sudo systemctl status gohttptemplate

sudo systemctl stop gohttptemplate

sudo journalctl -u gohttptemplate -f    # 查看日志
```

# 项目说明

本项目是一个 Go HTTP 服务的模板项目，支持以下功能：
- RESTful API 服务框架
- etcd 配置管理
- MySQL 数据库集成
- 中间件支持
- Docker 容器化部署
- Systemd 服务管理

## 项目结构

- `main.go` - 应用入口
- `configs/` - 配置管理
- `handler/` - HTTP 处理器和中间件
- `internal/` - 内部业务逻辑
- `register/` - 服务注册
- `common/` - 通用工具（命令行、etcd 客户端等）
- `resource/` - 资源文件（配置模板、数据库脚本）

# 技工拓展

- [Pingcap General advice](https://pingcap.github.io/style-guide/general.html)
