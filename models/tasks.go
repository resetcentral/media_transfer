package models

import "gorm.io/gorm"

type Status int

const (
	PENDING Status = iota
	COMPLETE
	FAILED
)

type Task struct {
	gorm.Model
	Status          Status
	Progress        float64
	MediaMetadataID int
	Destination     string
	DownloadUrl     string
	ErrorMsg        string
}
