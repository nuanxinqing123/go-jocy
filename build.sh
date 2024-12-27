## 编译前端文件
#cd /Users/nuanxinqing/Code/Vue/jocy-web || exit
#yarn build
#
## 返回项目目录
#cd /Users/nuanxinqing/Code/Golang/go-jocy || exit
#
## 删除现有前端文件
#rm -rf web/dist
#
## 复制前端文件到项目目录
#cp -r /Users/nuanxinqing/Code/Vue/jocy-web/dist web/
#
## 打包前端文件
#cd web || exit
#go-bindata -o=bindata/bindata.go -pkg=bindata -prefix "dist" dist/...
#cd ..

CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o Jocy-linux-arm64 -ldflags '-s -w -extldflags "-static"'
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o Jocy-linux-amd64 -ldflags '-s -w -extldflags "-static"'

# 压缩打包文件
upx Jocy-*