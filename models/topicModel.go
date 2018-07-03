package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Topic struct {
	Id               int64
	Uid              int64
	Title            string
	Content          string `orm:"size(5000)"`
	Attachment       string
	Created          time.Time `orm:"index"`
	Updated          time.Time `orm:"index"`
	Views            int64
	Author           string
	ReplyTime        time.Time `orm:"index"`
	ReplyCount       int64
	RepleyLastUserId int64
}

func SaveTopic(topic *Topic) error {
	o := orm.NewOrm()
	_, err := o.Insert(topic)
	return err
}

func DeleteTopic(id int64) error {
	o := orm.NewOrm()
	_, err := o.Delete(&Topic{Id: id})
	return err
}

func UpdateTopic(topic *Topic) error {
	o := orm.NewOrm()
	_, err := o.Update(topic)
	return err
}

func FindAllTopic() *[]Topic {
	o := orm.NewOrm()

	topics := make([]Topic, 0)
	o.QueryTable("topic").OrderBy("-created").All(&topics)
	return &topics
}

func FindTopicById(id int64) (*Topic, error) {
	o := orm.NewOrm()
	topic := Topic{Id: id}
	err := o.Read(&topic)
	return &topic, err
}
