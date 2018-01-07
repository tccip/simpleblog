package models

import (
	"fmt"
	"time"

	gorm "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	// import _ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	// DB_HOST     = ""
	DB_USER     = "root"
	DB_PASSWORD = "111111"
	DB_NAME     = "blog"
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
	db, err := gorm.Open("mysql", DB_USER+":"+DB_PASSWORD+"@/"+DB_NAME+"?charset=utf8&parseTime=True&loc=Local")
	// db, err := gorm.Open("postgres", "host="+DB_HOST+" user="+DB_USER+" dbname="+DB_NAME+" sslmode=disable password="+DB_PASSWORD)
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
	fmt.Println(topic)
	db.Model(&topic).Updates(map[string]interface{}{"Name": name, "Content": content, "LastModfiy": time.Now()})
}
