package handlers

import (
	"encoding/json"
	"minio_example/cmd/internal/cloudstorage"
	"net/http"
	"strings"
	"sync"

	"go.uber.org/zap"
)

type IndexHandler struct {
	logger   *zap.Logger
	s3Client *cloudstorage.S3Client
	bucket   string
	counter  int
	mutex    sync.Mutex
}

func NewIndexHandler(logger *zap.Logger, s3Client *cloudstorage.S3Client, bucket string) *IndexHandler {
	return &IndexHandler{logger: logger, s3Client: s3Client, bucket: bucket}
}

func (h *IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var data map[string]string
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		h.logger.Error("error reading post body", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fileName := data["file_name"]
	contents := data["contents"]
	if fileName == "" || contents == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("you must supply a file name and contents"))
		return
	}
	if err := h.s3Client.PutObject(h.bucket, fileName, strings.NewReader(contents)); err != nil {
		h.logger.Error("error writing to bucket", zap.String("bucket", h.bucket), zap.String("file name", fileName), zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write([]byte("success"))
}
