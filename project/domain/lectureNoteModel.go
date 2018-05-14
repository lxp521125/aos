package domain

import "strings"

type LectureNote struct {
	Title        string `col:"title" json:"title"`
	Project_name string `col:"project_name" json:"-"`
	Subject_name string `col:"subject_name" json:"-"`
	Description  string `col:"description" json:"description"`
	Tag_id       int64  `json:"tag_id"`
	Creator_id   int64  `json:"creator_id"`
	Partner_id   int64  `json:"partner_id"`
	Path         string `json:"path"`
}

func (this LectureNote) GetTypeAndKey() (string, string) {
	device_str := strings.Split(DEVICE_STR, ",")
	validate := func(str string) bool {
		for _, v := range device_str {
			if v == str {
				return true
			}
		}
		return false
	}

	if validate(this.Subject_name) {
		return TYPE_PROJECT, this.Project_name
	} else {
		return TYPE_SUBJECT, this.Subject_name
	}
}
