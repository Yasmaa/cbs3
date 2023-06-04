package domain

import (
	"encoding/xml"
	"gorm.io/gorm"
)

type Bucket struct {
	gorm.Model
	Name     string `gorm:"primaryKey;unique;not null"`
	Location string `type:"string"`

}

type CreateBucketConfiguration struct {
	XMLName  xml.Name `xml:"CreateBucketConfiguration"`
	Location string   `xml:"LocationConstraint" type:"string"`
}
