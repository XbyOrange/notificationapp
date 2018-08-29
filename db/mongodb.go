package db

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
	"gitlab.com/backend/notifications/notificationApp/config"
	"gitlab.com/backend/notifications/notificationApp/logger"
	"strings"
)

const (
	templateID       = "_id"
)

type MongoDB struct {}

func (m MongoDB) Insert(template TemplateDB) error {
	var db, _ = dialDB()
	var err error
	template.Template=strings.ToUpper(template.Template)
	if err = db.C(config.MongoCollection).Insert(template); err != nil {
		logger.Print("ERROR",config.SrvName,config.Database,"1","500",err.Error(), err)
	}
	return err
}

func (m MongoDB) FindAll() ([]TemplateDB, error) {
	var db, _ = dialDB()
	var msg []TemplateDB
	var err error
	if err := db.C(config.MongoCollection).Find(bson.M{templateID: bson.RegEx{Pattern: ".*", Options: "i"},}).All(&msg); err != nil {
		logger.Print("ERROR",config.SrvName,config.Database,"1","404","Error getting all templates : " + err.Error(), err)
	}
	return msg, err
}

func (m MongoDB) Find(ID string) (TemplateDB, error) {
	var db, _ = dialDB()
	var msg TemplateDB
	var err error
	if err := db.C(config.MongoCollection).Find(bson.M{templateID:  strings.ToUpper(ID) }).One(&msg); err != nil {
		logger.Print("ERROR",config.SrvName,config.Database,"1","404","Error finding template by Id : " + ID + err.Error(), err)
	}
	return msg, err
}

func (m MongoDB) Remove(ID string) bool {
	var db, _ = dialDB()
	if err := db.C(config.MongoCollection).Remove(bson.M{templateID: strings.ToUpper(ID)}); err != nil {
		return false
	}
	return true
}

/* UpdateResponse : Update Message with Response
func (message *Message) UpdateResponse(ID string, response MessageResponse) (Message, error) {
	var msg Message
	err := db.C(config.AppConfiguration.DbConfig.Collection).Update(bson.M{messageID: ID},
		bson.M{"$set": bson.M{messageResponse: response}})
	if err != nil {
		logger.Print("ERROR",config.SrvName,config.Database,"1","404","Error updating message " + err.Error())
		return msg, err
	}
	msg.MessageResponse = response
	return msg, err
}

func FindAllByReference(reference string) ([]TemplateDB, error) {
	var msgs []TemplateDB //add limit and sort
	var err error
	if err = db.C(config.MongoCollection).Find(bson.M{content: reference}).All(&msgs); err != nil {
		logger.Print("ERROR",config.SrvName,config.Database,"1","404","Error finding template by reference " + err.Error())
	}
	return msgs, err
}

func CountAllMessagesByReference(reference string) int {
	size, _ := db.C(config.MongoCollection).Find(bson.M{content: reference}).Count()
	return size
}

func RemoveAllMessagesByReference(reference string) {
	db.C(config.MongoCollection).RemoveAll(bson.M{content: reference})
}*/

func dialDB() (*mgo.Database, error) {
	var db *mgo.Database
	_, err := mgo.Dial(config.MongoHost)
	mongoDialInfo := &mgo.DialInfo{
		Addrs:    []string{config.MongoHost},
		Timeout:  60 * time.Second,
		Database: config.MongoDB,
		Username: config.MongoUsername,
		Password: config.MongoPassword,
	}
	logger.Print("INFO",config.SrvName,config.Database,"1","200","Creating connection with mongo")
	session, err := mgo.DialWithInfo(mongoDialInfo)
	if err != nil {
		logger.Print("ERROR",config.SrvName,config.Database,"1","500","Error connecting to database " + err.Error(), err)
		return db, err
	}
	logger.Print("INFO",config.SrvName,config.Productor,"1","200","Connection created to database")
	index := mgo.Index{
		Key:        []string{templateID},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
	db = session.DB(config.MongoDB)
	db.C(config.MongoCollection).EnsureIndex(index)
	logger.Print("INFO",config.SrvName,config.Productor,"1","200","Created Index into database")
	return db, err
}
