package registry

import (
	"github.com/cubbit/cbs3/internal/delivery/handler"
	"github.com/cubbit/cbs3/internal/repository"
	"github.com/cubbit/cbs3/internal/usecase"
	"gorm.io/gorm"
)

type interactor struct {
	db *gorm.DB
}

type Interactor interface {
	NewAppHandler() handler.AppHandler
}

func NewInteractor(pg *gorm.DB) Interactor {
	return &interactor{db: pg}
}

func (i *interactor) NewAppHandler() handler.AppHandler {
	return handler.AppHandler{
		BucketHandler: i.NewBucketHandler(),
		ObjectHandler: i.NewObjectHandler()}
}

// Bucket
func (i *interactor) NewBucketHandler() handler.BucketHandler {
	return handler.NewBucketHandler(i.NewBucketService())
}

func (i *interactor) NewBucketService() usecase.BucketService {
	return usecase.NewBucketService(i.NewBucketRepository())
}

func (i *interactor) NewBucketRepository() repository.BucketRepository {
	return repository.NewBucketRepository(i.db)
}

// Object
func (i *interactor) NewObjectHandler() handler.ObjectHandler {
	return handler.NewObjectHandler(i.NewObjectService())
}

func (i *interactor) NewObjectService() usecase.ObjectService {
	return usecase.NewObjectService(i.NewObjectRepository())
}

func (i *interactor) NewObjectRepository() repository.ObjectRepository {
	return repository.NewObjectRepository(i.db)
}
