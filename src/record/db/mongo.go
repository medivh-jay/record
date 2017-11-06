package db

import (
	"gopkg.in/mgo.v2"
	"record/log"
)

const (
	DatabaseName = "record"
	TableName    = "pubg_play_data"
)

type Mongo struct {
	dialInfo   *mgo.DialInfo
	session    *mgo.Session
	database   *mgo.Database
	Collection *mgo.Collection
	err        error
}

func (mongo *Mongo) selectCollection(tableName string) {
	mongo.Collection = mongo.database.C(tableName)
}

func (mongo *Mongo) connection() {
	mongo.session, mongo.err = mgo.DialWithInfo(mongo.dialInfo) //连接数据库
	if mongo.err != nil {
		log.Info("mongo连接错误" + mongo.err.Error())
	}
	mongo.session.SetMode(mgo.Monotonic, true)
	mongo.database = mongo.session.DB(DatabaseName) //数据库名称
	mongo.Collection = mongo.database.C(TableName)
}

func (mongo *Mongo) Insert(data interface{}) error {
	return mongo.Collection.Insert(data)
}

func (mongo *Mongo) Select(query interface{}) *mgo.Query {
	return mongo.Collection.Find(query)
}

func (mongo *Mongo) Update(selector interface{}, update interface{}) error {
	return mongo.Collection.Update(selector, update)
}

func New(dialInfo *mgo.DialInfo) *Mongo {
	mongo := &Mongo{dialInfo: dialInfo}
	mongo.connection()
	return mongo
}
