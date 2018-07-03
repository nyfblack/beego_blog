package controllers

import (
	"beeblog/models"
	"github.com/astaxie/beego"
	"log"
	"strconv"
	"time"
)

type TopicController struct {
	beego.Controller
}

func (this *TopicController) Get() {
	topics := models.FindAllTopic()
	this.Data["topics"] = topics
	this.TplName = "topicList.html"
}

func (this *TopicController) Post() {
	title := this.Input().Get("title")
	content := this.Input().Get("content")
	//this.Input().Get("category")
	id := this.Input().Get("id")

	var topic models.Topic
	topic.Title = title
	topic.Content = content
	topic.Updated = time.Now()
	topic.ReplyTime = time.Now()
	topic.Created = time.Now()

	if id != "" {
		i, err := strconv.Atoi(id)
		if err != nil {
			log.Fatal(err)
		}
		topic.Id = int64(i)
		err = models.UpdateTopic(&topic)
	} else {
		err := models.SaveTopic(&topic)
		if err != nil {
			log.Fatal(err)
		}
	}
	this.Redirect("/topic", 302)
}

func (this *TopicController) Add() {
	this.TplName = "topicAdd.html"
}

func (this *TopicController) View() {

	id, err := strconv.Atoi(this.Input().Get("id"))
	if err != nil {
		log.Fatal(err)
	}
	topic, err := models.FindTopicById(int64(id))
	if err != nil {
		log.Fatal(err)
	}
	this.Data["topic"] = topic
	this.TplName = "topicView.html"
}

func (this *TopicController) Delete() {
	id, err := strconv.Atoi(this.Input().Get("id"))
	if err != nil {
		log.Fatal("err")
	} else {
		models.DeleteTopic(int64(id))
	}
	this.Redirect("/topic", 302)
}
