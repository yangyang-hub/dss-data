# 基础镜像，基于golang镜像构建--编译阶段
FROM registry.cn-hangzhou.aliyuncs.com/dss-pod/golang-base:1.18.4 AS builder
# 全局工作目录
WORKDIR /usr/local/go/src/dss-data
# 设定时区
ENV TZ=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
COPY . /usr/local/go/src/dss-data
#  用于代理下载go项目依赖的包
RUN go env -w GO111MODULE="on"
RUN go env -w GOPROXY=https://goproxy.cn,direct
# 编译，关闭CGO，防止编译后的文件有动态链接，而alpine镜像里有些c库没有，直接没有文件的错误
# RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s" main.go
RUN GOOS=linux GOARCH=amd64 go build main.go

 
 
# 使用alpine这个轻量级镜像为基础镜像--运行阶段
FROM registry.cn-hangzhou.aliyuncs.com/dss-pod/alpine:1.0.0 AS runner
# 全局工作目录
WORKDIR /usr/local/go/src/dss-data
# 复制编译阶段编译出来的运行文件到目标目录
COPY --from=builder /usr/local/go/src/dss-data .
# docker run命令触发的真实命令(相当于直接运行编译后的可运行文件)
ENTRYPOINT ["./main"]