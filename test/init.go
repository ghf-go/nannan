package test

import "github.com/ghf-go/nannan/app"

func init() {
	app.RegisterCli("esFind", esFind)
	app.RegisterCli("esDelete", esDelete)
	app.RegisterCli("esMget", esMget)
	app.RegisterCli("esSave", esSave)
	app.RegisterCli("esSearch", esSearch)
}
