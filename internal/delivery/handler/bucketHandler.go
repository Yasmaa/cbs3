package handler

import (
	"fmt"
	"net/http"

	"github.com/cubbit/cbs3/internal/domain"
	"github.com/cubbit/cbs3/internal/usecase"

	"github.com/gin-gonic/gin"
)

type BucketHandler interface {
	CreateBucket(c *gin.Context)
}

type bucketHandler struct {
	BucketService usecase.BucketService
}

func NewBucketHandler(uc usecase.BucketService) BucketHandler {
	return &bucketHandler{BucketService: uc}
}

func (bh *bucketHandler) CreateBucket(c *gin.Context) {

	name := c.Param("bucket")

	var config domain.CreateBucketConfiguration

	if c.Request.Body != http.NoBody {
		if err := c.ShouldBindXML(&config); err != nil {
			fmt.Println(err.Error(), c.Request.Body == http.NoBody)
			c.XML(http.StatusBadRequest, nil)
			return
		}
	}

	if err := bh.BucketService.CreateBucket(name, config); err != nil {
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	}
	c.Status(http.StatusOK)
}
