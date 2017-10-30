package db

import (
	"fmt"
	"gopkg.in/mgo.v2"
)

const (
	DatabaseUrl  = "127.0.0.1:27017"
	DatabaseName = "record"
	TableName    = "pubg_play_data"
)

type Mongo struct {
	session    *mgo.Session
	database   *mgo.Database
	Collection *mgo.Collection
	err        error
}

func (mongo *Mongo) selectCollection(tableName string) {
	mongo.Collection = mongo.database.C(tableName)
}

func (mongo *Mongo) connection() {
	mongo.session, mongo.err = mgo.Dial(DatabaseUrl) //连接数据库
	if mongo.err != nil {
		fmt.Println(mongo.err.Error())
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

func New() *Mongo {
	mongo := &Mongo{}
	mongo.connection()
	return mongo
}
