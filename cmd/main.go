package main

import (
	"my/v1/api"
	"my/v1/cmd/server"
)

func main() {

	r, err := api.InitHandlers()
	if err != nil {
		panic(err)
	}
	err = server.Server(r)
	if err != nil {
		panic(err)
	}

}
