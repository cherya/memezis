package memezis

import (
	"bytes"
	"context"
	"io"
	"net/http"

	"github.com/cherya/memezis/pkg/errors"
)

type uploadMediaResponse struct {
	Filename string `json:"filename"`
}

func (m *memezis) UploadMedia(ctx context.Context, req *http.Request) (interface{}, error) {
	file, handler, err := req.FormFile("file")
	if err != nil {
		return nil, errors.WrapMC(err, "UploadMedia: error reading file", http.StatusInternalServerError)
	}
	defer file.Close()

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		return nil, errors.WrapMC(err, "UploadMedia: error reading file", http.StatusInternalServerError)
	}

	filePath, err := m.fs.UploadTemp(buf, handler.Filename)
	if err != nil {
		return nil, errors.WrapMC(err, "UploadMedia: error saving file", http.StatusInternalServerError)
	}

	return &uploadMediaResponse{
		Filename: filePath,
	}, nil
}
