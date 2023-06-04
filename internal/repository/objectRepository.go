package repository

import (
	"github.com/cubbit/cbs3/internal/domain"
	"gorm.io/gorm"
)

type ObjectRepository interface {
	PutObject(object *domain.Object) error
	GetObject(bucket string, key string) (domain.Object, error)
	ListObjects(bucket string, marker string, maxkeys int, prefix string) ([]domain.Object, error)
}

type objectRepository struct {
	db *gorm.DB
}

func NewObjectRepository(db *gorm.DB) ObjectRepository {
	return &objectRepository{db: db}
}

func (or *objectRepository) PutObject(object *domain.Object) error {

	return or.db.Create(object).Error

}

func (or *objectRepository) GetObject(bucket string, key string) (domain.Object, error) {

	obj := domain.Object{}

	err := or.db.Where("bucket_name = ? and key = ?", bucket, key).First(&obj).Error
	return obj, err

}

func (or *objectRepository) ListObjects(bucket string, marker string, maxkeys int, prefix string) ([]domain.Object, error) {

	objs := []domain.Object{}

	queryBuider := or.db.Limit(maxkeys)
	err := queryBuider.Select("key", "size", "e_tag").Where("bucket_name = ? and key like ? and key > ? ", bucket, prefix+"%", marker).Find(&objs).Error
	return objs, err

}
