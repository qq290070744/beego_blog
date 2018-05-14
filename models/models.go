package models

import (
	"time"
	_ "github.com/mattn/go-sqlite3"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"github.com/Unknwon/com"
	"os"
	"path"
)

const (
	_DB_name        = "data/beeblog.db"
	_SQLTTE3_DRIVER = "sqlite3"
)

type Category struct {
	Id             int64
	Title          string
	Created        time.Time `orm:"index"`
	Views          int64     `orm:"index"`
	TopicTime      time.Time `orm:"index"`
	TopicCount     int64
	TopicLasUserId int64
}

type Topic struct {
	Id               int64
	Uid              int64
	Title            string
	Category         string
	Content          string    `orm:"size(5000)"`
	Attachment       string
	Created          time.Time `orm:"index"`
	Updated          time.Time `orm:"index"`
	Views            int64     `orm:"index"`
	Author           string
	ReplyTime        time.Time `orm:"index"`
	ReplyCount       int64
	RepleyLastUserId int64
}

type Comment struct {
	Id      int64
	Tid     int64
	Name    string
	Content string    `orm:"size(1000)"`
	Created time.Time `orm:"index"`
}

func RegisterDb() {
	orm.DefaultTimeLoc = time.Local
	orm.RegisterModel(new(Category), new(Topic), new(Comment))
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/beego_blog?charset=utf8&loc=Local")

}

func RegisterDb_SQLTTE3() {
	if !com.IsExist(_DB_name) {
		os.MkdirAll(path.Dir(_DB_name), os.ModePerm)
		os.Create(_DB_name)
	}
	orm.DefaultTimeLoc = time.Local
	orm.RegisterModel(new(Category), new(Topic), new(Comment))
	orm.RegisterDriver(_SQLTTE3_DRIVER, orm.DRSqlite)
	orm.RegisterDataBase("default", _SQLTTE3_DRIVER, _DB_name, 10)

}

