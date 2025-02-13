package response

import (
	"encoding/json"
	"log"
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
	// TODO: when it's in []byte the output would be base64 encoded and when it's in string it's ugly in the output
	// make a workaround for it
	Body []byte `json:"body"`
}

func (r *Response) json() ([]byte, error) {
	// TODO: make timeout configurable
	_, body, err := helpers.DynamicRead(r.Body.Read, 1*time.Second)
	if err != nil {
		return []byte{}, err
	}
	err = r.Body.Close()
	if err != nil {
		// maybe do something better?
		log.Printf("Response body couldn't be closed!: %s\n", err.Error())
	}

	res := &jsonResponse{
		Header:     r.Header,
		Body:       body,
		Status:     r.Status,
		StatusCode: r.StatusCode,
	}
	return json.MarshalIndent(res, "", "    ")
}

func (r *Response) jsonAsync() *sobek.Promise {
	enqCallback := r.Vu.RegisterCallback()
	p, resolve, reject := r.Vu.Runtime().NewPromise()

	go func() {
		res, err := r.json()
		enqCallback(func() error {
			if err != nil {
				if er := reject(err); er != nil {
					return er
				}
			}
			if er := resolve(string(res)); er != nil {
				return er
			}
			return nil
		})
	}()

	return p
}
