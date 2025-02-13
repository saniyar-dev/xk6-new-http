package response

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/grafana/sobek"
	"github.com/saniyar-dev/xk6-new-http/pkg/helpers"
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
func (r *Response) Define() error {
	rt := r.Vu.Runtime()

	helpers.Must(rt, r.Obj.DefineDataProperty(
		"json", rt.ToValue(r.jsonAsync), sobek.FLAG_FALSE, sobek.FLAG_FALSE, sobek.FLAG_TRUE))
	return nil
}

type jsonResponse struct {
	Status     string      `json:"status"`
	StatusCode int         `json:"status_code"`
	Header     http.Header `json:"header"`
	Body       []byte      `json:"body"`
}

func (r *Response) json() ([]byte, error) {
	// TODO: make dynamic buffer for reading from body, maybe add some helper for it to use the functionality global
	_, body, err := helpers.DynamicRead(r.Body.Read, 1*time.Second)
	if err != nil {
		return []byte{}, err
	}

	res := &jsonResponse{
		Header:     r.Header,
		Body:       body,
		Status:     r.Status,
		StatusCode: r.StatusCode,
	}
	return json.Marshal(res)
}

func (r *Response) jsonAsync() *sobek.Promise {
	enqCallback := r.Vu.RegisterCallback()
	p, resolve, reject := r.Vu.Runtime().NewPromise()

	go func() {
		enqCallback(func() error {
			res, err := r.json()
			if err != nil {
				if er := reject(err); er != nil {
					return er
				}
			}
			if er := resolve(res); er != nil {
				return er
			}
			return nil
		})
	}()

	return p
}
