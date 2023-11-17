# monologue

我的能力边界。

# PL

## Go

(Go语言圣经笔记)[gopl.io]

## Python

## Javascript/nodejs

# Network

[http 权威指南笔记](http_qwzn)

# Toolkit

# Worklog/Review

# Centralized log

# Distributed tracing

# Docker

## 用 docker 启动 nginx

ref: https://hub.docker.com/_/nginx

示例：

```bash
# 用 host 网络，这是我的一般用法，因为我基本是想启动多个 nginx 来代理本机可访问的服务，而不是为 docker 创建的其它网络进行 clb，所以用 host
# -v 用自己的配置文件，注意，配置文件中删掉所有的 include，因为不知道容器中 include 了什么东西
# 比如容器中有 include 了监听 80 端口的配置，就很容易跟本机的 80 端口冲突
sudo docker run --network host --name test_nginx -v /data/nginx_for_e2ewhisperslice.conf:/etc/nginx/nginx.conf:ro -d nginx
```

# Socat

```
sudo apt install socat -y
```

示例：

```
# 端口转发，访问本机的 8777 端口就相当于访问本机的 8888 端口
socat tcp-listen:8777,fork,reuseaddr tcp:localhost:8888
```

作用是把 socat 命令指定的两个 addr 给连接起来。

# ssh 端口转发

ref: https://zhuanlan.zhihu.com/p/148825449

以下用 = 表示网络互通，用 -> 表示 ssh 连接方向

## 本地端口转发

在这样的场景下: AOther == A -> B == BOther
想使 AOther 访问 BOther 上的服务
则可以

```
shell_A > ssh -L AOther:Port_of_A:BOther:Port_of_BOther B
# 其中 AOther 可以是 localhost/0.0.0.0/某个特定host
# 此时 Aother 可以通过访问 A:Port_of_A 来实现访问 BOther:Port_of_BOther
```

## 远程端口转发

在这样的场景下: AOther == A -> B == BOther
想使 BOther 访问 AOther 上的服务
则可以

```
shell_A > ssh -R BOther:Port_of_B:AOther:Port_of_AOther B
# 其中 BOther 可以是 localhost/0.0.0.0/某个特定host
# 此时 Bother 可以通过访问 B:Port_of_B 来实现访问 AOther:Port_of_AOther
```

# git

stash 的替代方法：
1. git commit
2. 处理完事情后切回原分支，然后 git reset HEAD

## git hooks



# netstat

# lsof

# bash

```bash
# 前台任务转后台，并在前台任务结束后执行新的命令
sleep 10 
ctrl z
fg; echo hello
```

