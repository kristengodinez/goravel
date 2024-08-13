package main

import (
	"github.com/goravel/framework/facades"

	"goravel/bootstrap"
)

func main() {
	// This bootstraps the framework and gets it ready for use.
	bootstrap.Boot()

	//Start http server by facades.Route().
	go func() {
		if err := facades.Route().Run(); err != nil {
			facades.Log().Errorf("Route run error: %v", err)
		}
		if err := facades.Route().RunTLS(); err != nil {
			facades.Log().Errorf("Route run error: %v", err)
		}
		if err := facades.Route().RunTLSWithCert("127.0.0.1:3000", "ca.pem", "ca.key"); err != nil {
			facades.Log().Errorf("Route run error: %v", err)
		}
	}()

	select {}
}
