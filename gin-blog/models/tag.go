package models

type Tag struct {
	Model
	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

/* func (u Tag) TableName() string {
	//return "blog_tags"
	return "blog_tag"
} */

func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag, dbErr error) {
	err := db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags).Error
	if ( err != nil ){
		return nil,err
	}
	return tags, nil
}

func GetTagTotal(maps interface{}) (count int, dbErr error) {
	err := db.Model(&Tag{}).Where(maps).Count(&count).Error
	if ( err != nil ){
		return 0,err
	}
	return count, nil
}

func ExistTagByName(name string) (bool , error) {
	var tag Tag
	err := db.Select("id").Where("name = ?", name).First(&tag).Error
	if ( err != nil ){
		return false, err
	}
	if tag.ID > 0 {
		return true ,nil
	}
	return false ,nil

}

func ExistTagByID(id int) (bool, error) {
	var tag Tag
	err := db.Select("id").Where("id = ?", id).First(&tag).Error
	if ( err != nil ){
		return false, err
	}
	if tag.ID > 0 {
		return true ,nil
	}
	return false ,nil

}

func AddTag(name string, state int, createdBy string) (int, error) {
	newTag := Tag{
		Name:      name,
		State:     state,
		CreatedBy: createdBy,
	}
	err := db.Create(&newTag).Error
	if ( err != nil ){
		return 0, err
	}
	return newTag.ID ,nil
}

/*func (tag *Tag) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())
	return nil
}
func (tag *Tag) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())
	return nil
}*/

func DeleteTag(id int) (bool, error) {
	err := db.Where("id = ?", id).Delete(&Tag{}).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
func EditTag(id int, data interface{}) (bool,error) {
	err := db.Model(&Tag{}).Where("id = ?", id).Updates(data).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func CleanAllArticle() (bool, error) {
	err := db.Unscoped().Where("deleted_on != ? ", 0).Delete(&Article{}).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
