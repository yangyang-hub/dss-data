# 基础镜像，基于golang镜像构建--编译阶段
FROM registry.cn-hangzhou.aliyuncs.com/dss-pod/golang-base:1.18.4 AS builder
# 全局工作目录
WORKDIR /usr/local/go/src/dss-data
# 把运行Dockerfile文件的当前目录所有文件复制到目标目录
COPY /output /usr/local/go/src/dss-data
#  用于代理下载go项目依赖的包
RUN go env -w GO111MODULE="on"
RUN go env -w GOPROXY=https://goproxy.cn,direct
# 编译，关闭CGO，防止编译后的文件有动态链接，而alpine镜像里有些c库没有，直接没有文件的错误
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s" main.go
 
 
 
 
# 使用alpine这个轻量级镜像为基础镜像--运行阶段
FROM alpine AS runner
# 全局工作目录
WORKDIR /usr/local/go/src/dss-data
# 复制编译阶段编译出来的运行文件到目标目录
COPY --from=builder /usr/local/go/src/dss-data/main .
# 复制编译阶段里的configs文件夹到目标目录
COPY --from=builder /usr/local/go/src/dss-data/configs ./configs
# 将时区设置为东八区
RUN echo "https://mirrors.aliyun.com/alpine/v3.8/main/" > /etc/apk/repositories \
    && echo "https://mirrors.aliyun.com/alpine/v3.8/community/" >> /etc/apk/repositories \
    && apk add --no-cache tzdata \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime  \
    && echo Asia/Shanghai > /etc/timezone \
    && apk del tzdata
# docker run命令触发的真实命令(相当于直接运行编译后的可运行文件)
ENTRYPOINT ["./main"]