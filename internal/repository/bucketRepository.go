package repository

import (
	"github.com/cubbit/cbs3/internal/domain"
	"gorm.io/gorm"
)

type BucketRepository interface {
	CreateBucket(bucket *domain.Bucket) error
}

type bucketRepository struct {
	db *gorm.DB
}

func NewBucketRepository(db *gorm.DB) BucketRepository {
	return &bucketRepository{db: db}
}

func (br *bucketRepository) CreateBucket(bucket *domain.Bucket) error {

	return br.db.Create(bucket).Error
}
