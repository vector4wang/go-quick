#### go zero 快速开发指令
1、对应目录下执行脚本`goctl api new core`

2、编写api文件,https://go-zero.dev/docs/tasks/dsl/api

3、生成代码
```
goctl api go --api core.api --dir .
```

4、go依赖初始化
```
go mod init core
```

5、数据表