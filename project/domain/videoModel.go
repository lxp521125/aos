package domain

import (
	"strings"
)

type Video struct {
	Title        string `col:"title" json:"title"`
	Project_name string `col:"project_name" json:"-"`
	Subject_name string `col:"subject_name" json:"-"`
	Description  string `col:"description" json:"description"`
	Video_id     string `col:"video_id" json:"video_id"`
	Duration     int32  `col:"duration" json:"duration"`
	Creator_id   int64  `json:"creator_id"`
	Tag_id       int64  `json:"tag_id"`
	Partner_id   int64  `json:"partner_id"`
}

const (
	TYPE_PROJECT = "project" // 类型为project
	TYPE_SUBJECT = "subject" // 类型为subject
	DEVICE_STR   = "全科,全部"
)

func (this Video) GetTypeAndKey() (string, string) {
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
