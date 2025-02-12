package main

import (
	"github.com/saniyar-dev/xk6-new-http/client"
	"go.k6.io/k6/js/modules"
)

func init() {
	modules.Register("k6/x/new-http", new(client.RootModule))
}
