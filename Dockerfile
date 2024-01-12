FROM golang:1.21 as builder

WORKDIR /App
COPY . .

RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w CGO_ENABLED=0 \
    && go env \
    && go mod tidy \
    && go build -o server .

FROM alpine:latest

LABEL MAINTAINER="2493381254@qq.com"

WORKDIR /App

COPY --from=0 /App ./
COPY --from=0 /App/config.docker.yaml ./

EXPOSE 8888
ENTRYPOINT ./server -c config.docker.yaml
# docker build -t chinese_learning_app .

# docker run  -d --name chinese_learning_app --network bridge --link docker_mysql:docker_mysql --link my_minio:my_minio -p 8888:8888 chinese_learning_app