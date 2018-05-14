package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	"strconv"
)

func AddCategory(name string) error {
	time.LoadLocation("Asia/Chongqing")
	o := orm.NewOrm()
	cate := &Category{Title: name, Created: time.Now(), TopicTime: time.Now()}
	qs := o.QueryTable("category")
	err := qs.Filter("title", name).One(cate)
	if err == nil {
		return err
	}
	_, err = o.Insert(cate)
	if err != nil {
		return err
	}
	return nil
}

func GetAllCategories() ([]*Category, error) {
	o := orm.NewOrm()
	cates := make([]*Category, 0)
	qs := o.QueryTable("category")
	_, err := qs.All(&cates)
	return cates, err

}
func DelCategories(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	cate := &Category{Id: cid}
	_, err = o.Delete(cate)
	if err != nil {
		return err
	}
	return err
}
