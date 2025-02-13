package request

import (
	"net/http"

	"github.com/grafana/sobek"
	"github.com/saniyar-dev/xk6-new-http/pkg/interfaces"
)

// Requestparams struct is the options for request
type Requestparams struct {
	header http.Header
}

func (r *Request) ParseParams(rt *sobek.Runtime, args []sobek.Value) (interfaces.Params, error) {
	return &Requestparams{}, nil
}
