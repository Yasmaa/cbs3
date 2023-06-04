package domain

import (
	"gorm.io/gorm"
)

type Object struct {
	gorm.Model
	Key        string
	Content    []byte `type:"bytea" json:"-" xml:"-"`
	BucketName string `json:"-" xml:"-"`
	Bucket     Bucket `gorm:"foreignKey:BucketName;references:Name;" json:"-" xml:"-"`
	Size       int
	ETag       string
}

type ListObjectsResponse struct {
	Xmlns    string
	Name     string
	Prefix   string
	Marker   string
	MaxKeys  int
	Contents []Object
}
