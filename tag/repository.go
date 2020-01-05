package tag

import (
	"github.com/goEventListing/entity"
)

//add event tags
	//AddEventTag(id []int)(*entity.Tag, []error) //?? how do we add multiple tags
	//notify(eventID uint, tagsID []int) []error //this should be done separatly in notification section
	//get the event tags
	//GetTags() ([]entity.Tag, []error)

type TagRepository interface{
	Tags()([]entity.Tag,[]error)
	Tag(id uint)(*entity.Tag,[]error)
	AddTag(tag *entity.Tag)(*entity.Tag,[]error)
	UpdateTag(tag *entity.Tag)(*entity.Tag,[]error)
	RemoveTag(id uint)(*entity.Tag,[]error)
	

}