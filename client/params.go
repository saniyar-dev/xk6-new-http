package client

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/grafana/sobek"
)

// ParseParams parses Client params and save them to it's instance
func (c *Client) ParseParams(rt *sobek.Runtime, args []sobek.Value) (Params, error) {
	parsed := &Clientparams{
		headers: make(http.Header),
	}
	if len(args) == 0 {
		c.params = parsed
		return parsed, nil
	}
	if len(args) > 1 {
		return nil, fmt.Errorf(
			"you can't have multiple arguments when creating a new Client, but you've had %d args",
			len(args),
		)
	}

	rawParams := args[0]
	params := rawParams.ToObject(rt)
	for _, k := range params.Keys() {
		switch k {
		case "headers":
			headers := params.Get(k)
			if sobek.IsUndefined(headers) || sobek.IsNull(headers) {
				continue
			}
			headersObj := headers.ToObject(rt)
			if headersObj == nil {
				continue
			}
			for _, key := range headersObj.Keys() {
				parsed.headers.Set(key, headersObj.Get(key).String())
			}

		case "proxy":
			proxy := params.Get(k)
			if sobek.IsUndefined(proxy) || sobek.IsNull(proxy) {
				continue
			}
			if v, ok := proxy.Export().(*url.URL); ok {
				parsed.proxy = *v
			}

		case "url":
			urlV := params.Get(k)
			if sobek.IsUndefined(urlV) || sobek.IsNull(urlV) {
				continue
			}
			if v, ok := urlV.Export().(*url.URL); ok {
				parsed.url = *v
			}

		default:
			return parsed, fmt.Errorf(
				"unknown Client's option: %s",
				k,
			)
		}
	}

	c.params = parsed

	return parsed, nil
}
