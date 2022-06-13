# 青训营之低配抖音

### [项目说明](https://bytedance.feishu.cn/docx/doxcnbgkMy2J0Y3E6ihqrvtHXPg)  

### [接口说明](https://www.apifox.cn/apidoc/shared-8cc50618-0da6-4d5e-a398-76f3b8f766c5/api-18345145)

## 功能特性
+ mysql
+ redis

使用的库
+ gorm
+ go-jwt
+ go-redis
+ gin
+ logrus

## 项目架构（接口设计）
### ER图  
![Diagram 1.png](https://wx1.sinaimg.cn/large/007WELPTly1h348w4eta4j30eo0d2tat.jpg)

为什么不使用外键
+ 业务数据生成顺序，未必一定可以先生成外键的值，再生成明细数据
+ 修数据：应该没有人没修过生产环境的数据吧？有外键约束，修数据就有一些麻烦了
+ 性能和扩展问题：级联控制，在应用层面做，可以降低数据库的压力。因为数据库的资源是有限资源，应用资源是可以通过加机器进行水平扩展的
+ 分库分表的场景，无法使用外键
### 中间件
![8f9e3de8c47b73fa2a82e4baa639f03.png](https://wx1.sinaimg.cn/large/007WELPTly1h34advzxelj31o40f4q6c.jpg)
### feed
![0137a676e03742b0e40452822a21b93.png](https://wx1.sinaimg.cn/large/007WELPTly1h348dvxwvbj30uw0qb40a.jpg)
### user/register
![1dfeee897b0054a37161b3c33d2823c.png](https://wx1.sinaimg.cn/large/007WELPTly1h34acnmkckj314o0ngtcn.jpg)
### user/login
![bfe11f15924830bc6bb59ac547a2195.png](https://wx1.sinaimg.cn/large/007WELPTly1h34ad2rifpj315s0f4tau.jpg)

### user
![bfe11f15924830bc6bb59ac547a2195.png](https://wx1.sinaimg.cn/large/007WELPTly1h34adm5d79j315s0f4tau.jpg)
### publish/action
![4abb172bbd09d35c5ebb670e05aff37.png](https://wx1.sinaimg.cn/large/007WELPTly1h35t0pf4wyj30jr080tb7.jpg)
### publish/list
![ce5813ab16f09b690b8483376e4586b.png](https://wx1.sinaimg.cn/large/007WELPTly1h35t1td75nj30hw07vgn0.jpg)
### favorite/action
![1dcd1f20bef523397b685340588247f.png](https://wx1.sinaimg.cn/large/007WELPTly1h370c1mh13j30v30j1q4z.jpg)
### favorite/list
![b009f4dc257f13b94bdb99497819845.png](https://wx1.sinaimg.cn/large/007WELPTly1h370cewtwqj30wd0ep0uk.jpg)
### comment/list
![8f3fca4e86c3035acff34a54d6bea72.png](https://wx1.sinaimg.cn/large/007WELPTly1h348dawg6jj311u0jxdhy.jpg)
### comment/action
![67038def61f4826b78dfd5a4dbf4891.png](https://wx1.sinaimg.cn/large/007WELPTly1h348dioax2j31cr0lyq64.jpg)


### relation/action
![image.png](https://wx1.sinaimg.cn/large/007WELPTly1h349f1fxqqj31200m8qav.jpg)
### relation/follow/list
![image.png](https://wx1.sinaimg.cn/large/007WELPTly1h349kqu3u2j30rs07q0te.jpg)
### relation/follower/list/
![image.png](https://wx1.sinaimg.cn/large/007WELPTly1h349kqu3u2j30rs07q0te.jpg)

## 快速使用
使用`docker-compose up`启动`mysql`和`redis`,再`go run ./`即可

### 测试数据

测试数据写在 demo_data.go 中，用于列表接口的 mock 测试

### 项目分工

| 姓名    | 任务        |
|-------|-----------|
| 陈友良   | 用户登录, 用户注册 |
| 焦阳,陈鹏 | 扩展接口-1    |
|李亚君|投稿,发布|
|于梓漪|视频流，用户信息|
|梁研骏|拓展接口-2|

## 提交方式 
自己新建分支如`dev-lyj`，修改完代码会`pull request`到`main`分支