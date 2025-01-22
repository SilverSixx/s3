package s3

import (
	"net/http"

	"github.com/silversixx/s3-go/internal/controller"
	"github.com/silversixx/s3-go/pkg/command"
	"github.com/silversixx/s3-go/pkg/httputils"
)

type S3 struct {
	controller.CommonController[interface{}, interface{}]
}

type responseBody struct {
	FileUrl string `json:"file_path"`
}

func (*S3) POST(w http.ResponseWriter, r *http.Request) {

	responseBody := &responseBody{}

	err := r.ParseMultipartForm(10 << 20) // 10MB
	if err != nil {
		httputils.ResponseJsonError(w, http.StatusInternalServerError, httputils.GetErrorMsg(err), err)
		return
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		httputils.ResponseJsonError(w, http.StatusBadRequest, httputils.GetErrorMsg(err), err)
		return
	}
	defer file.Close()

	fileUrl, err := command.UploadFile(r.Context(), "datpl", handler.Filename, file)
	if err != nil {
		httputils.ResponseJsonError(w, http.StatusInternalServerError, httputils.GetErrorMsg(err), err)
		return
	}

	responseBody.FileUrl = fileUrl

	httputils.ResponseJson(w, http.StatusOK, responseBody)
}
