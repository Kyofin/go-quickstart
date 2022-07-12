## 包引用示例
在`src/mymath`包里执行`go install`,然后在任意目录里执行`go install mymath`。
之后会在$GO_PATH里生成pkg目录，会生成`pkg/darwin_amd64/mymath.a`应用包。

之后就可以在`src/mathapp/main.go`里引用这个包了。
写完main.go后，就可以在main.go目录里执行`go build`。这样在该目录里就会生成mymath的可执行文件。
