package handler

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/cubbit/cbs3/internal/domain"
	"github.com/cubbit/cbs3/internal/usecase"
	"github.com/cubbit/cbs3/utils"

	"github.com/gin-gonic/gin"
)

type ObjectHandler interface {
	PutObject(c *gin.Context)
	GetObject(c *gin.Context)
	ListObjects(c *gin.Context)
}

type objectHandler struct {
	ObjectService usecase.ObjectService
}

func NewObjectHandler(uc usecase.ObjectService) ObjectHandler {
	return &objectHandler{ObjectService: uc}
}

func (oh *objectHandler) GetObject(c *gin.Context) {

	bucket := c.Param("bucket")
	key := strings.TrimPrefix(c.Param("key"), "/")

	file, err := oh.ObjectService.GetObject(bucket, key)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	fileSize := int64(len(file.Content))

	rangeHeader := c.Request.Header.Get("Range")
	if rangeHeader != "" {
		rg, err := utils.ParseRangeHeader(rangeHeader, fileSize)
		if err != nil {
			c.AbortWithError(http.StatusRequestedRangeNotSatisfiable, err)
			return
		}

		rangeStart := rg.Start
		rangeEnd := rg.End

		c.Header("Content-Range", rangeHeader)
		c.Header("Content-Length", fmt.Sprintf("%d", rangeEnd-rangeStart+1))
		c.Status(http.StatusOK)

		c.Writer.Write(file.Content[rangeStart : rangeEnd+1])
		return

	}

	c.Header("Content-Length", fmt.Sprintf("%d", fileSize))
	c.Header("ETag", file.ETag)

	c.Status(http.StatusOK)
	c.Writer.Write(file.Content)

}

func (oh *objectHandler) PutObject(c *gin.Context) {

	fileBytes, err := c.GetRawData()

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	etag := utils.CalculateETag(fileBytes)

	obj := &domain.Object{

		Key:        strings.TrimPrefix(c.Param("key"), "/"), // To retrieve the :key parameter without the leading slash (/)
		BucketName: c.Param("bucket"),
		Content:    fileBytes,
		Size:       len(fileBytes),
		ETag:       etag,
	}

	err = oh.ObjectService.PutObject(obj)

	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.Header("ETag", etag)
	c.Status(http.StatusOK)

}

func (oh *objectHandler) ListObjects(c *gin.Context) {

	bucket := c.Param("bucket")
	marker := c.Query("marker")
	maxKeys, err := strconv.Atoi(c.DefaultQuery("max-keys", "1000"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	prefix := c.DefaultQuery("prefix", "")

	objects, err := oh.ObjectService.ListObjects(bucket, marker, maxKeys, prefix)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	response := domain.ListObjectsResponse{
		Name:     bucket,
		Prefix:   prefix,
		Marker:   marker,
		MaxKeys:  maxKeys,
		Contents: objects,
	}

	xmlResponse, err := xml.MarshalIndent(response, "", "  ")
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.Writer.Write(xmlResponse)

}
