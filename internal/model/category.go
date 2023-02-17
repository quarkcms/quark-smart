package model

import (
	"time"

	appmodel "github.com/quarkcms/quark-go/pkg/app/model"
	"github.com/quarkcms/quark-go/pkg/dal/db"
	"gorm.io/gorm"
)

// 分类模型
type Category struct {
	Id          int            `json:"id" gorm:"autoIncrement"`
	Pid         int            `json:"pid" gorm:"size:11"`
	Title       string         `json:"title" gorm:"size:200;not null"`
	Sort        int            `json:"sort" gorm:"size:11;default:10;"`
	CoverId     string         `json:"cover_id" gorm:"size:500;"`
	Name        string         `json:"name" gorm:"size:100;"`
	Description string         `json:"description" gorm:"size:500;"`
	Count       int            `json:"count" gorm:"size:11;default:10;"`
	IndexTpl    string         `json:"index_tpl" gorm:"size:100;"`
	ListTpl     string         `json:"list_tpl" gorm:"size:100;"`
	DetailTpl   string         `json:"detail_tpl" gorm:"size:100;"`
	PageNum     int            `json:"page_num" gorm:"size:11;default:10;"`
	Type        string         `json:"type" gorm:"size:200;not null;default:ARTICLE"`
	Status      int            `json:"status" gorm:"size:1;not null;default:1"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}

// Seeder
func (m *Category) Seeder() {

	// 如果菜单已存在，不执行
	if (&appmodel.Menu{}).IsExist(19) {
		return
	}

	// 创建菜单
	menuSeeders := []*appmodel.Menu{
		{Id: 19, Name: "分类列表", GuardName: "admin", Icon: "", Type: "engine", Pid: 18, Sort: 0, Path: "/api/admin/category/index", Show: 1, Status: 1},
	}
	db.Client.Create(&menuSeeders)

	// 创建默认文章
	postSeeders := []Category{
		{Title: "默认分类", Name: "default", Type: "ARTICLE", Status: 1},
	}
	db.Client.Create(&postSeeders)
}
