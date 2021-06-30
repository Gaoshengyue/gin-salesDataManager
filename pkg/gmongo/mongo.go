package gmongo

import (
	"context"
	"dolphin/salesManager/pkg/setting"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	_ "go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

var MongoClient *mongo.Database

func Setup() *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(setting.MongoDBSetting.Timeout))
	defer cancel()
	// 通过传进来的uri连接相关的配置
	fmt.Println(setting.MongoDBSetting.MongoDBURI)
	o := options.Client().ApplyURI(setting.MongoDBSetting.MongoDBURI)
	// 设置最大连接数 - 默认是100 ，不设置就是最大 max 64
	o.SetMaxPoolSize(uint64(setting.MongoDBSetting.ConnectNum))

	// 发起链接
	client, err := mongo.Connect(ctx, o)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	// 判断服务是不是可用
	if err = client.Ping(context.Background(), readpref.Primary()); err != nil {
		log.Fatal(err)
		return nil
	}
	// 返回 client
	MongoClient = client.Database(setting.MongoDBSetting.Database)
	return nil
}

func Collection(name string) *mongo.Collection {
	return MongoClient.Collection(name)
}
