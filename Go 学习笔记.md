## 								Go 学习笔记

###### 安装

1. 默认会将 GOPATH 设置为 C 盘, 需要手动修改, 最好是放在 `Golang` 的安装目录, 因为之后使用 go get 获取的包都会在这个目录下



###### 自定义模块

1. 直接将模块放置在项目文件夹下即可
2. 如果自定义包中使用了其他模块, 进入到`项目文件夹`使用 `go install` 安装对应的模块, 之后即可正常使用
3. 模块接口命名规范 `AaBb`


###### 包管理
```base
set    GO111MODULE=on    //windows
export GO111MODULE=on    //linux
```

```base
go mod init project_name
```
使用 `go get` 命令进行安装, 就会自动将包的相关信息导入到 `*.mod` 文件中


###### 常用命令
```base
  go mod init:初始化modules
  go mod download:下载modules到本地cache
  go mod edit:编辑go.mod文件，选项有-json、-require和-exclude，可以使用帮助go help mod edit
  go mod graph:以文本模式打印模块需求图
  go mod tidy:检查，删除错误或者不使用的modules，下载没download的package
  go mod vendor:生成vendor目录
  go mod verify:验证依赖是否正确
  go mod why：查找依赖
  go test    执行一下，自动导包
  go list -m  主模块的打印路径
  go list -m -f={{.Dir}}  print主模块的根目录
  go list -m all  查看当前的依赖和版本信息
```