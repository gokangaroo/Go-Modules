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

## idiomApp本体

使用cmd命令行, 开启诗句玩法, 好像有点问题...

```bash
go build -o idiom.exe ./
idiom -cmd start -poem 两岸猿声啼不住
```

