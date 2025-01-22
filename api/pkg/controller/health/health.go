package health

import (
	"net/http"

	"github.com/silversixx/s3-go/internal/controller"
	"github.com/silversixx/s3-go/pkg/httputils"
)

type Health struct {
	controller.CommonController[interface{}, interface{}]
}

func (*Health) GET(w http.ResponseWriter, r *http.Request) {
	httputils.ResponseJson(w, http.StatusOK, nil)
}
