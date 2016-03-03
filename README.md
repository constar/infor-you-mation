# infor-you-mation

![logo](http://7xloce.com1.z0.glb.clouddn.com/udylogo.png)

## 简介

网址戳→ [邮订阅]

基于AngularJS做的一个适配Web和手机端显示的个人信息订阅网站

（目前还很粗糙，还在不断完善中...）

## 网站功能及名字由来  

提供关键词订阅服务，推送用户订阅的实习兼职招聘信息..

因为订阅信息来源目前只针对北邮人论坛 
所以取名 [邮订阅]  

希望能把这个网站做完善~~~
以后再做推广通用版 O(∩\_∩)O~

## 技术关键字

+ Node.js
+ Express
+ AngularJs
+ Redis
+ OpenSource

## 用法

首先下载源码

```
git clone https://github.com/constar/infor-you-mation.git
```

运行源码

```
cd infor-you-mation
npm install
PORT=3001 node ./bin/www
```

不过此时启动应该会报错，是因为本服务需要依赖Redis数据库(默认IP和端口 127.0.0.1:6379)。

假设你的开发环境是Mac 。
你可以通过以下命令安装和启动 Redis 数据库

```
brew install redis
redis-server
```

启动了Redis服务之后再运行 

```
PORT=3001 node ./bin/www
```

应该就可以看到服务启动成功的日志输出：

```
The server is now ready to accept connections on port 3001
```

然后则可以在浏览器上面打开 [http://127.0.0.1:3001](http://127.0.0.1:3001) 则可以看到 [邮订阅] 的首页显示。

但是还没完。因为此时数据库中的数据都是空，所以首页里面显示空荡荡的。

此时需要运行一下爬虫系统 [infor-you-mation-spider] ，
没错，爬虫是使用 Golang 开发的。

所以运行爬虫系统就需要你的机器已经可以运行 Golang 。

因此请运行我再次假设你的机器已经可以运行 go 和已经配置好 $GOPATH 环境变量，
那么直接运行下面的代码就可以下载 [infor-you-mation-spider] 并运行它一次。

```
go get github.com/constar/infor-you-mation-spider
$GOPATH/bin/infor-you-mation-spider
```

从此之后 Redis 里面已经有了完整的我们所需要的数据，
所以此时再次打开 [http://127.0.0.1:3001](http://127.0.0.1:3001) 
就可以看到完整的 [邮订阅] 展示样例了。

## 客服

+ Email: kaiyi0707@163.com
+ QQ: 798205246

![logo](http://7xloce.com1.z0.glb.clouddn.com/logo.png)
[邮订阅]:http://youdingyue.luckykaiyi.com
[infor-you-mation-spider](https://github.com/constar/infor-you-mation-spider)
