## 模块3 Docker核心技术

1. 从系统架构谈起
2. 理解Docker
3. Docker核心技术（一）
4. Docker核心技术（二）
5. 容器网络
6. Dockerfile的最佳实践

## 模块3 作业

1. 构建本地镜像
2. 编写 Dockerfile 将练习 2.2 编写的 httpserver 容器化
3. 将镜像推送至 docker 官方镜像仓库
4. 通过 docker 命令本地启动 httpserver
5. 通过 nsenter 进入容器查看 IP 配置

### 预备

1. [Docker hub](https://hub.docker.com/)注册免费个人版docker id
2. 创建个人仓库cncamp3
3. docker login

### 写作业

1. 构建本地镜像

```shell
#docker build . -t httpserver:0.0.1
```

2. 编写[Dockerfile](/module3/homework/Dockerfile)
3. 将镜像推送至 docker 官方镜像仓库

```shell
#docker push cncamp3/httpserver:0.0.1
```

4. 通过 docker 命令本地启动 httpserver

```shell
#docker run -d httpserver:0.0.1
```

5. 通过 nsenter 进入容器查看 IP 配置

```shell
#PID=$(docker inspect --format "{{ .State.Pid }}" nervous_shannon)
#nsenter -t $PID -n ip a
```