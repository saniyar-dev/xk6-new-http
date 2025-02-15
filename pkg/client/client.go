package client

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/grafana/sobek"
	"github.com/saniyar-dev/xk6-new-http/pkg/helpers"
	"github.com/saniyar-dev/xk6-new-http/pkg/interfaces"
	"github.com/saniyar-dev/xk6-new-http/pkg/request"
	"github.com/saniyar-dev/xk6-new-http/pkg/response"
	"go.k6.io/k6/js/modules"
)

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

// this function would handle any type of request and do the actuall job of requesting
func (c *Client) do(req *request.Request) (*response.Response, error) {
	rt := c.Vu.Runtime()

	resp := &response.Response{
		Obj: rt.NewObject(),
		Vu:  c.Vu,
	}

	httpResp, err := c.Do(&req.Request)
	if err != nil {
		return resp, err
	}

	resp.Response = *httpResp

	helpers.Must(rt, resp.Define())
	return resp, nil
}

func createRequest(method string, arg sobek.Value) (*request.Request, error) {
	// TODO: Optimize this function
	req := &request.Request{}
	req.Method = method

	if v, ok := arg.Export().(*request.Request); ok {
		req = v
		req.Method = method
		return req, nil
	} else if v, ok := arg.Export().(string); ok {
		addr, err := url.Parse(v)
		if err != nil {
			return req, err
		}
		req.URL = addr
		return req, nil
	}

	return req, fmt.Errorf(
		"invalid input! couldn't make the request from argument: %+v",
		arg.Export())
}

func (c *Client) getAsync(arg sobek.Value) *sobek.Promise {
	enqCallback := c.Vu.RegisterCallback()
	p, resolve, reject := c.Vu.Runtime().NewPromise()

	req, err := createRequest(http.MethodGet, arg)
	if err != nil {
		// TODO: find a way to handle the rejection error
		_ = reject(err)
		return p
	}

	go func() {
		res, err := c.do(req)
		enqCallback(func() error {
			if err != nil {
				if er := reject(err); er != nil {
					return er
				}
			}
			if er := resolve(res.Obj); er != nil {
				return er
			}
			return nil
		})
	}()

	return p
}
