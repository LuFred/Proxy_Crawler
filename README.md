

# Golang实现的IP代理爬虫


### 3、安装及使用

另外，本项目用到的依赖库有：
```
gopkg.in/mgo.v2
github.com/PuerkitoBio/goquery
```

下载本项目：
```
go get -u github.com/LuFred/Proxy_Crawler
```

数据存储使用mongodb：
>需要安装mongodb

然后配置好相应的config.json并启动：
```
go build
./Proxy_Crawler
```


### 4、感谢

感谢 [henson](https://github.com/henson/ProxyPool) 提供参考。