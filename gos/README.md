####介绍
    Gos是一个基于Golang的游戏服务器框架

####启动
	git checkout https://github.com/mafei198/gos.git
	go install server && bin/server

######已实现功能
1. GenServer类似于Erlang的Behavior，封装了Goroutine的基本启动、查找、call、cast，以及对其callback module方法的动态调用
2. Routes 根据请求协议号，动态的将请求分发到对应的玩家Goroutine处理
3. Connection 建立TCP连接后自动分配Player Goroutine
4. 热数据按需加载、定时持久化至MySQL
5. 通讯协议Generator
6. Table Structs Generator
7. Broadcast

####试玩总结
Golang 优点:

1. 运行效率高：单核TCP ECHO Server， 每秒能处理7w+ 请求(PS: Erlang 4核 7w+)
2. 典型的C语言风格，流水式开发，面向过程、面向对象都非常方便，容易理解
3. 结构体与内置Map的支持使得开发非常方便

Golang 缺点:

1. 多核Scheduler目前实现的还不是很好，GOMAXPROCS的多核利用性能还不能达到其单核 * 多进程的功效，所以如果要充分发挥其性能需要配合多进程的使用，但是进程间的rpc调用将会是一个性能瓶颈
2. 强类型，静态语言，在写框架的时候比较麻烦
3. 没有内置REPL交互环境，需要自己写交互Inspect运行时的环境
4. 没有内置监督树支持，goroutine需要自行管理(锁死、crash)

####TODO
* Timertask Manager
  1. 增加、删除、修改定时服务
  2. work pool管理
  
* Timertask worker
  1. 定时任务执行

* 游戏配置文件生成器

* 运行时REPL交互环境
