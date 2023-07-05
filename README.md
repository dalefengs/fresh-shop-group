# 启运冻品商城

# 运行环境
Golang：>= 1.18

Node.js：= v16.19.1 

Mysql: >= 5.7

服务器：CentOS7

# 运行环境安装

## Golang 1.18 安装
1. 下载安装包
```shell
wget https://golang.google.cn/dl/go1.18.10.linux-amd64.tar.gz
```

2. 安装
```shell
tar -zxf go1.18.10.linux-amd64.tar.gz -C /usr/local
```

3. 添加到环境变量
```shell
vim /etc/profile
```
将一下内容放入 /etc/profile 的末尾
```shell
# golang 环境变量
export GO111MODULE=on
export GOROOT=/usr/local/go
export GOPATH=/home/gopath
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
``
4. 应用环境变量
```shell
source /etc/profile
```

5. 验证
```shell
[root@QYDP-CENTOS ~]# go version
go version go1.18.10 linux/amd64
```

## 安装 Node.js
1. yum 安装设置 Node.js v16 版本
```shell
curl --silent --location https://rpm.nodesource.com/setup_16.x | sudo bash
```
yum 方式安装
```shell
sudo yum -y install nodejs
```

验证
```shell
[root@QYDP-CENTOS ~]# node --version
v16.20.1
```

2. 安装 yarn
```shell
npm install --global yarn
```
验证
```shell
[root@QYDP-CENTOS ~]# yarn -version
1.22.19
```

# 项目部署
1. 使用 FTP 或宝塔 将 `fresh-shop-group.zip` 上传至服务器 `/www/wwwroot/qiyundongpin.cn` 目录下
2. 解压 `fresh-shop-group.zip`
```shell
cd /www/wwwroot/qiyundongpin.cn
unzip x fresh-shop-group.zip
```
**配置文件路径：** `fresh-shop-group\fresh-shop\server\config.yaml

**配置文件路径：** `fresh-shop-group\fresh-shop\server\config.yaml

**配置文件路径：** `fresh-shop-group\fresh-shop\server\config.yaml



## 数据库配置
1. 使用 Navicat 工具导入`fresh-shop` 数据库
数据库文件目录: fresh-shop-group/sql/fresh-shop.sql 
2. 导入成功后修改数据库配置文件的 `mysql` 字段（大概在109行）  


## 后端服务