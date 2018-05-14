package persistence

import (
	"aos/pkg/utils"
	"fmt"

	"github.com/go-xorm/xorm"
)

// GdSubject 科目表
type GdSubject struct {
	Id   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// GdSubjectFac 工厂
type GdSubjectFac struct {
	Table        GdSubject
	engine       *xorm.Engine
	RowsSlicePtr []GdSubject
}

// NewGdSubjectFacFac 初始化
func NewGdSubjectFacFac() *GdSubjectFac {
	var fac GdSubjectFac
	fac.engine, _ = utils.InitEng(0)
	return &fac
}

// FindAll 查询所有
func (myM *GdSubjectFac) FindAll(where string) error {
	err := myM.engine.Where(where).Find(&(myM.RowsSlicePtr))
	if err != nil {
		fmt.Println(err)
	}
	return err
}
