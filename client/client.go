package client

import (
	"net/http"
	"net/url"

	"github.com/grafana/sobek"
	"go.k6.io/k6/js/modules"
)

// Clientparams struct
type Clientparams struct {
	// dial    interface{}

	// URL represents the default URL client object would use to do requests.
	URL url.URL

	// Proxy represents the default proxy client object would use to do requests.
	Proxy url.URL

	// Headers represents the default headers client object would use to do requests.
	Headers http.Header
}

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
	vu modules.VU

	// Each vu can create multiple Client objects so we need to have access to the returned sobek.Object returning to vu.
	obj *sobek.Object

	// Params is the way to config the global params for Client object to do requests.
	Params Clientparams
}
