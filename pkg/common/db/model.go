package db

import (
	"strings"

	"github.com/dtm-labs/rockscache"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
	"github.com/idchats/user_score/pkg/common/config"
	"go.mongodb.org/mongo-driver/x/bsonx"

	//"github.com/idchats/user_score/pkg/common/log"
	"context"
	"time"

	go_redis "github.com/go-redis/redis/v8"
	"github.com/idchats/user_score/pkg/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2"
)

var DB DataBases

type DataBases struct {
	MysqlDB    mysqlDB
	mgoSession *mgo.Session
	//redisPool   *redis.Pool
	mongoClient *mongo.Client
	RDB         go_redis.UniversalClient
	Rc          *rockscache.Client
	WeakRc      *rockscache.Client
	Pool        *redsync.Redsync
}

type RedisClient struct {
	client  *go_redis.Client
	cluster *go_redis.ClusterClient
	go_redis.UniversalClient
	enableCluster bool
}

func key(dbAddress, dbName string) string {
	return dbAddress + "_" + dbName
}

func init() {
	//log.NewPrivateLog(constant.LogFileName)
	var err error
	//mysql init
	initMysqlDB()
	// mongo init
	// mongodb create index

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if config.Config.Redis.EnableCluster {
		DB.RDB = go_redis.NewClusterClient(&go_redis.ClusterOptions{
			Addrs:    config.Config.Redis.DBAddress,
			Username: config.Config.Redis.DBUserName,
			Password: config.Config.Redis.DBPassWord, // no password set
			PoolSize: 50,
		})
		_, err = DB.RDB.Ping(ctx).Result()
		if err != nil {
			panic(err.Error())
		}
		clientPool := goredis.NewPool(DB.RDB)
		DB.Pool = redsync.New(clientPool)
	} else {
		DB.RDB = go_redis.NewClient(&go_redis.Options{
			Addr:     config.Config.Redis.DBAddress[0],
			Username: config.Config.Redis.DBUserName,
			Password: config.Config.Redis.DBPassWord, // no password set
			DB:       0,                              // use default DB
			PoolSize: 100,                            // 连接池大小
		})
		clientPool := goredis.NewPool(DB.RDB)
		DB.Pool = redsync.New(clientPool)
		_, err = DB.RDB.Ping(ctx).Result()
		if err != nil {
			panic(err.Error())
		}
	}
	// 强一致性缓存，当一个key被标记删除，其他请求线程会被锁住轮询直到新的key生成，适合各种同步的拉取, 如果弱一致可能导致拉取还是老数据，毫无意义
	DB.Rc = rockscache.NewClient(DB.RDB, rockscache.NewDefaultOptions())
	DB.Rc.Options.StrongConsistency = true

	// 弱一致性缓存，当一个key被标记删除，其他请求线程直接返回该key的value，适合高频并且生成很缓存很慢的情况 如大群发消息缓存的缓存
	DB.WeakRc = rockscache.NewClient(DB.RDB, rockscache.NewDefaultOptions())
	DB.WeakRc.Options.StrongConsistency = false
}

func createMongoIndex(client *mongo.Client, collection string, isUnique bool, keys ...string) error {
	db := client.Database(config.Config.Mongo.DBDatabase).Collection(collection)
	opts := options.CreateIndexes().SetMaxTime(10 * time.Second)

	indexView := db.Indexes()
	keysDoc := bsonx.Doc{}

	// 复合索引
	for _, key := range keys {
		if strings.HasPrefix(key, "-") {
			keysDoc = keysDoc.Append(strings.TrimLeft(key, "-"), bsonx.Int32(-1))
		} else {
			keysDoc = keysDoc.Append(key, bsonx.Int32(1))
		}
	}

	// 创建索引
	index := mongo.IndexModel{
		Keys: keysDoc,
	}
	if isUnique == true {
		index.Options = options.Index().SetUnique(true)
	}
	result, err := indexView.CreateOne(
		context.Background(),
		index,
		opts,
	)
	if err != nil {
		return utils.Wrap(err, result)
	}
	return nil
}
