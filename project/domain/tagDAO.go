package domain

type TagDAO interface {
	GetTag(tagType, tagName string) (TagModel, error)
}
