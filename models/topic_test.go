package models

import (
	"fmt"
	"testing"
	"time"
)

var topic Topic = Topic{
	Title:     "testTopic",
	Content:   "测试用Topic",
	Created:   time.Now(),
	Updated:   time.Now(),
	ReplyTime: time.Now(),
}

func TestSaveTopic(t *testing.T) {
	err := SaveTopic(&topic)
	if err != nil {
		t.Error(err)
	}
}
func TestUpdateTopic(t *testing.T) {
	topic, err := FindTopicById(int64(1))
	if err != nil {
		t.Error(err)
	}
	topic.Title = "test update"
	err = UpdateTopic(topic)
	if err != nil {
		t.Error(err)
	}
}

func TestDeleteTopic(t *testing.T) {
	err := DeleteTopic(int64(2))
	if err != nil {
		t.Error(err)
	}
}

func TestFindTopicById(t *testing.T) {
	topic, err := FindTopicById(int64(1))
	if err != nil {
		t.Error(err)
	}
	fmt.Println(*topic)
}

func TestFindAllTopic(t *testing.T) {
	topics := FindAllTopic()
	for _, topic := range *topics {
		fmt.Println(topic)
	}
}
