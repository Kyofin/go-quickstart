在windows上编译linux上能运行的程序。在这个目录下打开cmd。
执行下面内容
```
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64

go env CGO_ENABLED
 go env GOOS
 go env GOARCH
 
 go build
```
