package response

import (
	"net/http"

	"github.com/grafana/sobek"
	"github.com/saniyar-dev/xk6-new-http/pkg/interfaces"
	"go.k6.io/k6/js/modules"
)

// Response object
type Response struct {
	http.Response

	Vu modules.VU

	Obj *sobek.Object
}

var _ interfaces.Object = &Response{}

// Define func defines data properties on obj attatched to Client struct.
func (c *Response) Define() error {
	return nil
}
