package client

import (
	"github.com/grafana/sobek"
	"github.com/saniyar-dev/xk6-new-http/helpers"
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
			"Client": i.initClient,
		},
	}
}

func (i *HTTPAPI) initClient(c sobek.ConstructorCall) *sobek.Object {
	rt := i.vu.Runtime()

	client := &Client{
		vu:  i.vu,
		obj: rt.NewObject(),
	}

	helpers.Must(rt, func() error {
		_, err := client.ParseParams(rt, c.Arguments)
		return err
	}())

	return client.obj
}
