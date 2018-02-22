package api

import (
	"bytes"
	"net/http"
	"strconv"

	"github.com/labstack/gommon/log"
	"github.com/minio/minio-go"
)

func (a *service) getLogoHandler(w http.ResponseWriter, r *http.Request) {
	obj, err := a.fileStorage.GetObject(a.conf.StoragePayloadBucket, "logo.png", minio.GetObjectOptions{})
	if err != nil {
		log.Error(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	info, err := obj.Stat()
	if err != nil {
		log.Error(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	var b bytes.Buffer
	_, err = b.ReadFrom(obj)
	if err != nil {
		log.Error(err)
		http.Error(w, http.StatusText(500), 500)
		return
	}

	w.Header().Add("Content-Length", strconv.Itoa(int(info.Size)))
	w.Header().Add("Content-Disposition", "attachment; filename="+info.Key)
	w.WriteHeader(http.StatusOK)

	b.WriteTo(w)
}
