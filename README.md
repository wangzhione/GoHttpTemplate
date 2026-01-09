gohttptemplate

# 测试

## postman

**1. heath 存活测试**

```bash
curl --location 'localhost:8089/health'
```

# 编译

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

# 项目克隆

**所有 gohttptemplate 部分, 不区分大小写, 统一替换你的新 名称.** 

# 技工拓展

- [Pingcap General advice](https://pingcap.github.io/style-guide/general.html)
