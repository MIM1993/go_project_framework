package internal
//
//import (
//	"math/rand"
//	"strings"
//	"time"
//)
//
//import (
//	"github.com/sirupsen/logrus"
//	// "gopkg.in/mgo.v2"
//	// "gopkg.in/mgo.v2/bson"
//	"gopkg.in/redis.v3"
//)
//// RedisHandle redis handle
//type RedisHandle []*redis.Client
//
//// Init redis handle
//func (handle *RedisHandle) Init() error {
//	config := Global.Config
//	for _, addr := range strings.Split(config.RedisAddresses, ",") {
//		dialTimeout := time.Duration(config.RedisDialTimeout) * time.Second
//		readTimeout := time.Duration(config.RedisReadTimeout) * time.Second
//		writeTimeout := time.Duration(config.RedisWriteTimeout) * time.Second
//		client := redis.NewClient(&redis.Options{
//			Addr:         addr,
//			DialTimeout:  dialTimeout,
//			ReadTimeout:  readTimeout,
//			WriteTimeout: writeTimeout,
//			PoolSize:     config.RedisPoolSize,
//			Password:     config.RedisPassword,
//			DB:           config.RedisDB,
//		})
//		_, err := client.Ping().Result()
//		if err != nil {
//			logrus.Fatalf("redis init failed:%s", err.Error())
//			return err
//		}
//		handle.Append(client)
//	}
//	logrus.Infof("init redis OK, server num:%v", len(*handle))
//	return nil
//}
//
//// Append a redis client
//func (handle *RedisHandle) Append(client *redis.Client) {
//	*handle = append(*handle, client)
//}
//
//// GetClient get a redis client
//func (handle *RedisHandle) GetClient() *redis.Client {
//	if len(*handle) <= 0 {
//		err := handle.Init()
//		if err != nil {
//			logrus.Errorf("GetClient redis fail err :%v", err)
//			return nil
//		}
//	}
//	idx := rand.Intn(len(*handle))
//	return (*handle)[idx]
//}
