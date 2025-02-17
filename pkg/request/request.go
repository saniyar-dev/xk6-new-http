package request

import (
	"log"
	"net/http"

	"github.com/grafana/sobek"
	"github.com/saniyar-dev/xk6-new-http/pkg/interfaces"
	"go.k6.io/k6/js/modules"
)

// Request object
type Request struct {
	*http.Request

	Vu modules.VU

	M map[string]sobek.Value

	params *Requestparams
}

var _ interfaces.Object = &Request{}

func (r *Request) Delete(k string) bool {
	delete(r.M, k)
	return true
}

func (r *Request) Get(k string) sobek.Value {
	return r.M[k]
}

func (r *Request) Has(k string) bool {
	_, exists := r.M[k]
	return exists
}

func (r *Request) Keys() []string {
	keys := make([]string, 0, len(r.M))
	for k := range r.M {
		keys = append(keys, k)
	}
	return keys
}

func (r *Request) Set(k string, val sobek.Value) bool {
	r.M[k] = val
	return true
}

func (r *Request) print() {
	log.Print("from print function")
}

// Define func defines data properties on obj attatched to Request struct.
func (r *Request) Define() error {
	rt := r.Vu.Runtime()
	r.Request.URL = r.params.url
	r.Request.Header = r.params.headers.Clone()
	r.Set("print", rt.ToValue(r.print))
	return nil
}
