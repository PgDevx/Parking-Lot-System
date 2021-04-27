package app

import "my/v1/db"

type App struct {
	MongoDB *db.MongoStorage
}

func NewApp() *App {
	m := db.NewMongoStorage()
	app := App{
		MongoDB: m,
	}
	return &app
}
