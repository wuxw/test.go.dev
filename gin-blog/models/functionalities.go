package models

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"

	//	"github.com/jinzhu/gorm"
	"time"
)

type Functionalities struct {
	Action         string         `gorm:"column:action" json:"action"`
	Controller     string         `gorm:"column:controller" json:"controller"`
	CreatedAt      time.Time      `gorm:"column:created_at" json:"created_at"`
	Description    string         `gorm:"column:description" json:"description"`
	Disabled       int            `gorm:"column:disabled" json:"disabled"`
	ForefatherIds  string         `gorm:"column:forefather_ids" json:"forefather_ids"`
	Forefathers    string    `gorm:"column:forefathers" json:"forefathers"`
	ID             int            `gorm:"column:id" json:"id"`
	Menu           int            `gorm:"column:menu" json:"menu"`
	NeedCurd       int            `gorm:"column:need_curd" json:"need_curd"`
	NeedLog        int            `gorm:"column:need_log" json:"need_log"`
	NeedSearch     int            `gorm:"column:need_search" json:"need_search"`
	Parent         string         `gorm:"column:parent" json:"parent"`
	ParentID       int       `gorm:"column:parent_id" json:"parent_id"`
	Realm          int       `gorm:"column:realm" json:"realm"`
	RefreshCycle   int       `gorm:"column:refresh_cycle" json:"refresh_cycle"`
	SearchConfigID int       `gorm:"column:search_config_id" json:"search_config_id"`
	Sequence       int            `gorm:"column:sequence" json:"sequence"`
	Title          string    `gorm:"column:title" json:"title"`
	UpdatedAt      time.Time      `gorm:"column:updated_at" json:"updated_at"`
	Children []*Functionalities `orm:"-"`
}

// TableName sets the insert table name for this struct type
func (f *Functionalities) TableName() string {
	return "functionalities"
}

func GetAllFunctionalities( maps interface{}) (funs []*Functionalities, er error) {
	err := db.Where(maps).Find(&funs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil,nil
	}
	return funs, nil
}

func GetTree(list []*Functionalities) string {
	data := buildData(list)
	result := makeTreeCore(0, data)
	body, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err)
	}
	return string(body)
}

func buildData(list []*Functionalities) map[int]map[int]*Functionalities {
	//var data map[int]map[int]*Functionalities
	data := make(map[int]map[int]*Functionalities)
	for _, v := range list {
		 id := v.ID
		 parentId := v.ParentID
		if _, ok := data[parentId]; !ok {
			data[parentId] = make(map[int]*Functionalities)
		}
		data[parentId][id] = v
	}
	return  data
}


func makeTreeCore(index int, data map[int]map[int]*Functionalities) []*Functionalities {
	tmp := make([]*Functionalities, 0)
	for id, item := range data[index] {
		if data[id] != nil {
			item.Children = makeTreeCore(id, data)
		}
		tmp = append(tmp, item)
	}
	return tmp
}
