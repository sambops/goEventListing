package repository

import (
	"github.com/goEventListing/entity"
	"github.com/jinzhu/gorm"
)

//TagRepoImpl implements the Tag.EventRepository
type TagRepoImpl struct{
	conn *gorm.DB
}
//NewTagRepoImpl ... will create an object of EventRepoImpl
func NewTagRepoImpl(con *gorm.DB) *TagRepoImpl{
	return &TagRepoImpl{conn:con}
}
// Tags()([]entity.Tag,[]error)
// 	Tag(id uint)(*entity.Tag,[]error)
// 	AddTag(id []int)(*entity.Tag,[]error)
// 	UpdateTag(tag *entity.Tag)(*entity.Tag,[]error)
// 	RemoveTag(id uint)(*entity.Tag,[]error)

//Tags ... returns all Tags from the database
func (tri *TagRepoImpl) Tags()([]entity.Tag,[]error){
	
}

//Tag ... returns tags associated with the given id
func(tri *TagRepoImpl) Tag(id uint)(*entity.Tag,[]error){

}

//AddTag ... adds new tag to our tag table
func (tri *TagRepoImpl) AddTag(id []int)(*entity.Tag,[]error){

}

//UpdateTag ... updates a given tag table
func (tri *TagRepoImpl) UpdateTag(tag *entity.Tag)(*entity.Tag,[]error){

}

//RemoveTag ... delete tag from  tag table with the given tag id
func (tri *TagRepoImpl) RemoveTag(id uint)(*entity.Tag,[]error){

}
