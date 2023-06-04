package usecase

import (
	"github.com/cubbit/cbs3/internal/domain"
	"github.com/cubbit/cbs3/internal/repository"
)

type BucketService interface {
	CreateBucket(name string, config domain.CreateBucketConfiguration) error
}

type bucketService struct {
	BucketRepository repository.BucketRepository
}

func NewBucketService(br repository.BucketRepository) BucketService {
	return &bucketService{BucketRepository: br}
}

func (us *bucketService) CreateBucket(name string, config domain.CreateBucketConfiguration) error {

	bucket := &domain.Bucket{
		Name:     name,
		Location: config.Location,
	}

	return us.BucketRepository.CreateBucket(bucket)

}
