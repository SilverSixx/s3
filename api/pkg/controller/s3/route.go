package s3

import (
	"github.com/silversixx/s3-go/internal/controller"
	"github.com/silversixx/s3-go/internal/interfaces"
)

func init() {
	interfaces.AddRoute(&S3{
		controller.CommonController[interface{}, interface{}]{
			Path: "/upload",
		},
	})
}
