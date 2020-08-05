package databases

import (
	//"github.com/sirupsen/logrus"
	//"gopkg.in/mgo.v2"
	//"time"
)


//读取配置,链接mongo,最后将构建好的链接句柄放置到公共变量

//链接mongo
// MongoInit 初始化mongo
//func MongoInit() error {
//	session, err := mgo.DialWithTimeout("127.0.0.1:3714", time.Duration(1)*time.Second)
//	if err != nil {
//		return err
//	}
//
//	if config.MongoUsername != "" || config.MongoPassword != "" {
//		cred := &mgo.Credential{}
//		cred.Username = config.MongoUsername
//		cred.Password = config.MongoPassword
//		cred.Source = config.MongoAuth
//
//		err = session.Login(cred)
//		if err != nil {
//			logrus.Fatalf("failed to login mongodb, %s", err.Error())
//			return err
//		}
//	}
//
//	session.SetMode(mgo.Monotonic, true)
//	session.SetPoolLimit(config.MongoDBPoolLimit)
//	Global.MongoSession = session
//	return nil
//}