package controllers

import (
	"beeblog/models"
	"github.com/astaxie/beego"
	"log"
	"strconv"
	"time"
)

type CategoryController struct {
	beego.Controller
}

func (this *CategoryController) Get() {
	categorys := models.FindAllCategory()
	this.Data["categorys"] = categorys
	this.TplName = "categoryList.html"
}

func (this *CategoryController) Post() {
	id := this.Input().Get("id")
	title := this.Input().Get("title")

	if id != "" {
		i, err := strconv.Atoi(id)
		if err != nil {
			log.Fatal(err)
		}
		category, err := models.FindCategoryById(int64(i))
		if err != nil {
			log.Fatal(err)
		}
		category.Title = title
		models.UpdateCategory(category)
	} else {
		var category models.Category
		category.Title = title
		category.Created = time.Now()
		category.TopicTime = time.Now()

		err := models.SaveCategory(&category)
		if err != nil {
			log.Fatal(err)
		}
	}

	this.Redirect("/category", 302)
}

func (this *CategoryController) Add() {
	this.TplName = "categoryAdd.html"
}

func (this *CategoryController) Delete() {
	id, err := strconv.Atoi(this.Input().Get("id"))
	if err != nil {
		log.Fatal(err)
	}
	err = models.DeleteCategory(int64(id))
	if err != nil {
		log.Fatal(err)
	}

	this.Redirect("/category", 302)
}

func (this *CategoryController) View() {
	id, err := strconv.Atoi(this.Input().Get("id"))
	if err != nil {
		log.Fatal(err)
	}
	category, err := models.FindCategoryById(int64(id))
	if err != nil {
		log.Fatal(err)
	}
	this.Data["category"] = category
	this.TplName = "categoryView.html"
}
