FROM registry.cn-hangzhou.aliyuncs.com/dss-pod/golang-base:1.18.4
WORKDIR /usr/local/go/src/dss-data
# 设定时区
ENV TZ=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
COPY /output /usr/local/go/src/dss-data

RUN go env -w GO111MODULE="on"
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go build -o main main.go
RUN chmod +X main
ENTRYPOINT ["./main"]