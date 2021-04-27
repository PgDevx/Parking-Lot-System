package api

import "my/v1/app"

type API struct {
	APP *app.App
}

func NewAPI() *API {
	var api API
	api.APP = app.NewApp()
	return &api
}
