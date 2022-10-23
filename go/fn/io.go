package fn

import (
	"bytes"
	"encoding/json"
	"io"

	rctxv1 "github.com/yndd/lcnc-runtime/pkg/api/resourcecontext/v1"
)

type ByteReadWriter struct {
	// Reader is where ResourceContext are decoded from.
	Reader io.Reader

	// Writer is where ResourceContext are encoded.
	Writer io.Writer
}

func (rw *ByteReadWriter) Read() (*rctxv1.ResourceContext, error) {
	input := &bytes.Buffer{}
	_, err := io.Copy(input, rw.Reader)
	if err != nil {
		return nil, err
	}
	rc := &rctxv1.ResourceContext{}
	if err := json.Unmarshal(input.Bytes(), rc); err != nil {
		return nil, err
	}
	return rc, nil
}