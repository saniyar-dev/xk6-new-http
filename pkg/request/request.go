package request

import (
	"net/http"

	"github.com/grafana/sobek"
	"github.com/saniyar-dev/xk6-new-http/pkg/interfaces"
	"go.k6.io/k6/js/modules"
)

// Request object
type Request struct {
	http.Request

	Vu modules.VU

	Obj    *sobek.Object
	params *Requestparams
}

var _ interfaces.Object = &Request{}

// Define func defines data properties on obj attatched to Request struct.
func (r *Request) Define() error {
	// rt := r.Vu.Runtime()

	return nil
}
