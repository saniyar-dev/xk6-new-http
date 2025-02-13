package api

import (
	"github.com/saniyar-dev/xk6-new-http/pkg/api"
	"go.k6.io/k6/js/modules"
)

func init() {
	modules.Register("k6/x/net/http", new(api.RootModule))
}
