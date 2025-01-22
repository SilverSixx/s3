package health

import (
	"github.com/silversixx/s3-go/internal/controller"
	"github.com/silversixx/s3-go/internal/interfaces"
)

func init() {
	interfaces.AddRoute(&Health{
		controller.CommonController[interface{}, interface{}]{
			Path: "/health",
		},
	})
}
