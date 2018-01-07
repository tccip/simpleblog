package models

import (
	"time"

	gorm "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	USER     = "root"
	PASSWORD = "111111"
	DB_NAME  = "bolog"
)

type Topic struct {
	Id           int `gorm:"AUTO_INCREMENT"`
	Name         string
	Content      string `gorm:"size:3000"`
	Views        int    `gorm:"default:0"`
	ReplyAccount int    `gorm:"default:0"`
	AddTime      time.Time
	LastModify   time.Time
}

var db *gorm.DB = Conn()

func Conn() *gorm.DB {
	db, err := gorm.Open("mysql", USER+":"+PASSWORD+"@/"+DB_NAME+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		err.Error()
	}
	if !db.HasTable("topics") {
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Topic{})
	}
	return db
}

func AddTopic(title, content string) bool {
	topic := Topic{Name: title, Content: content, AddTime: time.Now(),
		LastModify: time.Now()}
	db.Create(&topic)
	return true
}

func GetAllTopic() []Topic {
	var topics []Topic
	db.Find(&topics)
	return topics
}

func DelTopic(id int) {
	var topic Topic
	db.Where("id=?", id).Delete(&topic)
}

func QueryTopic(id int) Topic {
	var topic Topic
	db.Where("id=?", id).Find(&topic)

	topic.Views++
	db.Save(&topic)
	return topic
}

func ModifyTopic(id int, name string, content string) {
	var topic Topic
	db.Where("id=?", id).Find(&topic)
	// topic.Id = id
	// topic.Name = name
	// topic.Content = content
	// topic.LastModify = time.Now()
	// db.Save(&topic)
	db.Model(&topic).Updates(map[string]interface{}{"name": name, "content": content, "LastModify": time.Now})
}
