#
# 服务名称（默认取当前目录名小写）
#
BINARY := $(shell basename $(CURDIR) | tr A-Z a-z)

#
# 如果是 windows 请在 git bash 中运行
#
# Debug   : make
# Clean   : make clean
# Gonew   : make gonew   （NAME 从 go.mod 第一行 module 路径最后一段读取）
#
# windows
# $env:CGO_ENABLED="0"; $env:GOOS="linux"; $env:GOARCH="amd64"; go build -buildvcs=true -o $(BINARY) main.go

# or

# go build -ldflags="-s -w" -o $(BINARY) main.go ; go tool addr2line -e $(BINARY) {ptr}

.PHONY : all clean gonew

all :
	go build -buildvcs=true -o $(BINARY) main.go

# 清除操作
clean :
	-rm -rf *.exe
	-rm -rf $(BINARY)

# ---------------------------------------- gonew 衍生补充 ----------------------------------------
NAME := $(shell head -1 go.mod | sed 's/^module //' | sed 's|.*/||')

SED_INPLACE := sed -i
ifeq ($(shell uname), Darwin)
SED_INPLACE := sed -i ''
endif

# 用法: make gonew（NAME 取自 go.mod 第一行 module 路径最后一段）
# README 中不替换 "gonew github.com/wangzhione/gohttptemplate ..." 该行（模板示例命令）

gonew :
	@[ -f go.mod ] || (echo "go.mod not found" && exit 1)
	@[ -n "$(NAME)" ] || (echo "Could not get NAME from go.mod" && exit 1)
	@echo "即将执行 gonew：将 gohttptemplate 替换为 $(NAME)，并修改 Dockerfile、README.md、.service、.gitignore。"; \
	echo "是否继续? [y/N]"; \
	read -r confirm; \
	case "$$confirm" in [yY]) ;; *) echo "已取消"; exit 1;; esac
	@$(SED_INPLACE) 's|gohttptemplate|$(NAME)|g' Dockerfile
	@echo "Dockerfile: gohttptemplate -> $(NAME)"
	@$(SED_INPLACE) '/gonew github\.com\/wangzhione\/gohttptemplate/!s|gohttptemplate|$(NAME)|g' README.md
	@echo "README.md: gohttptemplate -> $(NAME) (kept gonew example line)"
	@if [ -f gohttptemplate.service ]; then \
		mv gohttptemplate.service $(NAME).service; \
		echo "Renamed: gohttptemplate.service -> $(NAME).service"; \
	fi
	@printf '\n%s\n' "$(NAME)" >> .gitignore
	@echo ".gitignore: appended $(NAME)"
	@echo "Done."
