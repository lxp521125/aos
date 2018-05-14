package domain

import "time"

type TagModel struct {
	Id           int64     `xorm:"id"`
	Left_leaf    int32     `xorm:"left_leaf"`
	Right_leaf   int32     `xorm:"right_leaf"`
	Depth        int32     `xorm:"depth"`
	Parent_id    int32     `xorm:"parent_id"`
	Path         string    `xorm:"path"`
	Name         string    `xorm:"name"`
	Slug         string    `xorm:"slug"`
	Partner_id   int32     `xorm:"partner_id"`
	PayLoad_id   int32     `xorm:"payload_id"`
	Payload_type string    `xorm:"payload_type"`
	Payload      string    `xorm:"payload"`
	Deleted_at   time.Time `xorm:"deleted_at"`
	Created_at   time.Time `xorm:"created_at"`
	Updated_at   time.Time `xorm:"updated_at"`
}
