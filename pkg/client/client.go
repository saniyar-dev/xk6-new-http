package client

import (
	"net/http"
	"net/url"

	"github.com/grafana/sobek"
	"github.com/saniyar-dev/xk6-new-http/pkg/helpers"
	"github.com/saniyar-dev/xk6-new-http/pkg/interfaces"
	"github.com/saniyar-dev/xk6-new-http/pkg/response"
	"go.k6.io/k6/js/modules"
)

// Clientparams struct is the default global options for Client struct
type Clientparams struct {
	// dial    interface{}

	// url represents the default URL client object would use to do requests.
	url url.URL

	// proxy represents the default proxy client object would use to do requests.
	proxy url.URL

	// headers represents the default headers client object would use to do requests.
	headers http.Header
}

var _ interfaces.Params = &Clientparams{}

// Client struct is the Client object type that users is going to use in js like this:
//
// const client = new Client();
// const response = await client.get('https://httpbin.test.k6.io/get');
//
// you can see more usage examples in js through examples dir.
type Client struct {
	// The http.Client struct to have all the functionalities of a http.Client in Client struct
	http.Client

	// Multiple vus in k6 can create multiple Client objects so we need to have access the vu Runtime, etc.
	Vu modules.VU

	// Each vu can create multiple Client objects so we need to have access to the returned sobek.Object returning to vu.
	Obj *sobek.Object

	// Params is the way to config the global params for Client object to do requests.
	params *Clientparams
}

var _ interfaces.Object = &Client{}

// Define func defines data properties on obj attatched to Client struct.
func (c *Client) Define() error {
	rt := c.Vu.Runtime()

	helpers.Must(rt, c.Obj.DefineDataProperty(
		"get", rt.ToValue(c.getAsync), sobek.FLAG_FALSE, sobek.FLAG_FALSE, sobek.FLAG_TRUE))
	return nil
}

func (c *Client) get(url string) (*response.Response, error) {
	rt := c.Vu.Runtime()

	resp := &response.Response{
		Obj: rt.NewObject(),
		Vu:  c.Vu,
	}

	httpResp, err := c.Get(url)
	if err != nil {
		return resp, err
	}

	resp.Response = *httpResp
	return resp, nil
}

func (c *Client) getAsync(url string) *sobek.Promise {
	enqCallback := c.Vu.RegisterCallback()
	p, resolve, reject := c.Vu.Runtime().NewPromise()

	go func() {
		res, err := c.get(url)
		enqCallback(func() error {
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
