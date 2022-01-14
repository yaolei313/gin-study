# gin-study
* 学习代码
```
go mod init github.com/yaolei313/gin-study
go get -u github.com/gin-gonic/gin
// direct表示若前面都失败，则到此终止，直接尝试访问文件的原地址
go env -w GOPROXY=https://goproxy.cn,direct
go env -w GOPROXY=https://goproxy.cn,https://gocenter.io,https://goproxy.io,direct
// 若有的话，多个用逗号分割
go env -w GOPRIVATE=git.mycompany.com,github.com/my/private
https://github.com/golang-standards/project-layout/blob/master/README.md
```
