package usecase

import (
	"github.com/cubbit/cbs3/internal/domain"
	"github.com/cubbit/cbs3/internal/repository"
)

type ObjectService interface {
	PutObject(object *domain.Object) error
	GetObject(bucket string, key string) (domain.Object, error)
	ListObjects(bucket string, marker string, maxkeys int, prefix string) ([]domain.Object, error)
}

type objectService struct {
	ObjectRepository repository.ObjectRepository
}

func NewObjectService(or repository.ObjectRepository) ObjectService {
	return &objectService{ObjectRepository: or}
}

func (os *objectService) PutObject(object *domain.Object) error {
	return os.ObjectRepository.PutObject(object)

}

func (os *objectService) GetObject(bucket string, key string) (domain.Object, error) {

	return os.ObjectRepository.GetObject(bucket, key)
}

func (os *objectService) ListObjects(bucket string, marker string, maxkeys int, prefix string) ([]domain.Object, error) {

	return os.ObjectRepository.ListObjects(bucket, marker, maxkeys, prefix)
}
