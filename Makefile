#
# 服务名称
#
# BINARY := $(shell basename $(CURDIR) | tr A-Z a-z)

#
# 如果是 windows 请在 git bash 中运行 
#
# Debug   : make
# Clean   : make clean
#
# windows
# $env:CGO_ENABLED="0"; $env:GOOS="linux"; $env:GOARCH="amd64"; go build -buildvcs=true -o $(BINARY) main.go

# or

# go build -ldflags="-s -w" -o $(BINARY) main.go ; go tool addr2line -e $(BINARY) {ptr}

.PHONY : all clean

all :
	go build -buildvcs=true

# 清除操作
clean :
	-rm -rf *.exe
	-rm -rf $(BINARY)
