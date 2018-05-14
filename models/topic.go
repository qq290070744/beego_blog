package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	"strconv"
)

func AddTopic(title, content, category string) error {
	time.LoadLocation("Asia/Chongqing")
	o := orm.NewOrm()
	topic := &Topic{
		Uid:        0,
		Title:      title,
		Category:   category,
		Content:    content,
		Created:    time.Now(),
		Updated:    time.Now(),
		Attachment: "",
		Views:      0,
		Author:     "",
		ReplyTime:  time.Now(),
	}
	_, err := o.Insert(topic)
	if err != nil {
		return err
	}

	cate := new(Category)
	qs := o.QueryTable("category")
	err = qs.Filter("title", category).One(cate)
	if err != nil {
		return err
	}
	cate.TopicCount++
	_, err = o.Update(cate)
	return err
}

func ModifyTopic(tid, title, content, category string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()

	old_category := new(Topic)
	qs := o.QueryTable("topic")
	err = qs.Filter("id", tidNum).One(old_category)
	//fmt.Println("old_category---------------------:", old_category.Category)

	topic := &Topic{
		Id: tidNum,
	}
	if o.Read(topic) == nil {
		topic.Title = title
		topic.Category = category
		topic.Content = content
		topic.Updated = time.Now()
		o.Update(topic)
		if old_category.Category != category {
			old_cate := new(Category)
			qs := o.QueryTable("category")
			err = qs.Filter("title", old_category.Category).One(old_cate)
			if err != nil {
				return err
			}
			old_cate.TopicCount--
			_, err = o.Update(old_cate)
			if err != nil {
				return err
			}

			new_cate := new(Category)
			qs = o.QueryTable("category")
			err = qs.Filter("title", category).One(new_cate)
			if err != nil {
				return err
			}
			new_cate.TopicCount++
			_, err = o.Update(new_cate)
			if err != nil {
				return err
			}
		}
	}

	return err
}
func DeleteTopic(category string, tid string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	topic := &Topic{
		Id: tidNum,
	}
	o.Delete(topic)

	cate := new(Category)
	qs := o.QueryTable("category")
	err = qs.Filter("title", category).One(cate)
	if err != nil {
		return err
	}
	cate.TopicCount--
	_, err = o.Update(cate)
	return err

}

func GetAllTopics(cate string, isDesc bool) ([]*Topic, error) {
	o := orm.NewOrm()
	topics := make([]*Topic, 0)

	qs := o.QueryTable("topic")
	var err error
	if isDesc {
		if len(cate) > 0 {
			qs = qs.Filter("category", cate)
		}
		_, err = qs.OrderBy("-created").All(&topics)
	} else {
		_, err = qs.All(&topics)
	}

	return topics, err

}

func GetTopic(tid string) (*Topic, error) {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()
	topic := new(Topic)
	qs := o.QueryTable("topic")
	//fmt.Println("tidNum: ", tidNum)
	err = qs.Filter("id", tidNum).One(topic)
	if err != nil {
		return nil, err
	}
	topic.Views++
	_, err = o.Update(topic)
	return topic, err
}
