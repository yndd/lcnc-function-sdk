package fn

import (
	"encoding/json"

	rctxv1 "github.com/yndd/lcnc-runtime/pkg/api/resourcecontext/v1"
)

func ParseResourceContext(input []byte) (*rctxv1.ResourceContext, error) {
	rc := &rctxv1.ResourceContext{}
	if err := json.Unmarshal(input, rc); err != nil {
		return nil, err
	}
	return rc, nil
}
