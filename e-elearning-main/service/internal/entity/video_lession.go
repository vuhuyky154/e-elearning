package entity

import "gorm.io/gorm"

type VideoLession struct {
	gorm.Model
	Code     string                `json:"code" gorm:"unique"`
	Thumnail string                `json:"thumnail"`
	Url360p  *string               `json:"url360p"`
	Url480p  *string               `json:"url480p"`
	Url720p  *string               `json:"url720p"`
	Url1080p *string               `json:"url1080p"`
	Status   *VIDEO_LESSION_STATUS `json:"status"`

	LessionId uint     `json:"lessionId"`
	Lession   *Lession `json:"lession" gorm:"foreignKey:LessionId; constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type VIDEO_LESSION_STATUS string

const (
	VIDEO_LESSION_PENDING VIDEO_LESSION_STATUS = "VIDEO_LESSION_PENDING"
	VIDEO_LESSION_DONE    VIDEO_LESSION_STATUS = "VIDEO_LESSION_DONE"
)
