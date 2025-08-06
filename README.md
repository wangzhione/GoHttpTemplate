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
$env:CGO_ENABLED=0; $env:GOOS="linux"; $env:GOARCH="amd64"; go build -o gohttptemplate main.go
```

# 环境配置

## vscode

`.vscode/setting.json`

```json
{

    "go.toolsManagement.autoUpdate": true,
    "go.testEnvVars": {},
    "go.testFlags": [
        "-v",
        "-count=1"
    ],
    "go.useLanguageServer": true, // 启用 gopls 作为 Go 语言服务器
    "editor.formatOnSave": true, // 每次保存时自动格式化代码
    "gopls": {
        "ui.diagnostic.staticcheck": true, // 启用 StaticCheck 代码分析
        "ui.completion.usePlaceholders": true, // 启用代码补全占位符
        "formatting.gofumpt": true, // 使用更严格的 gofumpt 代码格式
        "hoverKind": "FullDocumentation" // 悬停时显示完整文档
    },
    "[go]": {
        "editor.defaultFormatter": "golang.go",
        "editor.formatOnSave": true // 可选：启用自动格式化
    },
    "go.testTimeout": "120s"

}

```

`.vscode/launch.json`

```json
{
    // 使用 IntelliSense 了解相关属性。 
    // 悬停以查看现有属性的描述。
    // 欲了解更多信息，请访问: https://go.microsoft.com/fwlink/?linkid=830387
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch Package",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "${fileDirname}",

            "args": [
                "-f", "${fileDirname}/etc/prod.toml"
            ]
        }
    ]
}

```

随后就可以 F5 打断点, 开始 Debug

# 项目克隆

所有 gohttptemplate 部分, 不区分大小写, 统一替换你的新 名称. 

# 拓展信息

- [Effective Go](https://golang.org/doc/effective_go)
- [Pingcap General advice](https://pingcap.github.io/style-guide/general.html)
- [Uber Go Style Guide](https://github.com/uber-go/guide/blob/master/style.md)

## **git 小笔记**

```shell
git tag

git tag v0.0.2

git tag -d v0.0.2
git push origin :refs/tags/v0.0.2

git push origin --tags
```

```shell
-- 远端强制覆盖本地
git fetch --all
git reset --hard origin/master
```


