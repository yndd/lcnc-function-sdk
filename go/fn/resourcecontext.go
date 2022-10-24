package fn

import (
	"encoding/json"
	"fmt"

	rctxv1 "github.com/yndd/lcnc-runtime/pkg/api/resourcecontext/v1"
)

func ParseResourceContext(input []byte) (*rctxv1.ResourceContext, error) {
	rc := &rctxv1.ResourceContext{}
	
	if err := json.Unmarshal(input, rc); err != nil {
		return nil, fmt.Errorf("error: %s with bytes: %s", err.Error(), string(input))
	}
	return rc, nil
}
