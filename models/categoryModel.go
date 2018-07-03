package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Category struct {
	Id              int64
	Title           string
	Created         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	TopicTime       time.Time `orm:"index"`
	TopicCount      int64
	TopicLastUserId int64
}

func SaveCategory(category *Category) error {
	o := orm.NewOrm()
	_, err := o.Insert(category)
	return err
}

func DeleteCategory(id int64) error {
	o := orm.NewOrm()
	_, err := o.Delete(&Category{Id: id})
	return err
}

func UpdateCategory(category *Category) error {
	o := orm.NewOrm()
	_, err := o.Update(category)
	return err
}

func FindAllCategory() *[]Category {
	o := orm.NewOrm()
	categorys := make([]Category, 0)
	o.QueryTable("category").All(&categorys)
	return &categorys
}

func FindCategoryById(id int64) (*Category, error) {
	o := orm.NewOrm()
	category := Category{Id: id}
	err := o.Read(&category)
	return &category, err
}
