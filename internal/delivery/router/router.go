package router

import (
	"github.com/cubbit/cbs3/internal/delivery/handler"
	"github.com/gin-gonic/gin"
)

func NewRouter(handler handler.AppHandler) *gin.Engine {

	router := gin.Default()
	router.Handle("PUT", "/:bucket", handler.BucketHandler.CreateBucket)
	router.Handle("PUT", "/:bucket/*key", handler.ObjectHandler.PutObject)
	router.Handle("GET", "/:bucket", handler.ObjectHandler.ListObjects)
	router.Handle("GET", "/:bucket/*key", handler.ObjectHandler.GetObject)

	return router
}
