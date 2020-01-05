package repository

import (
	"github.com/goEventListing/API/entity"
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
	tags :=[]entity.Tag{}
	errs := tri.conn.Find(&tags).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return tags,errs
	
}

//Tag ... returns tags associated with the given id
func(tri *TagRepoImpl) Tag(id uint)(*entity.Tag,[]error){
tags := entity.Tag{}
errs := tri.conn.First(&tags,id).GetErrors()
if len(errs) > 0{
	return nil,errs
}
return &tags,errs

}

//AddTag ... adds new tag to our tag table
func (tri *TagRepoImpl) AddTag(tag *entity.Tag)(*entity.Tag,[]error){
	tagg := tag
	errs := tri.conn.Create(tagg).GetErrors()
	if len(errs) > 0{
		return nil,errs
	}
	return tagg,errs
}

//UpdateTag ... updates a given tag table
func (tri *TagRepoImpl) UpdateTag(tag *entity.Tag)(*entity.Tag,[]error){
tagg := tag
errs :=tri.conn.Save(tagg).GetErrors()
if len(errs)> 0{
	return nil,errs
}
return tagg,errs
}

//RemoveTag ... delete tag from  tag table with the given tag id
func (tri *TagRepoImpl) RemoveTag(id uint)(*entity.Tag,[]error){
tag,errs:= tri.Tag(id)
if len(errs) > 0 {
	return nil, errs
}
errs = tri.conn.Delete(tag,tag.ID).GetErrors()
if len(errs) > 0 {
	return nil, errs
}
return tag,errs

}
