package tag

import (
	"github.com/goEventListing/entity"
)

//TagServices ... services related to tags
type TagServices interface{
	Tags()([]entity.Tag,[]error)
	Tag(id uint)(*entity.Tag,[]error)
	AddTag(id []int)(*entity.Tag,[]error)
	UpdateTag(tag *entity.Tag)(*entity.Tag,[]error)
	RemoveTag(id uint)(*entity.Tag,[]error)
}