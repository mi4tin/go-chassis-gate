package main

import (
	"github.com/go-chassis/go-chassis"
	"github.com/go-chassis/go-chassis/core/lager"
	"github.com/go-chassis/go-chassis/core/server"

	"github.com/mi4tin/go-chassis-gate/examples/schemas"
	_ "github.com/mi4tin/go-chassis-gate/gate"
)

//if you use go run main.go instead of binary run, plz export CHASSIS_HOME=/{path}/{to}/rest/server/

func main() {
	chassis.RegisterSchema("gate", &schemas.RestGateFulHello{}, server.WithSchemaID("RestHelloService1111"))
	if err := chassis.Init(); err != nil {
		lager.Logger.Error("Init failed." + err.Error())
		return
	}
	chassis.Run()
}
