package services

import (
	"github.com/goEventListing/entity"
	"github.com/goEventListing/tag"
)

// Tags()([]entity.Tag,[]error)
// 	Tag(id uint)(*entity.Tag,[]error)
// 	AddTag(tag *entity.Tag)(*entity.Tag,[]error)
// 	UpdateTag(tag *entity.Tag)(*entity.Tag,[]error)
// 	RemoveTag(id uint)(*entity.Tag,[]error)

//TagServiceImpl ... implements tag.TagRepository interface
type TagServiceImpl struct{
tagrepo tag.TagRepository
}

//NewTagServiceImpl will create new TagService object
func NewTagServiceImpl(tagRep tag.TagRepository) *TagServiceImpl{
	return &TagServiceImpl{tagrepo:tagRep}
}

//Tags returns list of tags
func(tsi *TagServiceImpl) Tags()([]entity.Tag,[]error){
	tags,errs := tsi.tagrepo.Tags()
	if len(errs) > 0 {
		return nil, errs
	}
	return tags,nil
}
//Tag returns a tag object with the given id
func(tsi *TagServiceImpl) Tag(id uint)(*entity.Tag,[]error){
	tag,errs := tsi.tagrepo.Tag(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return tag,nil
}

//AddTag adds new tag
func(tsi *TagServiceImpl) AddTag(tag *entity.Tag)(*entity.Tag,[]error){
	tag,errs := tsi.tagrepo.AddTag(tag)
	if len(errs) > 0 {
		return nil,errs
	}
	return tag,nil

}
//UpdateTag updates a tag with new data
func (tsi *TagServiceImpl) UpdateTag(tag *entity.Tag)(*entity.Tag,[]error){
	tag,errs := tsi.tagrepo.UpdateTag(tag)
	if len(errs) > 0 {
		return nil, errs
	}
	return tag,nil
}
//RemoveTag delete a tag by its id
func (tsi *TagServiceImpl) RemoveTag(id uint)(*entity.Tag,[]error){
tag,errs :=tsi.tagrepo.RemoveTag(id)
if len(errs) > 0 {
	return nil, errs
}
return tag,nil

}