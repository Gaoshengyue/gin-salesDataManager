# SalesDataManager
各展业平台电销数据接口提供平台
from go gin
## Installation
```
$ go mod vendor
$ go mod download
```

## How to run

### Required

- Mysql
- Redis
- Mongo

### Ready

展业数据需要提前抓取到Mongo

### Conf

You should modify `conf/app.ini`

```
[app]
PageSize = 10
JwtSecret = example
PrefixUrl = http://127.0.0.1:8030

RuntimeRootPath = runtime/

ImageSavePath = upload/images/
# MB
ImageMaxSize = 5
ImageAllowExts = .jpg,.jpeg,.png

ExportSavePath = export/
QrCodeSavePath = qrcode/
FontSavePath = fonts/

LogSavePath = logs/
LogSaveName = log
LogFileExt = log
TimeFormat = 20060102

[server]
#debug or release
RunMode = debug
HttpPort = 8030
ReadTimeout = 60
WriteTimeout = 60

[database]
Type = mysql
User = root
Password = root
Host = 127.0.0.1:3306
Name = example
TablePrefix = example_

[redis]
Host = 127.0.0.1:6379
Password =
MaxIdle = 30
MaxActive = 30
IdleTimeout = 200

[mongo]
MongoDBURI=mongodb://127.0.0.1:27017
Database=example
Timeout=100
ConnectNum=200

### Run
```
$ cd $GOPATH/src/salesDataManager

$ go run main.go 
```

Project information and existing API


[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)


Listening port is 8030
Actual pid is 4393
```
#### 热加载
```gin run main.go```
```
[gin] Listening on port 3000
[gin] Building...
[gin] Build finished
```





Swagger doc

![image](https://i.imgur.com/bVRLTP4.jpg)

```shell
swag init
```

## Features

- RESTful API
- Gorm
- Swagger
- logging
- Jwt-go
- Gin
- Graceful restart or stop (fvbock/endless)
- App configurable
- Cron
- Redis
- MongoDB-driver