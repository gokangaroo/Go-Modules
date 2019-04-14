# 成语API
## 根据关键字查成语
http://route.showapi.com/1196-1?showapi_appid=91358&showapi_sign=a0eef1cfdd2346c8a4ffe590c50eba4c&keyword=肉&page=1&rows=10
## 根据成语名查询详情
http://route.showapi.com/1196-2?showapi_appid=91358&showapi_sign=a0eef1cfdd2346c8a4ffe590c50eba4c&keyword=不知肉味&page=1&rows=10

# 具体操作

## demoApp

使用cmd命令行, 分为模糊查询和精确查询

```bash
go build -o idiom.exe ./
idiom.exe -cmd ambiguous -keyword 肉
idiom.exe -cmd accurate -keyword 不知肉味
```

