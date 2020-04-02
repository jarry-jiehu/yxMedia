package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type PlayModel struct {
	Id         int64  `orm:"pk;column(id)"`       //id
	Name       string `orm:"column(name)"`        //影片名字
	Title      string `orm:"column(title)"`       //影片标题
	Introduct  string `orm:"column(introduct)"`   //影片介绍
	Type       string `orm:"column(type)"`        // 影片类型，如电影，电视剧
	Menu       string `orm:"column(menu)"`        // 菜单项
	Country    string `orm:"column(country)"`     //国家
	FilmUrl    string `orm:"column(film_url)"`    //播放地址
	IconUrl    string `orm:"column(icon_url)"`    // 图标地址
	Heat       int64  `orm:"column(heat)"`        // 热度
	CreateTime string `orm:"column(create_time)"` // 影片创建时间
}

func (this *PlayModel) TableName() string {
	return "media"
}

func (this *PlayModel) Insert(playinfo *PlayModel) error {
	o := orm.NewOrm()
	id, err := o.Insert(playinfo)
	beego.Info("insert ID ", id, err)

	return err
}

func (this *PlayModel) GetByName(name string) (*PlayModel, error) {

	o := orm.NewOrm()
	qs := o.QueryTable(new(PlayModel))

	qs = qs.Filter("name__contains", name)

	var playinfo PlayModel
	err := qs.One(&playinfo)

	return &playinfo, err
}

func (this *PlayModel) GetAll(menu string) []*PlayModel {
	o := orm.NewOrm()
	qs := o.QueryTable(new(PlayModel))

	qs = qs.Filter("menu__contains", menu)
	qs = qs.OrderBy("heat")

	var playlist []*PlayModel
	num, err := qs.All(&playlist)

	beego.Info("num :", num, "  err :", err)

	return playlist
}
