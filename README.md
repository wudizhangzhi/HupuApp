# 虎扑篮球命令行版  
![Duncan](images/Duncan.jpg)

灵感来自另一位[JR](https://github.com/chenjiandongx/HupuLive), 本来想fork他的工程继续改的,但是发现自己都没用到他的代码。所以就自己新建一个项目了.

~~目前大部分功能支持python3.x， python2.x, 部分功能python2.x的适配还在加油中...~~

~~window的适配也在努力中...~~

# go 版本的架构

- cmd 入口
- gohupu
    - api
        - live.go 直播接口
        - base.go 接口基础
    - message
        - ~~ws.go   websocket消息~~
        - http.go http消息
    - live
        - client.go 直播客户端
    - menu
        - base.go   基础
        - live.go   直播菜单
    - logger  日志
- constant.go  常量，设置等
- utils.go     公用方法

### 测试环境
* go version go1.16.4 windows/amd64

## 安装
在[release](https://github.com/wudizhangzhi/HupuApp/releases)页面中选择符合自己系统版本的下载，
解压缩后进入文件夹直接执行`./hupu`

## 使用指南
方向键上下控制选项，回车选择，ctrl+c中途退出。

## 虎扑直播
![hupu_live](images/hupu_live.gif)

![hupu_live_contrast](images/hupu_live_contrast.gif)



## 球队数据排行

![hupu_teamranks](images/hupu_teamranks.gif)


## 球员数据

![hupu_playerdata](images/hupu_playerdata.gif)


![hupu_playerdata_injury](images/hupu_playerdata_injury.gif)

## 虎扑新闻

![hupu_news](images/hupu_news.gif)


## Hupu Api 

[Document](https://wudizhangzhi.github.io/%E7%88%AC%E8%99%AB/hupu-go-2-0-http-api.html)

## To-do list
* [ ] go版本其他功能(新闻，数据等)
* [ ] 文字美化


