# 虎扑篮球直播命令行版  
![Duncan](images/Duncan.jpg)

灵感来自另一位JR, 本来想fork他的工程继续改的,但是发现自己都没用到他的代码。所以就自己新建一个项目了.

### 测试环境
* macos 10.12.6
* Python 3.6.1 

## requirements

```
lxml==3.7.3
requests==2.18.4
six==1.11.0
docopt==0.6.2
user_agent==0.1.9
```

<!-- ### 如何安装
1. ``` git clone https://github.com/chenjiandongx/HupuLive.git ```
2. ``` cd HupuLive ```  
3. ``` pip install -r requirements.txt```
3. ``` python setup.py install ```   -->

<!-- ### 使用指南  
```hupu -h``` 或 ```hupu --help``` 能够查看如何使用，明细各项参数功能  

![使用指南](https://github.com/chenjiandongx/HupuLive/blob/master/images/hupu-0.gif)  

### 获取比赛直播场次  
```hupu -l``` 或 ```hupu --list``` 查询当天比赛的直播的场次，结果返回比赛场次，包括对阵双方以及场次的序号  

![获取比赛直播场次](https://github.com/chenjiandongx/HupuLive/blob/master/images/hupu-2.gif)  

### 选取比赛开始直播  
```hupu -w``` 或 ```hupu --watch``` 根据获得的场次序号来选择具体的比赛，比如这里的 0  

![选取比赛开始直播](https://github.com/chenjiandongx/HupuLive/blob/master/images/hupu-3.gif)  

对齐看起来很舒服有没有，强迫症的福音有没有！！！  
如果不想看了可以按 Ctrl-C 来中断直播，或者直接关闭终端就行了  

### 获取比赛统计数据  
```hupu -d``` 或 ```hupu --data``` 根据获取的场次序号来选择具体比赛的统计数据  

![获取比赛统计数据](https://github.com/chenjiandongx/HupuLive/blob/master/images/hupu-4.gif)  
数据也是对齐的看起来也是很爽的有没有！！！  

### 获取比赛赛后新闻
```hupu -n``` 或 ```hupu --news``` 同样根据获取的场次序号来选择具体比赛的赛后新闻  

![获取比赛赛后新闻](https://github.com/chenjiandongx/HupuLive/blob/master/images/hupu-5.gif)  

### 获取近期比赛赛程
```hupu -s``` 或 ```hupu --schedule``` 查看近七天的比赛赛程  

![获取近期比赛赛程](https://github.com/chenjiandongx/HupuLive/blob/master/images/hupu-1.gif)

### 如何卸载
使用 ```pip uninstall HupuLive``` 卸载 -->


## To-do list
* [x] http接口get, post, 参数规则
* [x] 获取比赛列表
* [x] 获取直播文字历史记录
* [ ] 命令行界面的设计
* [ ] 框架的设计
* [ ] 基类的设计
* [ ] websocket的研究
* [x] 直播数据tcp的连接
* [x] package 的 import 问题
* [ ] 心跳延续的问题
* [ ] curses 熟悉
* [ ] websocket多次返回比赛列表之后没有更多文字直播信息了？
