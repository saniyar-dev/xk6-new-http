package api

import (
	"github.com/grafana/sobek"
	"github.com/saniyar-dev/xk6-new-http/pkg/client"
	"github.com/saniyar-dev/xk6-new-http/pkg/helpers"
	"github.com/saniyar-dev/xk6-new-http/pkg/request"
	"go.k6.io/k6/js/modules"
)

// RootModule is the root module for new-HTTP API to register the api extension.
type RootModule struct{}

var _ modules.Module = &RootModule{}

// NewModuleInstance implments NewModuleInstance func from go.k6.io/k6/js/modules.RootModule interface.
func (m *RootModule) NewModuleInstance(vu modules.VU) modules.Instance {
	return &HTTPAPI{
		vu: vu,
	}
}

// HTTPAPI is the main new-http api instance struct which implments go.k6.io/k6/js/modules.Instance interface.
type HTTPAPI struct {
	vu modules.VU
}

var _ modules.Instance = &HTTPAPI{}

// Exports implments Exports func from go.k6.io/k6/js/modules.Instance interface.
func (i *HTTPAPI) Exports() modules.Exports {
	return modules.Exports{
		Named: map[string]interface{}{
			"Client":  i.initClient,
			"Request": i.initRequest,
		},
	}
}

func (i *HTTPAPI) initClient(sc sobek.ConstructorCall) *sobek.Object {
	rt := i.vu.Runtime()

	c := &client.Client{
		Vu:  i.vu,
		Obj: rt.NewObject(),
	}

	helpers.Must(rt, func() error {
		_, err := c.ParseParams(rt, sc.Arguments)
		return err
	}())
	helpers.Must(rt, c.Define())

	return c.Obj
}

func (i *HTTPAPI) initRequest(sc sobek.ConstructorCall) *sobek.Object {
	rt := i.vu.Runtime()

	r := &request.Request{
		Vu: i.vu,
	}

	helpers.Must(rt, func() error {
		_, err := r.ParseParams(rt, sc.Arguments)
		return err
	}())

	// TODO: find another way to reconstruct the original Object cause this way we cannot implement other functionality to the object
	r.Obj = rt.ToValue(r).ToObject(rt)
	helpers.Must(rt, r.Define())

	return r.Obj
}
