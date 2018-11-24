package models

import "github.com/jinzhu/gorm"

type Article struct {
	Model
	TagID      int    `json:"tag_id" gorm:"index"`
	Tag        Tag    `json:"tag"`
	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Content    string `json:"content"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func ExistArticleByID(id int) (bool, error) {
	var article Article
	err := db.Select("id").Where("id = ? AND deleted_on = ? ", id, 0).First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if article.ID > 0 {
		return true, nil
	}

	return false, nil
}
func GetArticleTotal(maps interface{}) (count int,err error) {
	err = db.Model(&Article{}).Where(maps).Count(&count).Error
	if err != nil {
		return count, err
	}
	return
}
func GetArticles(pageNum int, pageSize int, maps interface{}) (articles []*Article, err error) {
	err = db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles).Error
	if err != nil {
		return nil,nil
	}
	return articles, nil
}

/*func GetArticle(id int) (article Article) {
	db.Where("id = ?", id).First(&article)
	db.Model(&article).Related(&article.Tag)
	return
}*/

func GetArticle(id int) (*Article, error) {
	var article Article
	//err := db.Where("id = ? AND deleted_on = ? ", id, 0).First(&article).Related(&article.Tag).Error
	err := db.Where("id = ? AND deleted_on = ? ", id, 0).First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &article, nil
}

/*func EditArticle(id int, data interface{}) bool {
	db.Model(&Article{}).Where("id = ?", id).Updates(data)
	return true
}*/

func EditArticle(id int, data interface{}) error {
	if err := db.Model(&Article{}).Where("id = ? AND deleted_on = ? ", id, 0).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

func AddArticle(data map[string]interface{}) (int, error) {
	var newObj Article
	newObj = Article{
		TagID:     data["tag_id"].(int),
		Title:     data["title"].(string),
		Desc:      data["desc"].(string),
		Content:   data["content"].(string),
		CreatedBy: data["created_by"].(string),
		State:     data["state"].(int),
	}
	if err := db.Create(&newObj).Error ; err != nil {
		return 0, err
	}

	return  newObj.ID, nil
}

func DeleteArticle(id int) (bool, error) {
	if err := db.Where("id = ?", id).Delete(Article{}).Error; err != nil {
		return  false, err
	}
	return true, nil
}

/*
func (article *Article) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())
	return nil
}
func (article *Article) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())
	return nil
}
*/

func CleanAllTag() bool {
	db.Unscoped().Where("deleted_on != ? ", 0).Delete(&Tag{})
	return true
}
