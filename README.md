# 技术文档

### 1.整体架构

使用gin框架实现一个接口，接口功能接收字符串表达式，进行运算，输出结果

### 2.目录结构

.
├── README.md
├── app
│   └── startRun.go
│   └── main.go
├── go.mod
├── go.sum

├── util
│   └── stackUse.go

├── internal
│   ├── ctrl
│   │   └── calController.go
│   ├── handler
│   │   ├── conversion.go
│   │   └── verification.go
│   ├── judgmenterr
│   │   ├── errStru.go
│   │   └── errHandler.go
│   ├── router
│   │   └── router.go
├── locustReport.html
├── contNumber.py
└── 程序流程图.png



### 3.代码逻辑分层

|    层     |      文件夹      |   主要职责   |        调用关系        |   其他说明   |
| :-------: | :--------------: | :----------: | :--------------------: | :----------: |
|  应用层   | app/startRun.go  |  启动服务器  |       调用路由层       | 不可同层调用 |
|  路由层   |  internet/route  |   转发路由   | 被应用层调用调用控制层 | 不可同层调用 |
|  控制层   |  internet/ctrl   | 请求参数处理 |     调用handler层      | 不可同层调用 |
| Handler层 | internet/service | 处理具体业务 |      被控制层调用      | 不可同层调用 |
|  工具层   |       util       |  创建栈结构  |     被handler调用      | 不可同层调用 |
| 压力测试  |    locust.py     | 进行压力测试 |      无敌调用关系      | 不可同层调用 |



### 4.存储设计

使用字符串切片保存分割后的中缀表达式，使用栈计算后缀表达式



### 5.接口设计

#### 请求方法

http post

#### 接口地址

localhost:8080/ping

#### 请求参数

| Key  |   Value   |
| :--: | :-------: |
| exp  | 1+2*4-6/3 |



### 6、第三方库

### gin

```
go语言的web框架
https://github.com/gin-gonic/gin
```



### 7.流程图

![未命名文件 (3)](https://user-images.githubusercontent.com/87186547/127139426-60273fe8-b032-4a63-8d17-621b9595d41f.jpg)


