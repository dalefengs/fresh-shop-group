# 启运冻品商城
预览：微信小程序搜索启运冻品“启运冻品”

# 运行环境
Golang：>= 1.18

Node.js：= v16.19.1 

Mysql: >= 5.7

服务器：CentOS7

# 运行环境安装

## 微信小程序开发工具
Windows 系统中下载 微信小程序开发工具
下载地址：https://developers.weixin.qq.com/miniprogram/dev/devtools/download.html

## HBuilder X
Windows 系统中下载 HBuilder X
下载地址：https://www.dcloud.io/hbuilderx.html

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
go env -w GOPROXY=https://goproxy.cn,direct
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
unzip fresh-shop-group.zip
```
**配置文件路径：** `fresh-shop-group\fresh-shop\server\config.yaml

**配置文件路径：** `fresh-shop-group\fresh-shop\server\config.yaml

**配置文件路径：** `fresh-shop-group\fresh-shop\server\config.yaml



## 数据库配置
1. 使用宝塔工具导入`fresh-shop` 数据库
数据库文件目录: fresh-shop-group/sql/fresh-shop.sql 
2. 导入成功后修改数据库配置文件的 `mysql` 字段（大概在109行）  

## 后端服务构建
**请确保配置文件已经修改完毕**
**构建时间较长请耐心等待**
```shell
cd /www/wwwroot/qiyundongpin.cn/fresh-shop-group/fresh-shop/server
go mod tidy
go build ./
```
验证
```shell
[root@QYDP-CENTOS server]# ls | grep server
server
```

## 后台管理系统构建
1. 构建
**构建时间较长请耐心等待**
```shell
cd /www/wwwroot/qiyundongpin.cn/fresh-shop-group/fresh-shop/web
yarn install 
yarn build
```
验证
```shell
[root@QYDP-CENTOS server]# ls | grep dist
server
```

2. 部署
宝塔面板 -> 网站 -> Go项目 -> 添加项目
项目名称：启运冻品Server
项目执行文件：/www/wwwroot/qiyundongpin.cn/fresh-shop-group/fresh-shop/server/server
项目端口：48888

## 小程序构建

### 微信小程序开发工具
Windows 系统重打开 微信小程序开发工具并登录，
点击右上角的设置 -> 安全 -> 将服务端口打开

### HBuilder X
1. Windows 系统打开 HBuilder X
点击 文件 -> 导入 -> 从目录中导入 -> fresh-shop-group/fresh-shop/fresh-shop-uniapp

2. 发行
然后点击 发行 -> 小程序 - 微信 -> 填写相关信息 -> 发行
小程序名称：启运冻品
小程序Appid：wxdd9fa969ce0103ce
其他选项不需要勾选，如果提示需要登录则注册一个账号

3. 发布
发行成功后会自动打开微信小程序开发者工具
点击上传输入版本号 1 确认即可

4. 提交审核
https://mp.weixin.qq.com/wxamp/wacodepage/getcodepage?token=599980036&lang=zh_CN
打开微信小程序后台，找到版本管理，提交审核即可
