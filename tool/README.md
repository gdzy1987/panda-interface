# pgrpc (panda-tool)

The pgrpc 是面向 Panda 开发者的开发工具，方便快速生成框架代码。

## Getting Started
### Required
- [go](https://golang.org/dl/)
- [protoc](https://github.com/protocolbuffers/protobuf)
- [protoc-gen-go](https://github.com/protocolbuffers/protobuf-go)

### Quick Start
```bash
# 安装
$ go get -u github.com/XM-GO/panda-interface/tool/cmd/pgrpc

# 创建项目模板
$ pgrpc new github.com/PandaGoAdmin/helloworld

$ cd helloworld

# 下载必须的插件
$ make init

# 生成 proto 模板
$ pgrpc proto add api/helloworld/v1/helloworld.proto

# 生成 error proto 模板
$ pgrpc proto add api/helloworld/v1/error.proto

# 下载必须的插件
$ make api

# 生成 service 模板
$ pgrpc proto service api/helloworld/v1/helloworld.proto -t pkg/service

# 生成 server 模板(此输出需要手工加入 cmd/helloworld/main.go 中)
$ pgrpc proto server api/helloworld/v1/helloworld.proto

# 生成 api 的 makedown 文件
$ pgrpc markdown -f api/apidocs.swagger.json  -t third_party/markdown-templates/ -o ./docs/API/Greeter -m all

# 运行程序
$ go run cmd/helloworld/main.go
```



